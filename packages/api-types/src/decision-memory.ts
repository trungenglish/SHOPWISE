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

export interface SessionListResponse {
  data: DecisionSession[];
  meta: {
    limit: number;
    offset: number;
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

export interface PreferenceListResponse {
  data: UserPreference[];
}

export interface UpdatePreferenceRequest {
  category: string;
  value: string;
}
