import { render, screen } from "@testing-library/react";
import { expect, it } from "vitest";

import type { Laptop } from "../types";
import ComparisonModal from "./comparison-modal";

const product: Laptop = {
  id: "product-1",
  image: "https://example.com/laptop.png",
  matchExplanation: "Fits the request.",
  matchScore: 90,
  name: "Laptop 1",
  price: 37_475_000,
  specs: {
    cooling: "Dual fan",
    cpu: "Core Ultra 9",
    gpu: "RTX 5070",
    ram: "32GB",
    screen: "2.5K, 240Hz",
    warranty: "2 years",
  },
};

it("keeps the metric label and four products in the same grid row", () => {
  const products = ["1", "2", "3", "4"].map((id) => ({
    ...product,
    id: `product-${id}`,
    name: `Laptop ${id}`,
  }));

  render(
    <ComparisonModal isOpen onClose={() => undefined} products={products} />
  );

  expect(screen.getByText("Comparison Metrics").parentElement).toHaveStyle({
    gridTemplateColumns: "minmax(180px, 1fr) repeat(4, minmax(200px, 1fr))",
  });
});
