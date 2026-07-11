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
      // @ts-expect-error /dashboard route is in another branch
      navigate({ to: "/dashboard" });
    } catch (error) {
      toast.error("Failed to sign in. Please try again.");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Card className="animate-in slide-in-from-right-4 fade-in w-full duration-300">
      <CardHeader>
        <div className="mb-2 flex items-center gap-2">
          <Button
            variant="ghost"
            size="icon"
            className="text-muted-foreground hover:text-foreground -ml-2 min-h-[44px] min-w-[44px]"
            onClick={onBack}
            disabled={isSubmitting}
            aria-label="Go back"
          >
            <ArrowLeft className="h-4 w-4" />
          </Button>
          <CardTitle className="text-xl font-bold">Sign In</CardTitle>
        </div>
        <CardDescription>
          Enter your email or phone number to access your account.
        </CardDescription>
      </CardHeader>
      <form onSubmit={handleSubmit(onSubmit)}>
        <CardContent className="flex flex-col gap-4">
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
              />
              <FieldError errors={[errors.password]} />
            </FieldContent>
          </Field>
        </CardContent>
        <CardFooter>
          <Button
            type="submit"
            className="w-full"
            size="lg"
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
