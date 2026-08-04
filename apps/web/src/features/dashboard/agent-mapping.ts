import type { AgentEnvelope } from "../../api/decision-memory";

import type { Laptop } from "./types";

type RecommendationProduct = Exclude<
  AgentEnvelope,
  { type: "question" }
>["decision"]["products"][number];

const notProvided = "Not provided";

const displayValue = (value: unknown): string =>
  typeof value === "string" && value.length > 0 ? value : notProvided;

export const mapRecommendationProduct = (
  product: RecommendationProduct
): Laptop => {
  const {
    id,
    image,
    match_explanation: matchExplanation,
    match_score: matchScore,
    name,
    price,
    specifications,
  } = product;
  const memory =
    typeof specifications.ram_gb === "number"
      ? `${specifications.ram_gb}GB`
      : displayValue(specifications.ram);
  const refreshRate =
    typeof specifications.refresh_rate_hz === "number"
      ? `${specifications.refresh_rate_hz}Hz`
      : "";
  const resolution = displayValue(specifications.resolution);
  const screen = [resolution === notProvided ? "" : resolution, refreshRate]
    .filter(Boolean)
    .join(", ");

  return {
    id,
    image,
    matchExplanation,
    matchScore,
    name,
    price,
    specs: {
      cooling: displayValue(specifications.cooling),
      cpu: displayValue(specifications.cpu),
      gpu: displayValue(specifications.gpu),
      ram: memory,
      screen: screen || notProvided,
      warranty: displayValue(specifications.warranty),
    },
  };
};
