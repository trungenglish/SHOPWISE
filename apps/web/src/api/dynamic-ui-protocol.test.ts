import { applyUIOperations, type UIOperation } from "@shopwise/protocols";
import { expect, it } from "vitest";

it("applies UI operations deterministically without mutating prior state", () => {
  const original = {};
  const operations: UIOperation[] = [
    {
      id: "b",
      sequence: 2,
      operation: "append",
      target: "surface",
      component: { id: "second", type: "card" },
    },
    {
      id: "a",
      sequence: 1,
      operation: "append",
      target: "surface",
      component: { id: "first", type: "card" },
    },
  ];

  expect(Object.keys(applyUIOperations(original, operations))).toEqual([
    "first",
    "second",
  ]);
  expect(original).toEqual({});
});

it("replaces a nested component after ordering partial updates", () => {
  const result = applyUIOperations({}, [
    {
      id: "replace",
      sequence: 2,
      operation: "replace",
      target: "component",
      target_id: "question",
      component: { id: "answer", type: "chat_message" },
    },
    {
      id: "append",
      sequence: 1,
      operation: "append",
      target: "surface",
      component: {
        id: "page",
        type: "page",
        children: [{ id: "question", type: "card" }],
      },
    },
  ]);

  expect(result.page?.children).toEqual([
    { id: "answer", type: "chat_message" },
  ]);
});
