import React, { useState } from "react";
import { TrustFactors } from "../types";
import { ShieldCheck, CheckCircle2, Star, Loader2, X, ZoomOut, ZoomIn } from "lucide-react";
import { useReactFlow } from "@xyflow/react";

interface TrustCenterHeaderProps {
  readonly trustScore: number;
  readonly trustFactors: TrustFactors;
  readonly isSimulating: boolean;
}

const TrustCenterHeader = ({
  trustScore,
  trustFactors,
  isSimulating,
}: TrustCenterHeaderProps) => {
  const [isOpen, setIsOpen] = useState(false);
  const { zoomIn, zoomOut, fitView } = useReactFlow();
  return (
    <div className="pointer-events-none absolute top-4 right-6 z-50 flex flex-col items-end">
      <div className="flex flex-row items-center gap-2">
        <div className="bg-surface-lowest/60 border-outline-variant/15 flex gap-1.5 rounded-lg border p-1 backdrop-blur">
          <button
            type="button"
            onClick={() => zoomIn()}
            className="bg-surface-low hover:bg-surface-highest text-on-surface-variant hover:text-primary cursor-pointer rounded p-1.5 transition-colors"
            title="Zoom In"
          >
            <ZoomIn size={15} />
          </button>
          <button
            type="button"
            onClick={() => zoomOut()}
            className="bg-surface-low hover:bg-surface-highest text-on-surface-variant hover:text-primary cursor-pointer rounded p-1.5 transition-colors"
            title="Zoom Out"
          >
            <ZoomOut size={15} />
          </button>
          <button
            type="button"
            onClick={() => fitView()}
            className="bg-surface-low hover:bg-surface-highest text-on-surface-variant hover:text-primary cursor-pointer rounded px-2 py-1 font-mono text-[10px] font-bold transition-colors"
            title="Fit View"
          >
            FIT
          </button>
        </div>

        {/* Header Button */}
        <button
          type="button"
          onClick={() => setIsOpen(!isOpen)}
          className="bg-surface-high/80 hover:bg-surface-highest border-outline-variant/30 group pointer-events-auto flex cursor-pointer items-center gap-2 rounded-full border px-3 py-1.5 shadow-lg backdrop-blur-md transition-colors"
        >
          <div className="relative">
            {isSimulating ? (
              <Loader2 size={16} className="animate-spin text-[#4F7CFF]" />
            ) : (
              <ShieldCheck
                size={16}
                className="text-[#4F7CFF] transition-transform duration-200 ease-[var(--ease-out)] group-hover:scale-110 max-md:group-hover:scale-100"
              />
            )}
          </div>
          <span className="font-display text-on-surface text-[11px] font-bold tracking-wide">
            {isSimulating ? "VERIFYING" : "TRUST CENTER"}
          </span>
        </button>
      </div>

      {/* Expandable Panel */}
      {isOpen && (
        <div className="glass-card animate-fade-in-down pointer-events-auto mt-2 flex w-[280px] origin-top-right flex-col gap-2 rounded-xl border border-[#4F7CFF]/20 p-4 text-xs shadow-[0_15px_40px_-10px_rgba(79,124,255,0.2)]">
          <div className="border-outline-variant/15 flex items-center justify-between border-b pb-2">
            <span className="font-display text-on-surface-variant text-[10px] font-black tracking-wider">
              TRUST CENTER
            </span>
            <div className="flex items-center gap-1.5">
              <span className="text-on-surface font-mono text-xs font-extrabold">
                {trustScore || 0}%
              </span>
              <CheckCircle2 size={13} className="text-[#4F7CFF]" />
              <button
                type="button"
                onClick={() => setIsOpen(false)}
                className="text-on-surface-variant hover:text-on-surface ml-2 cursor-pointer transition-colors"
              >
                <X size={14} />
              </button>
            </div>
          </div>

          <div className="flex flex-col gap-2 pt-1">
            <div className="flex items-center justify-between text-[11px]">
              <span className="text-on-surface-variant font-mono text-[10px]">
                Benchmark Sources
              </span>
              <span className="text-on-surface font-mono font-semibold">
                {trustFactors?.benchmarkSources || 0}
              </span>
            </div>
            <div className="flex items-center justify-between text-[11px]">
              <span className="text-on-surface-variant font-mono text-[10px]">
                Review Coverage
              </span>
              <span className="text-on-surface font-mono font-semibold">
                {trustFactors?.reviewCoverage || 0}%
              </span>
            </div>
            <div className="border-outline-variant/10 mt-1 flex items-center justify-between border-t pt-2 text-[11px]">
              <span className="text-on-surface-variant font-mono text-[10px]">
                Retail Consensus
              </span>
              <span className="max-w-[140px] truncate text-right font-mono font-bold text-emerald-400">
                {trustFactors?.retailConsensus || "Awaiting Analysis"}
              </span>
            </div>
          </div>

          <div className="border-outline-variant/10 mt-1 flex flex-col gap-1 border-t pt-2">
            <div className="flex items-center justify-between text-[11px]">
              <span className="text-on-surface-variant font-mono text-[10px]">
                Cumulative Reliability
              </span>
              <span className="flex items-center gap-0.5 font-mono text-xs text-emerald-400">
                {trustFactors?.confidenceEvolution?.length > 0
                  ? "Increasing"
                  : "N/A"}
                <Star size={11} className="fill-emerald-400" />
              </span>
            </div>
            {/* Confidence bars */}
            <div className="mt-1 flex h-6 w-full items-end gap-1">
              {(trustFactors?.confidenceEvolution || [0, 0, 0, 0, 0]).map(
                (val, idx) => (
                  <div
                    key={idx}
                    className="w-1/5 rounded-t-sm bg-[#4F7CFF]/20 transition-[height,background-color] duration-300 ease-[var(--ease-out)]"
                    style={{
                      height: `${Math.max(10, val)}%`,
                      backgroundColor:
                        idx === 4
                          ? "#4F7CFF"
                          : `rgba(79, 124, 255, ${0.3 + idx * 0.15})`,
                      boxShadow: idx === 4 ? "0 0 8px #4F7CFF" : "none",
                    }}
                  />
                )
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default TrustCenterHeader;
