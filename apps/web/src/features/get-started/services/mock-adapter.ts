import type { SignInFormData, OnboardingFormData } from "../schema";

const delay = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

// Random latency between 800ms and 1500ms
const getNetworkLatency = () =>
  Math.floor(Math.random() * (1500 - 800 + 1)) + 800;

export const mockSubmitSignIn = async (data: SignInFormData): Promise<void> => {
  await delay(getNetworkLatency());
  // In a real app, this would validate credentials via an API.
  // For the mock, we just resolve successfully.
  console.log("Mock sign in successful for:", data.emailOrPhone);
};

export const mockSubmitOnboarding = async (
  data: OnboardingFormData
): Promise<void> => {
  await delay(getNetworkLatency());
  // In a real app, this would register the new user via an API.
  // For the mock, we just resolve successfully.
  console.log("Mock onboarding successful for:", data.email);
};
