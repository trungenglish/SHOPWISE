import { z } from "zod";

export interface SessionMessage {
  ID: string;
  SessionID: string;
  Role: string;
  Content: string;
  ReasoningGraph: string;
  PinnedProducts: string;
  CreatedAt: string;
}

export interface DecisionSession {
  ID: string;
  UserID?: string;
  AnonymousID?: string;
  Title: string;
  Status: string;
  ParentSessionID?: string;
  CreatedAt: string;
  UpdatedAt: string;
  Messages: SessionMessage[];
}

const API_BASE = "/api/v1/sessions";

function getHeaders(extraHeaders?: Record<string, string>): HeadersInit {
  const headers: Record<string, string> = {
    "Content-Type": "application/json",
  };
  const token = localStorage.getItem("token");
  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }
  let anonId = localStorage.getItem("anonymousId");
  if (!anonId) {
    anonId = crypto.randomUUID();
    localStorage.setItem("anonymousId", anonId);
  }
  headers["X-Anonymous-ID"] = anonId;

  if (extraHeaders) {
    Object.assign(headers, extraHeaders);
  }
  return headers;
}

export async function createSession(
  initialMessage: string
): Promise<DecisionSession> {
  const res = await fetch(API_BASE, {
    method: "POST",
    headers: getHeaders(),
    body: JSON.stringify({ initial_message: initialMessage }),
  });
  if (!res.ok) throw new Error("Failed to create session");
  return res.json();
}

