export interface UserResponse {
  id: string;
  email: string;
  name: string;
  phone?: string;
  created_at: string;
  updated_at: string;
}

export interface UpsertGuestRequest {
  email: string;
  name: string;
  phone: string;
}

export async function upsertGuestUser(data: UpsertGuestRequest): Promise<UserResponse> {
  const response = await fetch("/api/v1/users/guest", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    let errorMessage = "Failed to create guest profile";
    try {
      const errorData = await response.json();
      errorMessage = errorData.detail || errorData.title || errorMessage;
    } catch (e) {
      // Ignore JSON parse error
    }
    throw new Error(errorMessage);
  }

  return response.json();
}
