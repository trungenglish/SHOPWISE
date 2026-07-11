import React from "react";
import {
  ZoomIn,
  ZoomOut,
  CheckCircle2,
  Star,
  Sparkles,
  AlertTriangle,
  GitCompare,
  HelpCircle,
  Bell,
} from "lucide-react";
import { Laptop, TrustFactors, PriceAlert } from "../types";

interface SpatialWorkspaceProps {
  products: Laptop[];
  activeProductId: string;
  setActiveProductId: (id: string) => void;
  savedIds: string[];
  onToggleSave: (id: string, e: React.MouseEvent) => void;
  trustScore: number;
  trustFactors: TrustFactors;
  graphNodes: string[];
  onChipClick: (suggestion: string) => void;
  onCompareAll: () => void;
  onExplainReasoning: () => void;
  onAccessories: () => void;
  onCheckout: () => void;
  discountRate: number;
  priceAlerts: PriceAlert[];
  onOpenPriceAlert: (laptop: Laptop) => void;
}

export default function SpatialWorkspace({
  products,
  activeProductId,
  setActiveProductId,
  savedIds,
  onToggleSave,
  trustScore,
  trustFactors,
  graphNodes = ["GPU Priority", "Thermals"],
  onChipClick,
  onCompareAll,
  onExplainReasoning,
  onAccessories,
  onCheckout,
  discountRate,
  priceAlerts,
  onOpenPriceAlert,
}: SpatialWorkspaceProps) {
  const suggestions = [
    `Compare with ${products[1]?.name || "Razer Blade"}`,
    `Cooling: Vapor Chamber vs Dual Fans`,
    `Test AI Performance with TensorRT`,
  ];

  const activeProduct =
    products.find((p) => p.id === activeProductId) || products[0];

  return (
    <main className="relative flex h-full flex-1 flex-col overflow-hidden bg-transparent p-6 select-none">
      {/* Title block */}
      <div className="relative z-10 mb-4 flex items-start justify-between">
        <div>
          <h2 className="font-display text-on-surface flex items-center gap-2 text-2xl font-extrabold tracking-tight">
            Spatial Workspace
          </h2>
          <p className="text-on-surface-variant mt-1 text-xs font-medium">
            Visual hardware mapping space & compatible configuration comparison
          </p>
        </div>

        {/* Zoom controls */}
        <div className="bg-surface-lowest/60 border-outline-variant/15 flex gap-1.5 rounded-lg border p-1 backdrop-blur">
          <button className="bg-surface-low hover:bg-surface-highest text-on-surface-variant hover:text-primary cursor-pointer rounded p-1.5 transition-colors">
            <ZoomIn size={15} />
          </button>
          <button className="bg-surface-low hover:bg-surface-highest text-on-surface-variant hover:text-primary cursor-pointer rounded p-1.5 transition-colors">
            <ZoomOut size={15} />
          </button>
        </div>
      </div>

      {/* Main Canvas Area */}
      <div className="relative h-full min-h-[480px] w-full flex-1">
        {/* Floating AI Constraint Suggestion Chips */}
        <div className="absolute top-[280px] left-[20px] z-20 flex max-w-[280px] flex-col gap-2.5">
          {suggestions.map((s, idx) => (
            <button
              key={idx}
              onClick={() => onChipClick(s)}
              className="floating-chip bg-primary/10 border-primary/20 hover:border-primary/50 flex cursor-pointer items-center gap-2 rounded-full border px-3.5 py-1.5 text-left font-mono text-[10px] font-bold tracking-wide text-[#4F7CFF] backdrop-blur-md transition-all hover:scale-[1.03] active:scale-[0.97]"
              style={{ animationDelay: `${idx * 1.5}s` }}
            >
              <Sparkles size={11} />
              <span>{s}</span>
            </button>
          ))}
        </div>

        {/* Trust Center Overlay Panel */}
        <div className="glass-card absolute top-0 right-0 z-20 flex w-[280px] flex-col gap-2 rounded-xl border border-[#4F7CFF]/20 p-4 text-xs shadow-[0_0_20px_rgba(79,124,255,0.08)]">
          <div className="border-outline-variant/15 flex items-center justify-between border-b pb-1.5">
            <span className="font-display text-on-surface-variant text-[10px] font-black tracking-wider">
              TRUST CENTER
            </span>
            <div className="flex items-center gap-1.5">
              <span className="text-on-surface font-mono text-xs font-extrabold">
                {trustScore || 94}%
              </span>
              <CheckCircle2 size={13} className="text-[#4F7CFF]" />
            </div>
          </div>
          <div className="flex items-center justify-between text-[11px]">
            <span className="text-on-surface-variant font-mono text-[10px]">
              Benchmark Cross-References
            </span>
            <span className="text-on-surface font-mono font-semibold">
              {trustFactors?.benchmarkSources || 8}
            </span>
          </div>
          <div className="flex items-center justify-between text-[11px]">
            <span className="text-on-surface-variant font-mono text-[10px]">
              Review Coverage
            </span>
            <span className="text-on-surface font-mono font-semibold">
              {trustFactors?.reviewCoverage || 85}%
            </span>
          </div>
          <div className="border-outline-variant/10 flex items-center justify-between border-t pt-1.5 text-[11px]">
            <span className="text-on-surface-variant font-mono text-[10px]">
              Retail Consensus
            </span>
            <span className="font-mono font-bold text-green-400">
              {trustFactors?.retailConsensus || "Recommended"}
            </span>
          </div>
          <div className="border-outline-variant/10 flex flex-col gap-1 border-t pt-1.5">
            <div className="flex items-center justify-between text-[11px]">
              <span className="text-on-surface-variant font-mono text-[10px]">
                Cumulative Reliability
              </span>
              <span className="flex items-center gap-0.5 font-mono text-xs text-green-400">
                Increasing
                <Star size={11} className="fill-green-400" />
              </span>
            </div>
            {/* Confidence bars */}
            <div className="mt-1 flex h-6 w-full items-end gap-1">
              {(trustFactors?.confidenceEvolution || [20, 40, 50, 75, 94]).map(
                (val, idx) => (
                  <div
                    key={idx}
                    className="w-1/5 rounded-t-sm bg-[#4F7CFF]/20 transition-all duration-700"
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

        {/* SVG Flow Connections */}
        <svg className="pointer-events-none absolute inset-0 z-0 h-full w-full opacity-60">
          <defs>
            <filter
              id="glow-filter"
              x="-20%"
              y="-20%"
              width="140%"
              height="140%"
            >
              <feGaussianBlur stdDeviation="3" result="blur" />
              <feComposite in="SourceGraphic" in2="blur" operator="over" />
            </filter>
          </defs>

          {/* Lines from User Goal -> priorities */}
          <path
            d="M 60 160 C 100 160, 100 100, 140 100"
            fill="none"
            stroke="#4F7CFF"
            strokeWidth="2"
            className="flow-line-highlight"
            filter="url(#glow-filter)"
          />
          <path
            d="M 60 160 C 100 160, 100 220, 140 220"
            fill="none"
            stroke="#4F7CFF"
            strokeWidth="2"
            className="flow-line-highlight"
            filter="url(#glow-filter)"
          />

          {/* Lines from Priorities -> Laptop Match targets */}
          {/* We assume laptops coordinates are roughly top-right clustered */}
          <path
            d="M 230 100 C 310 100, 310 80, 450 80"
            fill="none"
            stroke="#4F7CFF"
            strokeWidth="1.5"
            className="flow-line-highlight"
          />
          <path
            d="M 230 220 C 310 220, 310 120, 450 120"
            fill="none"
            stroke="#4F7CFF"
            strokeWidth="1.5"
            className="flow-line-highlight"
          />

          {/* Distant targets */}
          <path
            d="M 230 100 C 310 100, 310 240, 450 250"
            fill="none"
            stroke="#434654"
            strokeWidth="1"
            className="flow-line"
          />
          <path
            d="M 230 220 C 310 220, 310 280, 450 290"
            fill="none"
            stroke="#434654"
            strokeWidth="1"
            className="flow-line"
          />

          <path
            d="M 230 100 C 310 100, 310 380, 450 420"
            fill="none"
            stroke="#434654"
            strokeWidth="1"
            className="flow-line"
          />

          {/* Terminal node circles */}
          <circle cx="450" cy="80" r="4" fill="#4F7CFF" />
          <circle cx="450" cy="250" r="3" fill="#4F7CFF" opacity="0.6" />
          <circle cx="450" cy="420" r="3" fill="#4F7CFF" opacity="0.6" />

          {/* Main User Goal Node */}
          <rect
            x="10"
            y="140"
            width="100"
            height="40"
            rx="20"
            fill="#18181B"
            stroke="#4F7CFF"
            strokeWidth="1.5"
            className="pulse-border"
          />

          {/* Priorities Rects */}
          <rect
            x="135"
            y="85"
            width="100"
            height="30"
            rx="4"
            fill="#18181B"
            stroke="#4F7CFF"
            strokeWidth="1.5"
          />
          <rect
            x="135"
            y="205"
            width="100"
            height="30"
            rx="4"
            fill="#18181B"
            stroke="#4F7CFF"
            strokeWidth="1.5"
          />
        </svg>

        {/* HTML Labels Over SVG coordinates */}
        <div className="font-display text-on-surface pointer-events-none absolute top-[150px] left-[25px] z-10 text-[11px] font-bold">
          Target
        </div>
        <div className="font-display text-on-surface pointer-events-none absolute top-[92px] left-[150px] z-10 w-[75px] truncate text-center text-[10px] font-semibold">
          {graphNodes[0] || "GPU Priority"}
        </div>
        <div className="font-display text-on-surface pointer-events-none absolute top-[212px] left-[150px] z-10 w-[75px] truncate text-center text-[10px] font-semibold">
          {graphNodes[1] || "Thermals"}
        </div>

        {/* Match zone label */}
        <div className="absolute top-[-5px] left-[420px] font-mono text-[9px] font-bold tracking-widest text-[#4F7CFF] uppercase opacity-80 select-none">
          /// PRIMARY MATCH CLUSTER
        </div>

        {/* Laptop Stack (Spatial Cards) */}
        <div className="absolute top-[20px] left-[420px] z-10 flex w-[310px] flex-col gap-4">
          {products.map((laptop, index) => {
            const isActive = laptop.id === activeProductId;
            const isSaved = savedIds.includes(laptop.id);

            return (
              <div
                key={laptop.id}
                onClick={() => setActiveProductId(laptop.id)}
                className={`glass-card cursor-pointer overflow-hidden rounded-xl transition-all duration-300 ${
                  isActive
                    ? "border-primary ring-primary/40 scale-[1.02] border-2 shadow-[0_0_25px_rgba(79,124,255,0.25)] ring-1"
                    : "border-outline-variant/15 border opacity-75 hover:scale-[1.01] hover:opacity-100"
                }`}
              >
                {/* Header Match Score */}
                <div className="border-outline-variant/30 absolute top-3 right-3 z-10 flex items-center gap-1 rounded-full border bg-black/60 px-2 py-0.5 backdrop-blur">
                  <Star
                    size={11}
                    className={
                      isActive
                        ? "fill-[#4F7CFF] text-[#4F7CFF]"
                        : "text-on-surface-variant"
                    }
                  />
                  <span
                    className={`font-mono text-[9px] font-bold ${isActive ? "text-[#4F7CFF]" : "text-on-surface"}`}
                  >
                    {laptop.matchScore}% MATCH
                  </span>
                </div>

                {/* Laptop Image */}
                <div className="relative h-28 w-full">
                  <img
                    src={laptop.image}
                    alt={laptop.name}
                    referrerPolicy="no-referrer"
                    className="absolute inset-0 h-full w-full object-cover"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-zinc-950 via-zinc-950/20 to-transparent"></div>
                </div>

                {/* Body Content */}
                <div className="p-3">
                  <div className="mb-1 flex items-start justify-between">
                    <h3 className="font-display text-on-surface flex-1 truncate pr-2 text-sm font-bold">
                      {laptop.name}
                    </h3>
                    <div className="flex shrink-0 items-center gap-1.5">
                      <div className="flex flex-col items-end">
                        {discountRate > 0 ? (
                          <>
                            <span className="text-on-surface-variant font-mono text-[9px] leading-none line-through">
                              ${laptop.price.toLocaleString("en-US")}
                            </span>
                            <span className="font-mono text-xs font-bold text-emerald-400">
                              $
                              {Math.round(
                                laptop.price * (1 - discountRate)
                              ).toLocaleString("en-US")}
                            </span>
                          </>
                        ) : (
                          <span className="font-mono text-xs font-bold text-emerald-400">
                            ${laptop.price.toLocaleString("en-US")}
                          </span>
                        )}
                      </div>

                      {/* Bell price alert trigger */}
                      <button
                        onClick={(e) => {
                          e.stopPropagation();
                          onOpenPriceAlert(laptop);
                        }}
                        className={`cursor-pointer rounded-md p-1.5 transition-all ${
                          priceAlerts.some(
                            (a) => a.productId === laptop.id && a.active
                          )
                            ? "animate-pulse border border-[#4F7CFF]/40 bg-[#4F7CFF]/20 text-[#4F7CFF]"
                            : "text-on-surface-variant hover:bg-surface-low hover:text-[#4F7CFF]"
                        }`}
                        title="Set Price Alert"
                      >
                        <Bell size={12} />
                      </button>
                    </div>
                  </div>

                  {/* Tiny specs list */}
                  <div className="text-on-surface-variant mb-2.5 flex flex-col gap-0.5 font-mono text-[10px]">
                    <div className="border-outline-variant/10 flex justify-between border-b pb-0.5">
                      <span>GPU</span>
                      <span className="text-on-surface max-w-[150px] truncate">
                        {laptop.specs.gpu}
                      </span>
                    </div>
                    <div className="border-outline-variant/10 flex justify-between border-b pb-0.5">
                      <span>RAM</span>
                      <span className="text-on-surface">
                        {laptop.specs.ram}
                      </span>
                    </div>
                    <div className="flex justify-between">
                      <span>Cooling</span>
                      <span className="text-on-surface max-w-[150px] truncate">
                        {laptop.specs.cooling}
                      </span>
                    </div>
                  </div>

                  {/* Price alert active badge */}
                  {priceAlerts.some(
                    (a) => a.productId === laptop.id && a.active
                  ) && (
                    <div className="mt-1.5 mb-2 flex items-center justify-between rounded-lg border border-[#4F7CFF]/20 bg-[#4F7CFF]/10 px-2 py-1 font-mono text-[10px] text-[#4F7CFF]">
                      <span className="flex items-center gap-1">
                        <Bell size={10} className="animate-bounce" />
                        Price alert: &le; $
                        {priceAlerts
                          .find((a) => a.productId === laptop.id && a.active)
                          ?.targetPrice.toLocaleString()}
                      </span>
                      <span className="text-[9px] font-bold text-emerald-400">
                        (Current: $
                        {Math.round(
                          laptop.price * (1 - discountRate)
                        ).toLocaleString()}
                        )
                      </span>
                    </div>
                  )}

                  {/* Match Parameter Bars (only expand for Active) */}
                  {isActive && (
                    <div className="border-outline-variant/15 flex flex-col gap-1.5 border-t pt-2 transition-all">
                      <div className="flex items-center justify-between text-[9px]">
                        <span className="text-on-surface-variant w-14">
                          AI Performance
                        </span>
                        <div className="bg-surface-low ml-2 h-1.5 flex-1 overflow-hidden rounded-full">
                          <div
                            className="h-full rounded-full bg-[#4F7CFF]"
                            style={{ width: `${laptop.aiPerf}%` }}
                          />
                        </div>
                        <span className="ml-1 w-6 text-right font-mono">
                          {laptop.aiPerf}%
                        </span>
                      </div>
                      <div className="flex items-center justify-between text-[9px]">
                        <span className="text-on-surface-variant w-14">
                          3D Rendering
                        </span>
                        <div className="bg-surface-low ml-2 h-1.5 flex-1 overflow-hidden rounded-full">
                          <div
                            className="h-full rounded-full bg-[#4F7CFF]"
                            style={{ width: `${laptop.rendering}%` }}
                          />
                        </div>
                        <span className="ml-1 w-6 text-right font-mono">
                          {laptop.rendering}%
                        </span>
                      </div>
                      <div className="flex items-center justify-between text-[9px]">
                        <span className="text-on-surface-variant w-14">
                          Cooling
                        </span>
                        <div className="bg-surface-low ml-2 h-1.5 flex-1 overflow-hidden rounded-full">
                          <div
                            className="h-full rounded-full bg-[#4F7CFF]"
                            style={{ width: `${laptop.thermals}%` }}
                          />
                        </div>
                        <span className="ml-1 w-6 text-right font-mono">
                          {laptop.thermals}%
                        </span>
                      </div>

                      {/* Quick Pin/Save Action */}
                      <button
                        onClick={(e) => onToggleSave(laptop.id, e)}
                        className="text-on-primary-container border-primary/20 mt-2 flex w-full cursor-pointer items-center justify-center gap-1 rounded-md border bg-[#4F7CFF]/10 px-3 py-1.5 text-[11px] font-semibold transition-colors hover:bg-[#4F7CFF]/20"
                      >
                        {isSaved ? "Unpin from Board" : "Pin to Board"}
                      </button>
                    </div>
                  )}
                </div>
              </div>
            );
          })}
        </div>
      </div>

      {/* Persistent Floating Action Dock */}
      <div className="bg-surface-highest/80 border-outline-variant/40 absolute bottom-6 left-1/2 z-20 flex w-full max-w-2xl -translate-x-1/2 items-center justify-between gap-5 rounded-full border px-6 py-3 shadow-[0_10px_30px_rgba(0,0,0,0.5)] backdrop-blur-xl">
        <button
          onClick={onCompareAll}
          className="text-on-surface flex cursor-pointer items-center gap-2 font-mono text-[12px] font-bold transition-colors hover:text-[#4F7CFF]"
        >
          <GitCompare size={15} /> Compare All
        </button>
        <button
          onClick={onExplainReasoning}
          className="text-on-surface flex cursor-pointer items-center gap-2 font-mono text-[12px] font-bold transition-colors hover:text-[#4F7CFF]"
        >
          <Sparkles size={15} className="text-[#4F7CFF]" /> Explain Reasoning
        </button>
        <button
          onClick={onAccessories}
          className="text-on-surface flex cursor-pointer items-center gap-2 font-mono text-[12px] font-bold transition-colors hover:text-[#4F7CFF]"
        >
          <HelpCircle size={15} /> Compatible Accessories
        </button>
        <button
          onClick={onCheckout}
          className="cursor-pointer rounded-full bg-[#4F7CFF] px-5 py-2 font-mono text-[12px] font-black text-white shadow-[0_0_15px_rgba(79,124,255,0.4)] transition-all hover:bg-[#4F7CFF]/90"
        >
          Checkout
        </button>
      </div>
    </main>
  );
}