export async function listSessions(
  limit = 20,
  offset = 0
): Promise<{ data: DecisionSession[]; meta: any }> {
  const res = await fetch(`${API_BASE}?limit=${limit}&offset=${offset}`, {
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error("Failed to list sessions");
  return res.json();
}

export async function getSession(id: string): Promise<DecisionSession> {
  const res = await fetch(`${API_BASE}/${id}`, {
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error("Failed to get session");
  return res.json();
}

export async function renameSession(id: string, title: string): Promise<void> {
  const res = await fetch(`${API_BASE}/${id}`, {
    method: "PATCH",
    headers: getHeaders(),
    body: JSON.stringify({ title }),
  });
  if (!res.ok) throw new Error("Failed to rename session");
}

export async function deleteSession(id: string): Promise<void> {
  const res = await fetch(`${API_BASE}/${id}`, {
    method: "DELETE",
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error("Failed to delete session");
}

export async function branchSession(
  id: string,
  newTitle: string
): Promise<DecisionSession> {
  const res = await fetch(`${API_BASE}/${id}/branch`, {
    method: "POST",
    headers: getHeaders(),
    body: JSON.stringify({ title: newTitle }),
  });
  if (!res.ok) throw new Error("Failed to branch session");
  return res.json();
}

export async function restoreSession(id: string): Promise<void> {
  const res = await fetch(`${API_BASE}/${id}/restore`, {
    method: "POST",
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error("Failed to restore session");
}

export async function autoSaveSession(
  id: string,
  clientTimestamp: string,
  pinnedProducts: string[]
): Promise<DecisionSession> {
  const res = await fetch(`${API_BASE}/${id}`, {
    method: "PUT",
    headers: getHeaders({ "X-Client-Timestamp": clientTimestamp }),
    body: JSON.stringify({ pinned_products: pinnedProducts }),
  });
  if (!res.ok) {
    if (res.status === 409) {
      throw new Error("conflict");
    }
    throw new Error("Failed to autosave session");
  }
  return res.json();
}

const recommendationProductSchema = z.object({
  id: z.string(),
  name: z.string(),
  price: z.number().int().nonnegative(),
  image: z.string(),
  specifications: z.record(z.string(), z.unknown()),
  match_score: z.number().int().min(0).max(100),
  match_explanation: z.string(),
});

const recommendationDecisionSchema = z.object({
  products: z.array(recommendationProductSchema).min(1),
  reasoning: z.string(),
});

const agentEnvelopeSchema = z.discriminatedUnion("type", [
  z.object({
    type: z.literal("question"),
    message: z.string().min(1),
  }),
  z.object({
    type: z.literal("recommendation"),
    message: z.string().min(1),
    decision: recommendationDecisionSchema,
  }),
  z.object({
    type: z.literal("comparison"),
    message: z.string().min(1),
    decision: recommendationDecisionSchema,
  }),
  z.object({
    type: z.literal("checkout_ready"),
    message: z.string().min(1),
    decision: recommendationDecisionSchema,
  }),
]);

export type AgentEnvelope = z.infer<typeof agentEnvelopeSchema>;

export async function sendChatMessage(
  sessionId: string,
  message: string
): Promise<AgentEnvelope> {
  const response = await fetch(`${API_BASE}/${sessionId}/chat`, {
    method: "POST",
    headers: getHeaders(),
    body: JSON.stringify({ message }),
  });
  if (!response.ok) {
    throw new Error("Agent request failed");
  }

  const parsed = agentEnvelopeSchema.safeParse(await response.json());
  if (!parsed.success) {
    throw new Error("Invalid Agent response");
  }
  return parsed.data;
}

export async function runAgentTurn(
  sessionId: string | null,
  message: string
): Promise<{ sessionId: string; envelope: AgentEnvelope }> {
  const activeSessionId = sessionId ?? (await createSession(message)).ID;
  return {
    sessionId: activeSessionId,
    envelope: await sendChatMessage(activeSessionId, message),
  };
}

function parseSSEBlock(block: string): { event: string; data: string } {
  let event = "";
  const data: string[] = [];
  for (const line of block.split(/\r?\n/)) {
    if (line.startsWith("event: ")) {
      event = line.slice(7);
    } else if (line.startsWith("data: ")) {
      data.push(line.slice(6));
    }
  }
  return { event, data: data.join("\n") };
}

export async function streamChatMessage(
  sessionId: string,
  message: string,
  onToken?: (text: string) => void
): Promise<AgentEnvelope> {
  const response = await fetch(`${API_BASE}/${sessionId}/chat/stream`, {
    method: "POST",
    headers: getHeaders(),
    body: JSON.stringify({ message }),
  });
  if (!response.ok || response.body === null) {
    throw new Error("Agent request failed");
  }

  const reader = response.body.getReader();
  const decoder = new TextDecoder();
  let buffer = "";
  let agentMessage = "";
  let decision: z.infer<typeof recommendationDecisionSchema> | undefined;
  let decisionType: "recommendation" | "comparison" | "checkout_ready" =
    "recommendation";

  const consumeBlock = (block: string) => {
    const event = parseSSEBlock(block);
    if (event.event === "token") {
      const token = z
        .object({ text: z.string() })
        .parse(JSON.parse(event.data));
      agentMessage += token.text;
      onToken?.(token.text);
    } else if (
      event.event === "recommendation" ||
      event.event === "comparison" ||
      event.event === "checkout_ready"
    ) {
      decision = recommendationDecisionSchema.parse(JSON.parse(event.data));
      decisionType = event.event;
    } else if (event.event === "error") {
      throw new Error("Agent request failed");
    }
  };

  while (true) {
    const { done, value } = await reader.read();
    buffer += decoder.decode(value, { stream: !done }).replaceAll("\r\n", "\n");
    let boundary = buffer.indexOf("\n\n");
    while (boundary >= 0) {
      consumeBlock(buffer.slice(0, boundary));
      buffer = buffer.slice(boundary + 2);
      boundary = buffer.indexOf("\n\n");
    }
    if (done) {
      break;
    }
  }
  if (buffer.trim().length > 0) {
    consumeBlock(buffer);
  }

  return agentEnvelopeSchema.parse(
    decision === undefined
      ? { type: "question", message: agentMessage }
      : { type: decisionType, message: agentMessage, decision }
  );
}

export async function runAgentTurnStream(
  sessionId: string | null,
  message: string,
  onToken?: (text: string) => void
): Promise<{ sessionId: string; envelope: AgentEnvelope }> {
  const activeSessionId = sessionId ?? (await createSession(message)).ID;
  return {
    sessionId: activeSessionId,
    envelope: await streamChatMessage(activeSessionId, message, onToken),
  };
}

export interface UserPreference {
  ID: string;
  UserID: string;
  Category: string;
  Value: string;
  SourceSessionID?: string;
  CreatedAt: string;
  UpdatedAt: string;
}

export async function listPreferences(): Promise<{ data: UserPreference[] }> {
  const res = await fetch("/api/v1/sessions/preferences", {
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error("Failed to list preferences");
  return res.json();
}

export async function updatePreference(
  id: string,
  category: string,
  value: string
): Promise<UserPreference> {
  const res = await fetch(`/api/v1/sessions/preferences/${id}`, {
    method: "PUT",
    headers: getHeaders(),
    body: JSON.stringify({ category, value }),
  });
  if (!res.ok) throw new Error("Failed to update preference");
  return res.json();
}

export async function deletePreference(id: string): Promise<void> {
  const res = await fetch(`/api/v1/sessions/preferences/${id}`, {
    method: "DELETE",
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error("Failed to delete preference");
}

export class ResumeSessionError extends Error {
  code: string;
  constructor(message: string, code: string) {
    super(message);
    this.name = "ResumeSessionError";
    this.code = code;
  }
}

export async function resumeSessionFromToken(
  token: string
): Promise<DecisionSession> {
  const res = await fetch(
    `/api/v1/session/resume?token=${encodeURIComponent(token)}`,
    {
      method: "GET",
      headers: getHeaders(),
    }
  );
  if (!res.ok) {
    const errorData = await res.json().catch(() => ({}));
    throw new ResumeSessionError(
      errorData.message || "Failed to resume session",
      errorData.error || `http_error_${res.status}`
    );
  }
  return res.json();
}
