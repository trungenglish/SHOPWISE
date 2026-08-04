import React, { createContext, useContext, useState, useEffect, useCallback, ReactNode } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { fetchPromotionStatus, startPromotion } from "../api/queries";
import { PromotionStatusResponse, TriggerType } from "@shopwise/api-types/promotions";

type DisplayState = "HIDDEN" | "ACTIVE" | "EXPIRED" | "SUCCESS" | "ERROR" | "RETRY";

interface CheckoutIncentiveContextValue {
  currentPromotion: PromotionStatusResponse | null;
  displayState: DisplayState;
  remainingSeconds: number;
  isCollapsed: boolean;
  startPromotion: (triggerType: TriggerType) => void;
  refreshPromotion: () => void;
  retryVoucherIssuance: () => void;
  collapse: () => void;
  expand: () => void;
}

const CheckoutIncentiveContext = createContext<CheckoutIncentiveContextValue | null>(null);

export const CheckoutIncentiveProvider = ({ children }: { children: ReactNode }) => {
  const queryClient = useQueryClient();
  const [isCollapsed, setIsCollapsed] = useState(false);
  const [remainingSeconds, setRemainingSeconds] = useState(0);

  const { data: currentPromotion, error, refetch, isError } = useQuery<PromotionStatusResponse>({
    queryKey: ["checkout-incentive-status"],
    queryFn: fetchPromotionStatus,
    retry: false,
    refetchOnWindowFocus: true,
  });

  const mutation = useMutation({
    mutationFn: startPromotion,
    onSuccess: (data) => {
      queryClient.setQueryData(["checkout-incentive-status"], data);
    },
    onError: (err) => {
      // If error is conflict (409), we ignore because they are ineligible or on cooldown
      console.error(err);
    }
  });

  const triggerPromotion = useCallback((triggerType: TriggerType) => {
    // Only trigger if we don't already have an active promotion
    if (!currentPromotion || currentPromotion.status === "EXPIRED") {
      mutation.mutate({ trigger_type: triggerType });
    }
  }, [currentPromotion, mutation]);

  useEffect(() => {
    if (!currentPromotion?.expires_at || currentPromotion.status !== "ACTIVE") {
      setRemainingSeconds(0);
      return;
    }

    const expiresAt = new Date(currentPromotion.expires_at).getTime();
    
    const updateCountdown = () => {
      const now = Date.now();
      const diff = Math.max(0, Math.floor((expiresAt - now) / 1000));
      setRemainingSeconds(diff);
      
      if (diff === 0 && currentPromotion.status === "ACTIVE") {
        // Trigger a refetch to get the EXPIRED state from backend
        refetch();
      }
    };

    updateCountdown();
    const interval = setInterval(updateCountdown, 1000);
    return () => clearInterval(interval);
  }, [currentPromotion, refetch]);

  let displayState: DisplayState = "HIDDEN";
  if (currentPromotion) {
    switch (currentPromotion.status) {
      case "ACTIVE":
      case "COMPLETED_PENDING_VOUCHER":
        displayState = remainingSeconds > 0 || currentPromotion.status === "COMPLETED_PENDING_VOUCHER" ? "ACTIVE" : "EXPIRED";
        break;
      case "EXPIRED":
        displayState = "EXPIRED";
        break;
      case "VOUCHER_ISSUED":
        displayState = "SUCCESS";
        break;
      case "ISSUE_RETRY_REQUIRED":
        displayState = "RETRY";
        break;
      case "ERROR_RECOVERABLE":
        displayState = "ERROR";
        break;
    }
  }

  // If there's an API error other than 404/not_found
  if (isError && error.message !== "not_found") {
    displayState = "ERROR";
  }

  // Auto-hide EXPIRED after a few seconds? The requirements say "Expired state: This promotional window has ended." 
  // We'll leave it up to the panel to display the expired state.

  const value: CheckoutIncentiveContextValue = {
    currentPromotion: currentPromotion || null,
    displayState,
    remainingSeconds,
    isCollapsed,
    startPromotion: triggerPromotion,
    refreshPromotion: () => { refetch(); },
    retryVoucherIssuance: () => { refetch(); /* Backend retry endpoint not specified, we'll just refetch */ },
    collapse: () => setIsCollapsed(true),
    expand: () => setIsCollapsed(false),
  };

  return (
    <CheckoutIncentiveContext.Provider value={value}>
      {children}
    </CheckoutIncentiveContext.Provider>
  );
};

export const useCheckoutIncentive = () => {
  const context = useContext(CheckoutIncentiveContext);
  if (!context) {
    throw new Error("useCheckoutIncentive must be used within CheckoutIncentiveProvider");
  }
  return context;
};
