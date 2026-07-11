import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "@tanstack/react-router";
import { ArrowLeft, Loader2, LogIn } from "lucide-react";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
  CardFooter,
} from "@shopwise/ui/components/card";
import { Button } from "@shopwise/ui/components/button";
import {
  Field,
  FieldLabel,
  FieldError,
  FieldContent,
} from "@shopwise/ui/components/field";
import { Input } from "@shopwise/ui/components/input";
import { toast } from "sonner";

import { signInSchema, type SignInFormData } from "../schema";
import { mockSubmitSignIn } from "../services/mock-adapter";

interface SignInStepProps {
  onBack: () => void;
}

export function SignInStep({ onBack }: SignInStepProps) {
  const navigate = useNavigate();
  const [isSubmitting, setIsSubmitting] = useState(false);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignInFormData>({
    resolver: zodResolver(signInSchema),
    defaultValues: {
      emailOrPhone: "",
      password: "",
    },
  });

  const onSubmit = async (data: SignInFormData) => {
    try {
      setIsSubmitting(true);
      await mockSubmitSignIn(data);
      toast.success("Signed in successfully!");
      // MOCKED FLOW: Bypass real authentication and navigate directly to dashboard
      // Note: We use replace: true to prevent navigating back to the sign-in step
      navigate({ to: "/dashboard", replace: true });
    } catch (error) {
      toast.error("Failed to sign in. Please try again.");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Card className="animate-in fade-in slide-in-from-right-4 w-full duration-200 border-outline-variant/20 bg-surface-low shadow-xl">
      <CardHeader className="p-8 pb-6">
        <Button
          variant="ghost"
          className="text-muted-foreground hover:text-foreground -ml-4 mb-2 w-fit px-4"
          onClick={onBack}
          disabled={isSubmitting}
        >
          <ArrowLeft className="mr-2 h-4 w-4" />
          Back
        </Button>
        <CardTitle className="text-foreground text-2xl font-bold tracking-tight">Sign In</CardTitle>
        <CardDescription className="text-muted-foreground mt-1">
          Enter your email or phone number to access your account.
        </CardDescription>
      </CardHeader>
      <form onSubmit={handleSubmit(onSubmit)}>
        <CardContent className="flex flex-col gap-5 px-8 pb-6">
          <Field data-invalid={!!errors.emailOrPhone}>
            <FieldLabel htmlFor="emailOrPhone">
              Email or Phone Number
            </FieldLabel>
            <FieldContent>
              <Input
                id="emailOrPhone"
                placeholder="Ex: 0912345678 or you@example.com"
                autoComplete="username"
                disabled={isSubmitting}
                {...register("emailOrPhone")}
                aria-invalid={!!errors.emailOrPhone}
                className="h-12"
              />
              <FieldError errors={[errors.emailOrPhone]} />
            </FieldContent>
          </Field>

          <Field data-invalid={!!errors.password}>
            <FieldLabel htmlFor="password">Password</FieldLabel>
            <FieldContent>
              <Input
                id="password"
                type="password"
                autoComplete="current-password"
                disabled={isSubmitting}
                {...register("password")}
                aria-invalid={!!errors.password}
                className="h-12"
              />
              <FieldError errors={[errors.password]} />
            </FieldContent>
          </Field>
        </CardContent>
        <CardFooter className="px-8 pb-8">
          <Button
            type="submit"
            className="w-full h-12"
            disabled={isSubmitting}
          >
            {isSubmitting ? (
              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
            ) : (
              <LogIn className="mr-2 h-4 w-4" />
            )}
            {isSubmitting ? "Signing in..." : "Sign in to ShopWise"}
          </Button>
        </CardFooter>
      </form>
    </Card>
  );
}
