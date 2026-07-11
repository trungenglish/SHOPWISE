import { useState } from "react";
import { useFormContext } from "react-hook-form";
import { useNavigate } from "@tanstack/react-router";
import {
  ArrowLeft,
  ArrowRight,
  CheckCircle2,
  Loader2,
} from "lucide-react";

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
import { mockSubmitOnboarding } from "../services/mock-adapter";

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
    getValues,
    formState: { errors },
  } = useFormContext<OnboardingFormData>();

  const nextStep = async () => {
    let isValid = false;
    const fields: (keyof OnboardingFormData)[] = ["fullName", "phoneNumber", "email"];
    
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
      await mockSubmitOnboarding(data);
      toast.success("Profile created successfully!");
      // @ts-expect-error /dashboard route is in another branch
      navigate({ to: "/dashboard" });
    } catch (error) {
      toast.error("Failed to create profile. Please try again.");
    } finally {
      setIsSubmitting(false);
    }
  };

  const formValues = getValues();

  return (
    <Card className="animate-in slide-in-from-right-4 fade-in w-full duration-300">
      <CardHeader>
        <div className="mb-2 flex items-center justify-between">
          <div className="flex items-center gap-2">
            <Button
              variant="ghost"
              size="icon"
              className="text-muted-foreground hover:text-foreground -ml-2 min-h-[44px] min-w-[44px]"
              onClick={prevStep}
              disabled={isSubmitting}
              aria-label="Go back"
            >
              <ArrowLeft className="h-4 w-4" />
            </Button>
            <CardTitle className="text-xl font-bold">
              {step === 1 && "What's your name?"}
              {step === 2 && "Your phone number"}
              {step === 3 && "Your email address"}
              {step === 4 && "Review your profile"}
            </CardTitle>
          </div>
          <div className="bg-primary/10 text-primary flex h-8 w-8 items-center justify-center rounded-full text-xs font-bold">
            {step}/{totalSteps}
          </div>
        </div>
        <CardDescription>
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
        <CardContent className="flex flex-col gap-4">
          <div
            className={
              step === 1 ? "animate-in fade-in zoom-in-95 block" : "hidden"
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
                />
                <FieldError errors={[errors.fullName]} />
              </FieldContent>
            </Field>
          </div>

          <div
            className={
              step === 2 ? "animate-in fade-in zoom-in-95 block" : "hidden"
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
                />
                <FieldError errors={[errors.phoneNumber]} />
              </FieldContent>
            </Field>
          </div>

          <div
            className={
              step === 3 ? "animate-in fade-in zoom-in-95 block" : "hidden"
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
                />
                <FieldError errors={[errors.email]} />
              </FieldContent>
            </Field>
          </div>

          <div
            className={
              step === 4 ? "animate-in fade-in zoom-in-95 block" : "hidden"
            }
          >
            <div className="bg-muted/50 space-y-4 rounded-lg border p-4 text-sm">
              <div className="flex justify-between border-b pb-2">
                <span className="text-muted-foreground">Full Name</span>
                <span className="text-foreground font-medium">
                  {formValues.fullName || "-"}
                </span>
              </div>
              <div className="flex justify-between border-b pb-2">
                <span className="text-muted-foreground">Phone</span>
                <span className="text-foreground font-medium">
                  {formValues.phoneNumber || "-"}
                </span>
              </div>
              <div className="flex justify-between">
                <span className="text-muted-foreground">Email</span>
                <span className="text-foreground font-medium">
                  {formValues.email || "-"}
                </span>
              </div>
            </div>
          </div>
        </CardContent>

        <CardFooter>
          {step < totalSteps ? (
            <Button
              type="button"
              className="w-full"
              size="lg"
              onClick={nextStep}
            >
              Continue
              <ArrowRight className="ml-2 h-4 w-4" />
            </Button>
          ) : (
            <Button
              type="submit"
              className="w-full"
              size="lg"
              disabled={isSubmitting}
            >
              {isSubmitting ? (
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              ) : (
                <CheckCircle2 className="mr-2 h-4 w-4" />
              )}
              {isSubmitting
                ? "Creating profile..."
                : "Confirm & Create Profile"}
            </Button>
          )}
        </CardFooter>
      </form>
    </Card>
  );
}
