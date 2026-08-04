import React from "react";
import { useCheckoutIncentive } from "../hooks/useCheckoutIncentive";
import { CountdownDisplay } from "./CountdownDisplay";
import { PromotionBadge } from "./PromotionBadge";
import { 
  VoucherIssuedState, 
  PromotionExpiredState, 
  PromotionErrorState, 
  IssueRetryRequiredState 
} from "./States";
import { X, ChevronDown } from "lucide-react";

export const CheckoutIncentivePanel: React.FC = () => {
  const { 
    displayState, 
    isCollapsed, 
    currentPromotion, 
    remainingSeconds, 
    collapse, 
    expand,
    retryVoucherIssuance
  } = useCheckoutIncentive();

  if (displayState === "HIDDEN" || !currentPromotion) {
    return null;
  }

  const rewardValue = currentPromotion.reward_value ?? 15;

  if (isCollapsed && displayState === "ACTIVE") {
    return (
      <PromotionBadge 
        remainingSeconds={remainingSeconds} 
        rewardValue={rewardValue} 
        onClick={expand} 
      />
    );
  }

  return (
    <div 
      className="fixed bottom-4 right-4 md:top-16 md:bottom-auto md:right-6 z-50 w-[300px] rounded-xl bg-slate-900/95 backdrop-blur-xl shadow-2xl border border-slate-700/50 overflow-hidden transition-all duration-500 ease-out transform origin-bottom-right md:origin-top-right hover:shadow-cyan-500/10"
      role="region"
      aria-label="Checkout Promotion"
    >
      {/* Dynamic gradient top border */}
      <div className="absolute top-0 left-0 w-full h-1 bg-gradient-to-r from-cyan-400 via-blue-500 to-purple-500" />
      
      {/* Header section with collapse/close controls */}
      <div className="px-4 py-3 border-b border-slate-800/60 flex justify-between items-center">
        <h2 className="font-semibold text-slate-100 flex items-center gap-2 text-sm">
          <span className="relative flex h-2 w-2">
            <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-cyan-400 opacity-75"></span>
            <span className="relative inline-flex rounded-full h-2 w-2 bg-cyan-500"></span>
          </span>
          Special Offer
        </h2>
        <div className="flex items-center gap-1">
          {displayState === "ACTIVE" && (
            <button 
              onClick={collapse}
              className="p-1 rounded-md text-slate-400 hover:bg-slate-800 hover:text-white transition-colors focus:outline-none focus:ring-2 focus:ring-cyan-500"
              aria-label="Collapse panel"
            >
              <ChevronDown className="w-3.5 h-3.5" />
            </button>
          )}
          {displayState !== "ACTIVE" && (
            <button 
              onClick={collapse}
              className="p-1 rounded-md text-slate-400 hover:bg-slate-800 hover:text-white transition-colors focus:outline-none focus:ring-2 focus:ring-cyan-500"
              aria-label="Close panel"
            >
              <X className="w-3.5 h-3.5" />
            </button>
          )}
        </div>
      </div>

      {/* Main Content */}
      <div className="p-4">
        {displayState === "ACTIVE" && (
          <div className="flex flex-col space-y-3">
            <p className="text-slate-300 leading-snug text-xs">
              Complete your payment within 15 minutes to receive a <strong className="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-blue-400 font-bold text-sm">{rewardValue}% accessory voucher</strong> for your next purchase.
            </p>
            <div className="bg-slate-950/50 p-2.5 rounded-lg flex justify-center border border-slate-800/80 shadow-inner">
              <CountdownDisplay remainingSeconds={remainingSeconds} />
            </div>
          </div>
        )}

        {displayState === "SUCCESS" && (
          <VoucherIssuedState rewardValue={rewardValue} />
        )}

        {displayState === "EXPIRED" && (
          <PromotionExpiredState />
        )}

        {displayState === "RETRY" && (
          <IssueRetryRequiredState onRetry={retryVoucherIssuance} />
        )}

        {displayState === "ERROR" && (
          <PromotionErrorState />
        )}
      </div>
    </div>
  );
};
