# Quickstart Validation: Resume Shopping Session

## Prerequisites
- Backend API running locally.
- A valid `DecisionSession` ID.

## Validation Scenarios

### Scenario 1: Generate and Use a Valid Token
1. **Generate Token (Mock Trigger)**
   Create a test script or use an internal debug endpoint to trigger the generation of a resume token for your session.
   *Result:* You should receive a JWT token and a log entry indicating Zalo notification "Accepted" (mocked).
2. **Resume Session**
   Navigate to the frontend at `http://localhost:3000/resume?token=<YOUR_TOKEN>`.
   *Result:* The UI should display a loading spinner, successfully call the backend, and restore the shopping UI with your previous items and chat history.

### Scenario 2: Token Replay Prevention
1. Use the same `<YOUR_TOKEN>` from Scenario 1 and navigate to `http://localhost:3000/resume?token=<YOUR_TOKEN>` again.
2. *Result:* The frontend should display the "Invalid or Already Used Link" error screen.

### Scenario 3: Expired Token
1. Generate a token but manipulate the database to set `expires_at` to a past timestamp.
2. Navigate to the resume route.
3. *Result:* The frontend should display the "Expired Link" error screen and offer to start a new session.
