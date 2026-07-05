import { z } from "zod";

export const RegisterRequestSchema = z.object({
  email: z.email(),
  name: z.string().min(1),
  password: z.string().min(8),
});

export const LoginRequestSchema = z.object({
  email: z.email(),
  password: z.string().min(1),
});

export const RefreshRequestSchema = z.object({
  refreshToken: z.string().min(1),
});

export const TokenResponseSchema = z.object({
  accessToken: z.string(),
  expiresIn: z.number().int().positive(),
  refreshToken: z.string(),
});

export const VerifyEmailRequestSchema = z.object({
  token: z.string().min(1),
});

export type RegisterRequest = z.infer<typeof RegisterRequestSchema>;
export type LoginRequest = z.infer<typeof LoginRequestSchema>;
export type RefreshRequest = z.infer<typeof RefreshRequestSchema>;
export type TokenResponse = z.infer<typeof TokenResponseSchema>;
export type VerifyEmailRequest = z.infer<typeof VerifyEmailRequestSchema>;
