import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { StartPromotionRequest, PromotionStatusResponse } from "@shopwise/api-types/promotions";

const BASE_URL = "/api/v1/checkout-incentive";

function getHeaders(): HeadersInit {
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
  // Use anonymousId as device/session identifier for anonymous users
  headers["X-Session-ID"] = anonId;
  headers["X-Device-ID"] = anonId;
  
  return headers;
}

export async function fetchPromotionStatus(): Promise<PromotionStatusResponse> {
  const res = await fetch(`${BASE_URL}/status`, {
    headers: getHeaders(),
  });
  if (!res.ok) {
    if (res.status === 404) {
      // 404 means no active promotion for this session
      throw new Error("not_found");
    }
    throw new Error("Failed to fetch promotion status");
  }
  return res.json();
}

export async function startPromotion(req: StartPromotionRequest): Promise<PromotionStatusResponse> {
  const res = await fetch(`${BASE_URL}/start`, {
    method: "POST",
    headers: getHeaders(),
    body: JSON.stringify(req),
  });
  if (!res.ok) {
    if (res.status === 409) {
      throw new Error("conflict"); // Cooldown or ineligible
    }
    throw new Error("Failed to start promotion");
  }
  return res.json();
}

export async function simulatePayment(): Promise<void> {
  let anonId = localStorage.getItem("anonymousId") || "";
  const res = await fetch(`${BASE_URL}/dev/simulate-payment`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ session_id: anonId }),
  });
  if (!res.ok) {
    throw new Error("Failed to simulate payment");
  }
}
