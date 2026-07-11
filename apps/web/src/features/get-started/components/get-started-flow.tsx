import { useState } from "react";
import { useForm, FormProvider } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { AccountStatusStep } from "./account-status-step";
import { SignInStep } from "./sign-in-step";
import { OnboardingWizardStep } from "./onboarding-wizard-step";
import { onboardingSchema, type OnboardingFormData } from "../schema";

export type FlowState = "choice" | "signin" | "onboarding";

export function GetStartedFlow() {
  const [flowState, setFlowState] = useState<FlowState>("choice");

  const methods = useForm<OnboardingFormData>({
    resolver: zodResolver(onboardingSchema),
    defaultValues: {
      fullName: "",
      phoneNumber: "",
      email: "",
    },
    mode: "onTouched",
  });

  return (
    <div className="bg-background flex min-h-svh flex-col items-center justify-start p-6 pt-12 md:justify-center md:p-10">
      <div className="w-full max-w-sm md:max-w-md">
        {flowState === "choice" && (
          <AccountStatusStep onSelect={setFlowState} />
        )}
        {flowState === "signin" && (
          <SignInStep onBack={() => setFlowState("choice")} />
        )}
        {flowState === "onboarding" && (
          <FormProvider {...methods}>
            <OnboardingWizardStep onBack={() => setFlowState("choice")} />
          </FormProvider>
        )}
      </div>
    </div>
  );
}
