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

const API_BASE = '/api/v1/sessions';

function getHeaders(extraHeaders?: Record<string, string>): HeadersInit {
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
  };
  const token = localStorage.getItem('token');
  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }
  let anonId = localStorage.getItem('anonymousId');
  if (!anonId) {
    anonId = crypto.randomUUID();
    localStorage.setItem('anonymousId', anonId);
  }
  headers['X-Anonymous-ID'] = anonId;
  
  if (extraHeaders) {
    Object.assign(headers, extraHeaders);
  }
  return headers;
}

export async function createSession(initialMessage: string): Promise<DecisionSession> {
  const res = await fetch(API_BASE, {
    method: 'POST',
    headers: getHeaders(),
    body: JSON.stringify({ initial_message: initialMessage }),
  });
  if (!res.ok) throw new Error('Failed to create session');
  return res.json();
}

export async function listSessions(limit = 20, offset = 0): Promise<{ data: DecisionSession[], meta: any }> {
  const res = await fetch(`${API_BASE}?limit=${limit}&offset=${offset}`, {
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error('Failed to list sessions');
  return res.json();
}

export async function getSession(id: string): Promise<DecisionSession> {
  const res = await fetch(`${API_BASE}/${id}`, {
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error('Failed to get session');
  return res.json();
}

export async function renameSession(id: string, title: string): Promise<void> {
  const res = await fetch(`${API_BASE}/${id}`, {
    method: 'PATCH',
    headers: getHeaders(),
    body: JSON.stringify({ title }),
  });
  if (!res.ok) throw new Error('Failed to rename session');
}

export async function deleteSession(id: string): Promise<void> {
  const res = await fetch(`${API_BASE}/${id}`, {
    method: 'DELETE',
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error('Failed to delete session');
}

export async function branchSession(id: string, newTitle: string): Promise<DecisionSession> {
  const res = await fetch(`${API_BASE}/${id}/branch`, {
    method: 'POST',
    headers: getHeaders(),
    body: JSON.stringify({ title: newTitle }),
  });
  if (!res.ok) throw new Error('Failed to branch session');
  return res.json();
}

export async function restoreSession(id: string): Promise<void> {
  const res = await fetch(`${API_BASE}/${id}/restore`, {
    method: 'POST',
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error('Failed to restore session');
}

export async function autoSaveSession(id: string, clientTimestamp: string, pinnedProducts: string[]): Promise<DecisionSession> {
  const res = await fetch(`${API_BASE}/${id}`, {
    method: 'PUT',
    headers: getHeaders({ 'X-Client-Timestamp': clientTimestamp }),
    body: JSON.stringify({ pinned_products: pinnedProducts }),
  });
  if (!res.ok) {
    if (res.status === 409) {
      throw new Error('conflict');
    }
    throw new Error('Failed to autosave session');
  }
  return res.json();
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
  const res = await fetch('/api/v1/sessions/preferences', {
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error('Failed to list preferences');
  return res.json();
}

export async function updatePreference(id: string, category: string, value: string): Promise<UserPreference> {
  const res = await fetch(`/api/v1/sessions/preferences/${id}`, {
    method: 'PUT',
    headers: getHeaders(),
    body: JSON.stringify({ category, value }),
  });
  if (!res.ok) throw new Error('Failed to update preference');
  return res.json();
}

export async function deletePreference(id: string): Promise<void> {
  const res = await fetch(`/api/v1/sessions/preferences/${id}`, {
    method: 'DELETE',
    headers: getHeaders(),
  });
  if (!res.ok) throw new Error('Failed to delete preference');
}

export class ResumeSessionError extends Error {
  code: string;
  constructor(message: string, code: string) {
    super(message);
    this.name = 'ResumeSessionError';
    this.code = code;
  }
}

export async function resumeSessionFromToken(token: string): Promise<DecisionSession> {
  const res = await fetch(`/api/v1/session/resume?token=${encodeURIComponent(token)}`, {
    method: 'GET',
    headers: getHeaders(),
  });
  if (!res.ok) {
    const errorData = await res.json().catch(() => ({}));
    throw new ResumeSessionError(
      errorData.message || 'Failed to resume session',
      errorData.error || `http_error_${res.status}`
    );
  }
  return res.json();
}
