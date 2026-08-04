import { readFileSync } from "node:fs";
import { resolve } from "node:path";

import { fireEvent, render, screen } from "@testing-library/react";
import { expect, it, vi } from "vitest";

import type { AgentEnvelope } from "../api/decision-memory";
import DynamicUIRenderer from "./dynamic-uirenderer";

const emptyDecision = {
  reasoning: "Fixture",
  products: [
    {
      id: "fixture-product",
      name: "Fixture product",
      price: 1,
      image: "",
      specifications: {},
      match_score: 1,
      match_explanation: "Fixture",
    },
  ],
};

it("submits a clarification only after the user confirms an answer", () => {
  const onInteraction = vi.fn();
  render(
    <DynamicUIRenderer
      envelope={{
        type: "question",
        message: "What is your budget?",
        schema_version: "1.0",
        turn_id: "turn-1",
        revision: 1,
        conversation_state: "collecting_requirements",
        ui_state: { status: "ready" },
        ui_operations: [],
        question: {
          id: "question-1",
          mode: "single",
          free_text_allowed: true,
          options: [
            { id: "under-25m", label: "Under 25 million VND" },
            { id: "25m-40m", label: "25–40 million VND" },
            { id: "flexible", label: "Flexible budget" },
          ],
        },
      }}
      onInteraction={onInteraction}
    />
  );

  expect(onInteraction).not.toHaveBeenCalled();
  fireEvent.click(screen.getByLabelText("Under 25 million VND"));
  expect(onInteraction).not.toHaveBeenCalled();
  fireEvent.click(screen.getByRole("button", { name: "Submit" }));
  expect(onInteraction).toHaveBeenCalledWith(
    expect.objectContaining({
      interaction_id: "question-1",
      action: "question.answer",
      payload: expect.objectContaining({
        selected_option_ids: ["under-25m"],
      }),
    })
  );
});

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

it("renders the shared fixture and ignores an unsupported future component", () => {
  const fixture = JSON.parse(
    readFileSync(
      resolve(
        process.cwd(),
        "../../packages/schemas/fixtures/dynamic-ui-all-components.json"
      ),
      "utf8"
    )
  ) as Omit<AgentEnvelope, "type" | "message" | "decision">;
  const envelope = {
    ...fixture,
    type: "recommendation",
    message: "Fixture",
    decision: emptyDecision,
    ui_operations: [
      ...(fixture.ui_operations ?? []),
      {
        id: "future-operation",
        sequence: 99,
        operation: "append" as const,
        target: "surface" as const,
        component: {
          id: "future-component",
          type: "future_widget",
          children: [
            {
              id: "known-child",
              type: "chat_message",
              props: { message: "Known fallback" },
            },
          ],
        },
      },
    ],
  } satisfies AgentEnvelope;

  render(<DynamicUIRenderer envelope={envelope} />);

  expect(screen.getByText("Card")).toBeInTheDocument();
  expect(screen.getByText("Known fallback")).toBeInTheDocument();
});
