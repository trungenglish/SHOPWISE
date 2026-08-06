# UI Contracts: Decision Memory

## Components Needed (React + Tailwind v4)

### 1. `SessionHistorySidebar`
- Location: Left drawer or sidebar in the shopping view.
- Responsibility: Fetches `/sessions`, displays active and archived sessions grouped by date (Today, Previous 7 Days, Older).
- Actions: Click to load session, Rename, Archive, Delete.

### 2. `AutoSaveIndicator`
- Location: Top right of the active decision workspace.
- Responsibility: Subscribes to the `useMutation` status of the auto-save API.
- States: `Saving...` (spinner), `Saved to Decision Memory` (checkmark, fades out), `Error saving` (warning icon).

### 3. `PreferenceUndoToast`
- Location: Bottom right toast notification.
- Responsibility: Appears when a preference is auto-extracted.
- Content: "Preference saved: [Value]".
- Actions: "Undo" button (Triggers `DELETE /preferences/:id`).

### 4. `ReasoningReplayModal`
- Location: Overlay when clicking "Why this?" on a recommendation.
- Responsibility: Renders the `reasoning_graph` JSON visually (e.g., as a stepper or bulleted list of constraints applied).

### 5. `BranchSessionButton`
- Location: Context menu on an old message in the conversation thread.
- Responsibility: Triggers `/sessions/:id/branch`. Upon success, updates URL to the new session ID.
