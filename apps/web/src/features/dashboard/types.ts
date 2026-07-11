export interface Laptop {
  id: string;
  name: string;
  price: number;
  image: string;
  matchScore: number;
  specs: {
    gpu: string;
    ram: string;
    cooling: string;
    cpu?: string;
    screen?: string;
  };
  aiPerf: number;
  rendering: number;
  thermals: number;
  matchExplanation?: string;
}

export interface PriceAlert {
  productId: string;
  targetPrice: number;
  active: boolean;
  createdAt: string;
}

export interface AuditLog {
  time: string;
  message: string;
  status: "done" | "running" | "pending";
}

export interface AgentStatus {
  id: string;
  name: string;
  progress: number;
  statusMessage: string;
  type: "shopping" | "inventory" | "warranty" | "finance";
}

export interface TrustFactors {
  benchmarkSources: number;
  reviewCoverage: number;
  retailConsensus: string;
  confidenceEvolution: number[];
}

export interface Session {
  id: string;
  title: string;
  userIntent: string;
  date: string;
  products: Laptop[];
  logs: AuditLog[];
  agents: AgentStatus[];
  trustScore: number;
  trustFactors: TrustFactors;
}
