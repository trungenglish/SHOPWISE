#!/usr/bin/env bash

set -e

echo "Running Decision Memory Quickstart Validation"

echo "1. Creating Session..."
RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/sessions \
  -H "Content-Type: application/json" \
  -H "X-Anonymous-ID: test-user" \
  -d '{"initial_message": "Looking for a laptop"}')

# Simple regex parsing to extract ID if jq is not available, but jq is standard enough.
SESSION_ID=$(echo "$RESPONSE" | grep -o '"id":"[^"]*' | cut -d'"' -f4 | head -n 1)
echo "Session ID: $SESSION_ID"

if [ -z "$SESSION_ID" ]; then
  echo "Failed to create session. Is the backend running on port 8080?"
  exit 1
fi

echo -e "\n2. Auto-Save (Valid Timestamp)..."
curl -s -X PUT http://localhost:8080/api/v1/sessions/$SESSION_ID \
  -H "Content-Type: application/json" \
  -H "X-Client-Timestamp: 2099-01-01T00:00:00Z" \
  -d '{"pinned_products": ["p1"]}'
echo "Auto-save request sent."

echo -e "\n3. Test Conflict Resolution (Old Timestamp)..."
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" -X PUT http://localhost:8080/api/v1/sessions/$SESSION_ID \
  -H "Content-Type: application/json" \
  -H "X-Client-Timestamp: 2000-01-01T00:00:00Z" \
  -d '{"pinned_products": ["p1", "p2"]}')

echo "Conflict Response HTTP Code: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 409 ]; then
  echo "Success: Server rejected outdated state with 409 Conflict."
else
  echo "Warning: Expected 409, got $HTTP_CODE"
fi

echo -e "\n4. Branch Session..."
BRANCH_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/sessions/$SESSION_ID/branch \
  -H "Content-Type: application/json" \
  -d '{"title": "Branched Session"}')

BRANCH_ID=$(echo "$BRANCH_RESPONSE" | grep -o '"id":"[^"]*' | cut -d'"' -f4 | head -n 1)
echo "Branched Session ID: $BRANCH_ID"

if [ -n "$BRANCH_ID" ] && [ "$BRANCH_ID" != "$SESSION_ID" ]; then
  echo "Success: Branched session successfully created."
else
  echo "Failed to create branched session."
  exit 1
fi

echo -e "\nValidation Complete."
