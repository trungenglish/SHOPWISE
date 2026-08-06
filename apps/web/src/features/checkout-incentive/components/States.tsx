import React from "react";
import { CheckCircle2, XCircle, AlertCircle, RefreshCw } from "lucide-react";

export const VoucherIssuedState: React.FC<{ rewardValue?: number }> = ({ rewardValue = 15 }) => (
  <div className="flex flex-col items-center space-y-3 text-center p-4">
    <CheckCircle2 className="w-12 h-12 text-green-500" />
    <h3 className="font-bold text-lg text-green-900">Voucher earned!</h3>
    <p className="text-sm text-green-800">
      You received {rewardValue}% off eligible accessories for your next purchase.
    </p>
  </div>
);

export const PromotionExpiredState: React.FC = () => (
  <div className="flex flex-col items-center space-y-3 text-center p-4">
    <XCircle className="w-12 h-12 text-slate-400" />
    <h3 className="font-bold text-lg text-slate-700">Offer Expired</h3>
    <p className="text-sm text-slate-500">
      This promotional window has ended.
    </p>
  </div>
);

export const PromotionErrorState: React.FC = () => (
  <div className="flex flex-col items-center space-y-3 text-center p-4">
    <AlertCircle className="w-12 h-12 text-red-500" />
    <h3 className="font-bold text-lg text-red-900">Something went wrong</h3>
    <p className="text-sm text-red-800">
      We couldn't load your promotion status.
    </p>
  </div>
);

export const IssueRetryRequiredState: React.FC<{ onRetry: () => void }> = ({ onRetry }) => (
  <div className="flex flex-col items-center space-y-3 text-center p-4">
    <AlertCircle className="w-12 h-12 text-orange-500" />
    <h3 className="font-bold text-lg text-orange-900">Voucher Pending</h3>
    <p className="text-sm text-orange-800">
      We received your payment, but there was a delay issuing your voucher.
    </p>
    <button 
      onClick={onRetry}
      className="mt-2 flex items-center gap-2 rounded bg-orange-100 px-4 py-2 text-orange-800 hover:bg-orange-200 transition-colors"
    >
      <RefreshCw className="w-4 h-4" />
      <span>Retry</span>
    </button>
  </div>
);
