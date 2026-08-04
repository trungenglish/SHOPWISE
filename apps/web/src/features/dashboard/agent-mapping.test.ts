import { describe, expect, it } from "vitest";

import { mapRecommendationProduct } from "./agent-mapping";

describe("mapRecommendationProduct", () => {
  it("maps authoritative catalog fields without inventing performance metrics", () => {
    const product = mapRecommendationProduct({
      id: "product-1",
      image: "https://example.com/laptop.png",
      match_explanation: "Fits gaming and budget requirements.",
      match_score: 92,
      name: "Catalog Laptop",
      price: 37_475_000,
      specifications: {
        cpu: "Core Ultra 9",
        gpu: "RTX 5070",
        ram_gb: 32,
        refresh_rate_hz: 240,
        resolution: "2.5K",
      },
    });

    expect(product.price).toBe(37_475_000);
    expect(product.specs.ram).toBe("32GB");
    expect(product.specs.screen).toBe("2.5K, 240Hz");
    expect(product.aiPerf).toBeUndefined();
    expect(product.rendering).toBeUndefined();
    expect(product.thermals).toBeUndefined();
  });
});
