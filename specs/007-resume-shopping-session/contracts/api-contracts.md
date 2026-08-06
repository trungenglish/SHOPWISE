# API Contracts: Resume Shopping Session

## 1. Resume Session

**Endpoint**: `GET /api/v1/session/resume`
**Description**: Validates a resume token, consumes it, and returns the full session state.

**Request Details**:
- **Headers**: None required (token is in query)
- **Query Params**:
  - `token` (string, required): The JWT resume token.

**Response (Success)**:
- **Status**: 200 OK
- **Body**: The standard `DecisionSession` payload, identical to standard session fetch.

**Response (Error - Invalid/Expired)**:
- **Status**: 401 Unauthorized or 400 Bad Request
- **Body**:
  ```json
  {
    "error": "token_expired",
    "message": "This resume link has expired. Please request a new one."
  }
  ```

**Response (Error - Already Consumed)**:
- **Status**: 409 Conflict
- **Body**:
  ```json
  {
    "error": "token_consumed",
    "message": "This resume link has already been used."
  }
  ```
