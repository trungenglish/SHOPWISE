import { useState } from "react";
import { useFormContext } from "react-hook-form";
import { useNavigate } from "@tanstack/react-router";
import { ArrowLeft, ArrowRight, CheckCircle2, Loader2 } from "lucide-react";

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

import type { OnboardingFormData } from "../schema";
import { upsertGuestUser } from "../../../api/users";

interface OnboardingWizardStepProps {
  onBack: () => void;
}

export function OnboardingWizardStep({ onBack }: OnboardingWizardStepProps) {
  const navigate = useNavigate();
  const [step, setStep] = useState(1);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const totalSteps = 4; // 1: Name, 2: Phone, 3: Email, 4: Review

  const {
    register,
    handleSubmit,
    trigger,
    watch,
    formState: { errors },
  } = useFormContext<OnboardingFormData>();

  const nextStep = async () => {
    let isValid = false;
    const fields: (keyof OnboardingFormData)[] = [
      "fullName",
      "phoneNumber",
      "email",
    ];

    if (step <= fields.length) {
      isValid = await trigger(fields[step - 1]);
    }

    if (isValid) {
      setStep((prev) => Math.min(prev + 1, totalSteps));
    }
  };

  const prevStep = () => {
    if (step === 1) {
      onBack();
    } else {
      setStep((prev) => Math.max(prev - 1, 1));
    }
  };

  const onSubmit = async (data: OnboardingFormData) => {
    try {
      setIsSubmitting(true);
      const guestUser = await upsertGuestUser({
        email: data.email,
        name: data.fullName,
        phone: data.phoneNumber,
      });
      localStorage.setItem("guestUser", JSON.stringify(guestUser));
      toast.success("Profile created successfully!");
      // Navigate directly to dashboard
      // Note: We use replace: true to prevent navigating back to the onboarding step
      navigate({ to: "/dashboard", replace: true });
    } catch (error) {
      toast.error(error instanceof Error ? error.message : "Failed to create profile. Please try again.");
    } finally {
      setIsSubmitting(false);
    }
  };

  const formValues = watch();

  return (
    <Card className="animate-in fade-in slide-in-from-right-4 w-full duration-200 border-outline-variant/20 bg-surface-low shadow-xl">
      <CardHeader className="p-8 pb-6">
        <div className="mb-2 flex items-center justify-between">
          <Button
            variant="ghost"
            className="text-muted-foreground hover:text-foreground -ml-4 w-fit px-4"
            onClick={prevStep}
            disabled={isSubmitting}
          >
            <ArrowLeft className="mr-2 h-4 w-4" />
            Back
          </Button>
          <div className="flex items-center gap-1.5">
            {[1, 2, 3, 4].map((i) => (
              <div
                key={i}
                className={`h-1.5 w-1.5 rounded-full transition-colors ${
                  step === i ? "bg-primary" : "bg-primary/20"
                }`}
              />
            ))}
          </div>
        </div>
        <CardTitle className="text-foreground text-2xl font-bold tracking-tight">
          {step === 1 && "What's your name?"}
          {step === 2 && "Your phone number"}
          {step === 3 && "Your email address"}
          {step === 4 && "Review your profile"}
        </CardTitle>
        <CardDescription className="text-muted-foreground mt-1">
          {step === 1 && "We'll use this to personalize your experience."}
          {step === 2 && "Required for delivery and order updates."}
          {step === 3 && "For receipts and important communications."}
          {step === 4 && "Make sure everything looks correct."}
        </CardDescription>
      </CardHeader>

      <form
        onSubmit={(e) => {
          if (step < totalSteps) {
            e.preventDefault();
            nextStep();
          } else {
            handleSubmit(onSubmit)(e);
          }
        }}
      >
        <CardContent className="flex flex-col gap-5 px-8 pb-6">
          <div
            className={
              step === 1 ? "animate-in fade-in slide-in-from-right-2 duration-200 block" : "hidden"
            }
          >
            <Field data-invalid={!!errors.fullName}>
              <FieldLabel htmlFor="fullName">Full Name</FieldLabel>
              <FieldContent>
                <Input
                  id="fullName"
                  placeholder="Ex: Nguyen Van A"
                  autoComplete="name"
                  disabled={isSubmitting}
                  {...register("fullName")}
                  aria-invalid={!!errors.fullName}
                  className="h-12"
                />
                <FieldError errors={[errors.fullName]} />
              </FieldContent>
            </Field>
          </div>

          <div
            className={
              step === 2 ? "animate-in fade-in slide-in-from-right-2 duration-200 block" : "hidden"
            }
          >
            <Field data-invalid={!!errors.phoneNumber}>
              <FieldLabel htmlFor="phoneNumber">Phone Number</FieldLabel>
              <FieldContent>
                <Input
                  id="phoneNumber"
                  placeholder="Ex: 0912345678"
                  type="tel"
                  autoComplete="tel"
                  disabled={isSubmitting}
                  {...register("phoneNumber")}
                  aria-invalid={!!errors.phoneNumber}
                  className="h-12"
                />
                <FieldError errors={[errors.phoneNumber]} />
              </FieldContent>
            </Field>
          </div>

          <div
            className={
              step === 3 ? "animate-in fade-in slide-in-from-right-2 duration-200 block" : "hidden"
            }
          >
            <Field data-invalid={!!errors.email}>
              <FieldLabel htmlFor="email">Email Address</FieldLabel>
              <FieldContent>
                <Input
                  id="email"
                  type="email"
                  placeholder="Ex: you@example.com"
                  autoComplete="email"
                  disabled={isSubmitting}
                  {...register("email")}
                  aria-invalid={!!errors.email}
                  className="h-12"
                />
                <FieldError errors={[errors.email]} />
              </FieldContent>
            </Field>
          </div>

          <div
            className={
              step === 4 ? "animate-in fade-in slide-in-from-right-2 duration-200 block" : "hidden"
            }
          >
            <div className="flex flex-col text-sm border-y border-outline-variant/30 py-2">
              <div className="flex justify-between py-3 border-b border-outline-variant/10 last:border-0">
                <span className="text-muted-foreground">Name</span>
                <span className="text-foreground font-medium">
                  {formValues.fullName || "-"}
                </span>
              </div>
              <div className="flex justify-between py-3 border-b border-outline-variant/10 last:border-0">
                <span className="text-muted-foreground">Phone</span>
                <span className="text-foreground font-medium">
                  {formValues.phoneNumber || "-"}
                </span>
              </div>
              <div className="flex justify-between py-3 border-b border-outline-variant/10 last:border-0">
                <span className="text-muted-foreground">Email</span>
                <span className="text-foreground font-medium">
                  {formValues.email || "-"}
                </span>
              </div>
            </div>
          </div>
        </CardContent>

        <CardFooter className="px-8 pb-8">
          {step < totalSteps ? (
            <Button
              type="button"
              className="w-full h-12"
              onClick={nextStep}
            >
              Continue
              <ArrowRight className="ml-2 h-4 w-4" />
            </Button>
          ) : (
            <div className="flex w-full flex-col gap-2">
              <Button
                type="submit"
                className="w-full h-12"
                disabled={isSubmitting}
              >
                {isSubmitting ? (
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                ) : (
                  <CheckCircle2 className="mr-2 h-4 w-4" />
                )}
                <span aria-live="polite">
                  {isSubmitting ? "Creating profile..." : "Create Profile"}
                </span>
              </Button>
              <Button
                type="button"
                variant="outline"
                className="w-full h-12"
                disabled={isSubmitting}
                onClick={() => setStep(3)}
              >
                Edit Information
              </Button>
            </div>
          )}
        </CardFooter>
      </form>
    </Card>
  );
}
