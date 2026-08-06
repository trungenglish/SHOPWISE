import type { SignInFormData } from "../schema";

export const submitSignIn = async (data: SignInFormData): Promise<void> => {
  const res = await fetch("/api/v1/identity/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email: data.emailOrPhone,
      password: data.password,
    }),
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => null);
    throw new Error(errorData?.title || "Failed to sign in");
  }

  const result = await res.json();
  localStorage.setItem("accessToken", result.accessToken);
  if (result.refreshToken) {
    localStorage.setItem("refreshToken", result.refreshToken);
  }
};
