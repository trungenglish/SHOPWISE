import { z } from "zod/v3";

export const signInSchema = z.object({
  emailOrPhone: z.string().min(1, "Email or phone number is required"),
  password: z.string().min(1, "Password is required"),
});

export type SignInFormData = z.infer<typeof signInSchema>;

export const onboardingSchema = z.object({
  fullName: z.string().min(1, "Full name is required"),
  phoneNumber: z
    .string()
    .trim()
    .min(1, "Phone number is required")
    .regex(
      /^(03|05|07|08|09)\d{8}$/,
      "Invalid Vietnamese phone number. Must be 10 digits and start with 03, 05, 07, 08, or 09."
    ),
  email: z.string().min(1, "Email is required").email("Invalid email address"),
});

export type OnboardingFormData = z.infer<typeof onboardingSchema>;
