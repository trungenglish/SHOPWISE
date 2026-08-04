import React from "react";
import { Tag } from "lucide-react";

interface PromotionBadgeProps {
  remainingSeconds: number;
  rewardValue?: number;
  onClick: () => void;
}

export const PromotionBadge: React.FC<PromotionBadgeProps> = ({ remainingSeconds, rewardValue = 15, onClick }) => {
  const minutes = Math.floor(remainingSeconds / 60);
  const seconds = remainingSeconds % 60;
  const timeString = `${minutes.toString().padStart(2, "0")}:${seconds.toString().padStart(2, "0")}`;

  return (
    <button
      onClick={onClick}
      className="fixed top-4 left-1/2 -translate-x-1/2 z-50 flex items-center gap-2 rounded-full bg-slate-900/90 backdrop-blur-md px-4 py-1.5 text-white shadow-xl shadow-cyan-500/10 border border-cyan-500/30 transition-all duration-300 ease-out hover:scale-105 hover:bg-slate-800 hover:border-cyan-400 focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:ring-offset-2 focus:ring-offset-slate-900 group"
      aria-label="Expand promotion panel"
    >
      <Tag className="w-3.5 h-3.5 text-cyan-400 group-hover:rotate-12 transition-transform duration-300" />
      <span className="font-semibold text-[11px] tracking-wide">
        {rewardValue}% Voucher <span className="opacity-50 mx-1">·</span> <span className="font-mono text-cyan-300 tabular-nums font-bold">{timeString}</span>
      </span>
    </button>
  );
};
