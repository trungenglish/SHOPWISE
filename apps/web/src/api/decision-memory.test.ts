import { beforeEach, describe, expect, it, vi } from "vitest";

import {
  runAgentTurn,
  sendChatMessage,
  streamChatMessage,
} from "./decision-memory";

describe("sendChatMessage", () => {
  beforeEach(() => {
    localStorage.clear();
    localStorage.setItem("anonymousId", "anon-1");
    vi.restoreAllMocks();
  });

  it("sends the message with session identity and returns a question", async () => {
    const fetchMock = vi.fn().mockResolvedValue(
      new Response(
        JSON.stringify({
          type: "question",
          message: "What is your budget?",
        }),
        { status: 200, headers: { "Content-Type": "application/json" } }
      )
    );
    vi.stubGlobal("fetch", fetchMock);

    const response = await sendChatMessage("session-1", "Gaming laptop");

    expect(response).toEqual({
      type: "question",
      message: "What is your budget?",
    });
    expect(fetchMock).toHaveBeenCalledWith(
      "/api/v1/sessions/session-1/chat",
      expect.objectContaining({
        method: "POST",
        body: JSON.stringify({ message: "Gaming laptop" }),
        headers: expect.objectContaining({ "X-Anonymous-ID": "anon-1" }),
      })
    );
  });

  it("rejects an invalid Agent envelope", async () => {
    vi.stubGlobal(
      "fetch",
      vi.fn().mockResolvedValue(
        new Response(
          JSON.stringify({
            type: "recommendation",
            message: "Missing decision",
          }),
          {
            status: 200,
            headers: { "Content-Type": "application/json" },
          }
        )
      )
    );

    await expect(sendChatMessage("session-1", "Gaming laptop")).rejects.toThrow(
      "Invalid Agent response"
    );
  });

  it("creates a session before the first Agent turn", async () => {
    const fetchMock = vi
      .fn()
      .mockResolvedValueOnce(
        new Response(
          JSON.stringify({
            ID: "session-1",
            Title: "Gaming laptop",
            Status: "active",
            CreatedAt: "2026-08-04T00:00:00Z",
            UpdatedAt: "2026-08-04T00:00:00Z",
            Messages: [],
          }),
          { status: 201, headers: { "Content-Type": "application/json" } }
        )
      )
      .mockResolvedValueOnce(
        new Response(
          JSON.stringify({ type: "question", message: "What is your budget?" }),
          { status: 200, headers: { "Content-Type": "application/json" } }
        )
      );
    vi.stubGlobal("fetch", fetchMock);

    const result = await runAgentTurn(null, "Gaming laptop");

    expect(result.sessionId).toBe("session-1");
    expect(result.envelope.type).toBe("question");
    expect(fetchMock).toHaveBeenNthCalledWith(
      1,
      "/api/v1/sessions",
      expect.objectContaining({ method: "POST" })
    );
  });

  it("parses split SSE chunks without a parser dependency", async () => {
    const encoder = new TextEncoder();
    const stream = new ReadableStream({
      start(controller) {
        controller.enqueue(
          encoder.encode('event: token\ndata: {"text":"Catalog match')
        );
        controller.enqueue(
          encoder.encode(
            ' found."}\n\nevent: recommendation\ndata: {"products":[{"id":"product-1","name":"Laptop","price":37475000,"image":"","specifications":{},"match_score":92,"match_explanation":"Fits"}],"reasoning":"Strong GPU"}\n\nevent: done\ndata: {}\n\n'
          )
        );
        controller.close();
      },
    });
    vi.stubGlobal(
      "fetch",
      vi.fn().mockResolvedValue(
        new Response(stream, {
          status: 200,
          headers: { "Content-Type": "text/event-stream" },
        })
      )
    );
    const onToken = vi.fn();

    const response = await streamChatMessage(
      "session-1",
      "Gaming laptop",
      onToken
    );

    expect(response.type).toBe("recommendation");
    expect(onToken).toHaveBeenCalledWith("Catalog match found.");
  });
});
