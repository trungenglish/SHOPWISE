import { fireEvent, render, screen } from "@testing-library/react";
import { expect, it, vi } from "vitest";

import DynamicUIRenderer from "./dynamic-uirenderer";

it("requires an explicit click before a checkout-ready action", () => {
  const onConfirmCheckout = vi.fn();
  render(
    <DynamicUIRenderer
      envelope={{
        type: "checkout_ready",
        message: "Ready for confirmation.",
        decision: {
          reasoning: "Selected by the user.",
          products: [
            {
              id: "product-1",
              name: "Catalog Laptop",
              price: 37_475_000,
              image: "",
              specifications: {},
              match_score: 90,
              match_explanation: "Confirmed choice.",
            },
          ],
        },
      }}
      onConfirmCheckout={onConfirmCheckout}
    />
  );

  expect(onConfirmCheckout).not.toHaveBeenCalled();
  fireEvent.click(screen.getByRole("button", { name: /confirm checkout/i }));
  expect(onConfirmCheckout).toHaveBeenCalledWith("product-1");
});
