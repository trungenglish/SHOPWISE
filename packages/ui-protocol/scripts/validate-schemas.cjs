const fs = require('fs');
const path = require('path');
const Ajv = require('ajv/dist/2020');

const ajv = new Ajv({ allErrors: true, strict: false });


function loadJson(filePath) {
  return JSON.parse(fs.readFileSync(filePath, 'utf8'));
}

const schemaDir = path.join(__dirname, '../src/schema');
const schemas = ['base-schema.json', 'event-schema.json', 'layout-components.json', 'ai-components.json', 'commerce-components.json', 'interaction-components.json', 'error-components.json', 'actions.json'];

console.log("Loading schemas...");
schemas.forEach(file => {
  const schema = loadJson(path.join(schemaDir, file));
  ajv.addSchema(schema, schema.$id);
  console.log(`Loaded ${file} (${schema.$id})`);
});

console.log("\nValidating quickstart sample payload...");

// Simulate a sample payload for validation
const samplePayload = {
  protocol: "dynamic-ui",
  schemaVersion: "1.0",
  documentId: "doc_123",
  timestamp: "2026-07-12T00:00:00Z",
  operation: "replace",
  metadata: {
    conversationId: "conv_456"
  },
  root: {
    id: "root_1",
    type: "ProductCardProps",
    props: {
      productId: "prod_1",
      name: "Super Phone",
      priceVND: 15000000
    }
  }
};

const validate = ajv.compile(ajv.getSchema('https://shopwise.local/schema/base-schema.json').schema);
const valid = validate(samplePayload);

if (!valid) {
  console.error("Validation failed:", validate.errors);
  process.exit(1);
} else {
  console.log("Validation passed successfully!");
}
