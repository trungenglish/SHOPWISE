import { readFileSync } from "node:fs";

import Ajv2020 from "ajv/dist/2020.js";
import type { AnySchema } from "ajv";
import { expect, it } from "vitest";

import { dynamicUIComponentTypes } from "@shopwise/protocols";

const readJson = (relativePath: string): unknown =>
  JSON.parse(
    readFileSync(new URL(relativePath, import.meta.url), "utf8")
  ) as unknown;

const schema = readJson(
  "../../../../packages/schemas/dynamic-ui-v1.schema.json"
) as AnySchema;
const fixture = readJson(
  "../../../../packages/schemas/fixtures/dynamic-ui-all-components.json"
);
const validator = new Ajv2020({
  formats: {
    uuid: /^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i,
  },
  strict: true,
}).compile(schema);

const componentTypes = (value: unknown): string[] => {
  if (typeof value !== "object" || value === null) {
    return [];
  }
  const record = value as Record<string, unknown>;
  const ownType = typeof record.type === "string" ? [record.type] : [];
  const children = Array.isArray(record.children) ? record.children : [];
  return [...ownType, ...children.flatMap(componentTypes)];
};

it("validates the shared fixture and covers every component type", () => {
  expect(validator(fixture), JSON.stringify(validator.errors)).toBe(true);
  const operations = (fixture as { ui_operations: { component: unknown }[] })
    .ui_operations;
  const actualTypes = new Set(
    operations.flatMap((operation) => componentTypes(operation.component))
  );
  expect(actualTypes).toEqual(new Set(dynamicUIComponentTypes));
});

it("rejects unsafe props and invalid interaction actions", () => {
  const unsafe = structuredClone(fixture) as {
    ui_operations: { component: { children: Record<string, unknown>[] } }[];
  };
  unsafe.ui_operations[0].component.children[0].props = {
    dangerouslySetInnerHTML: "<script>alert(1)</script>",
  };
  expect(validator(unsafe)).toBe(false);

  const invalidAction = structuredClone(fixture) as {
    ui_operations: {
      component: { children: { interaction?: { action: string } }[] };
    }[];
  };
  const button = invalidAction.ui_operations[0].component.children.find(
    (component) => component.interaction
  );
  if (!button?.interaction) {
    throw new Error("fixture must contain an interactive component");
  }
  button.interaction.action = "javascript.run";
  expect(validator(invalidAction)).toBe(false);
});
