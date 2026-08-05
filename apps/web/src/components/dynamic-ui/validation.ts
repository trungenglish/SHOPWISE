import Ajv from "ajv";

// Note: In a real environment, you would load the JSON schemas asynchronously or bundle them.
// For the MVP, we assume the schemas are passed in or bundled.
const ajv = new Ajv({ allErrors: true });

export function validateDynamicUIDocument(payload: any, schema: any): { valid: boolean; errors?: any } {
  const validate = ajv.compile(schema);
  const valid = validate(payload);
  if (!valid) {
    return { valid: false, errors: validate.errors };
  }
  return { valid: true };
}

export function validateInteractionEvent(payload: any, schema: any): { valid: boolean; errors?: any } {
  const validate = ajv.compile(schema);
  const valid = validate(payload);
  if (!valid) {
    return { valid: false, errors: validate.errors };
  }
  return { valid: true };
}

// Optional: VND Currency validation helper function
export function validateVNDCurrency(amount: number): boolean {
  // Principle XVI: All currency must be in VND, integer values without decimals, correctly formatted.
  return Number.isInteger(amount) && amount >= 0;
}
