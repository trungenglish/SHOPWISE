import React from "react";
import { Handle, Position } from "@xyflow/react";
import { Target, Cpu, Star, Bell, Sparkles } from "lucide-react";
import { Laptop, PriceAlert } from "../types";

// User Goal/Intent Node
export function UserGoalNode({ data }: any) {
  return (
    <div className="min-w-[200px] rounded-2xl border-2 border-[#4F7CFF] bg-[#18181B]/95 px-5 py-3 text-left shadow-[0_0_20px_rgba(79,124,255,0.15)] select-none">
      <div className="mb-1 flex items-center gap-2">
        <Target size={14} className="text-[#4F7CFF]" />
        <span className="font-display text-[9px] font-black tracking-widest text-[#4F7CFF] uppercase">
          CORE INTENT
        </span>
      </div>
      <p className="font-sans text-xs font-bold text-white">
        {data.label || "System requirements"}
      </p>

      {/* Source handle connecting to priorities */}
      <Handle
        type="source"
        position={Position.Right}
        id="a"
        style={{ background: "#4F7CFF", width: 8, height: 8, right: -5 }}
      />
    </div>
  );
}

// Priorities Node (dynamic requirements matched by AI)
export function PriorityNode({ data }: any) {
  return (
    <div className="border-outline-variant/30 min-w-[160px] rounded-xl border bg-[#18181B]/90 px-4 py-2.5 text-left shadow-lg transition-colors select-none hover:border-[#4F7CFF]/50">
      <div className="mb-1 flex items-center gap-2">
        <Cpu size={12} className="text-[#4F7CFF]" />
        <span className="font-display text-on-surface-variant text-[9px] font-bold tracking-wider uppercase">
          PRIORITY CRITERIA
        </span>
      </div>
      <p className="font-sans text-xs font-bold text-white">{data.label}</p>

      {/* Target handle from Goal */}
      <Handle
        type="target"
        position={Position.Left}
        id="in"
        style={{ background: "#4F7CFF", width: 6, height: 6, left: -4 }}
      />
      {/* Source handle to Laptops */}
      <Handle
        type="source"
        position={Position.Right}
        id="out"
        style={{ background: "#4F7CFF", width: 6, height: 6, right: -4 }}
      />
    </div>
  );
}

// Laptop Node: A beautiful interactive laptop node that allows selecting products and setting price alerts
export function LaptopNode({ data }: any) {
  const laptop = data.laptop as Laptop;
  const isActive = data.isActive;
  const discountRate = data.discountRate || 0;
  const priceAlerts = data.priceAlerts || [];
  const onOpenPriceAlert = data.onOpenPriceAlert;
  const onSelect = data.onSelect;
  const isHovered = data.isHovered;
  const isSaved = data.isSaved;
  const onToggleSave = data.onToggleSave;

  const currentDiscountedPrice = Math.round(laptop.price * (1 - discountRate));
  const hasActiveAlert = priceAlerts.some(
    (a: PriceAlert) => a.productId === laptop.id && a.active
  );
  const activeAlert = priceAlerts.find(
    (a: PriceAlert) => a.productId === laptop.id && a.active
  );

  return (
    <div
      onClick={data.onSelect}
      onMouseEnter={data.onHover}
      onMouseLeave={data.onHoverLeave}
      className={`glass-card relative w-[240px] cursor-pointer overflow-hidden rounded-xl text-left transition-[transform,opacity,border-color,box-shadow,background-color] duration-200 ease-[var(--ease-out)] select-none ${
        isActive
          ? "border-primary ring-primary/40 scale-[1.02] border-2 bg-[#18181B]/95 shadow-[0_0_20px_rgba(79,124,255,0.25)] ring-1"
          : "border-outline-variant/15 border bg-[#18181B]/80 opacity-80 hover:scale-[1.01] hover:opacity-100 max-md:hover:scale-100"
      }`}
    >
      {/* Target Handle from Priorities */}
      <Handle
        type="target"
        position={Position.Left}
        style={{ background: "#4F7CFF", width: 6, height: 6, left: -4 }}
      />

      {/* Header Match Score */}
      <div className="border-outline-variant/30 absolute top-2.5 right-2.5 z-10 flex items-center gap-1 rounded-full border bg-black/75 px-2 py-0.5 backdrop-blur">
        <Star
          size={9}
          className={
            isActive
              ? "fill-[#4F7CFF] text-[#4F7CFF]"
              : "text-on-surface-variant"
          }
        />
        <span
          className={`font-mono text-[8px] font-black tracking-wide ${isActive ? "text-[#4F7CFF]" : "text-on-surface"}`}
        >
          {laptop.matchScore}% MATCH
        </span>
      </div>

      {/* Laptop Image */}
      <div className="relative h-20 w-full">
        <img
          src={laptop.image}
          alt={laptop.name}
          referrerPolicy="no-referrer"
          className="absolute inset-0 h-full w-full object-cover"
        />
        <div className="absolute inset-0 bg-gradient-to-t from-[#18181B] via-[#18181B]/20 to-transparent"></div>
      </div>

      {/* Body Content */}
      <div className="p-2.5">
        <div className="mb-1.5 flex items-start justify-between">
          <h3 className="font-display text-on-surface flex-1 truncate pr-2 text-[11px] leading-tight font-bold">
            {laptop.name}
          </h3>
          <div className="flex shrink-0 flex-col items-end leading-none">
            {discountRate > 0 ? (
              <>
                <span className="text-on-surface-variant mb-0.5 font-mono text-[8px] line-through">
                  {laptop.price.toLocaleString("vi-VN")} ₫
                </span>
                <span className="font-mono text-xs font-bold text-emerald-400">
                  {currentDiscountedPrice.toLocaleString("vi-VN")} ₫
                </span>
              </>
            ) : (
              <span className="font-mono text-xs font-bold text-emerald-400">
                {laptop.price.toLocaleString("vi-VN")} ₫
              </span>
            )}
          </div>
        </div>

        {/* Tiny specs */}
        <div className="mb-1.5 flex flex-wrap gap-1">
          <span className="bg-surface-lowest border-outline-variant/10 text-on-surface-variant rounded border px-1 py-0.5 font-mono text-[8px]">
            {laptop.specs?.gpu || "GPU"}
          </span>
          <span className="bg-surface-lowest border-outline-variant/10 text-on-surface-variant rounded border px-1 py-0.5 font-mono text-[8px]">
            {laptop.specs?.ram || "RAM"}
          </span>
        </div>

        {/* Alert configuration & status */}
        <div className="border-outline-variant/10 mt-2 flex items-center justify-between border-t pt-2">
          {hasActiveAlert ? (
            <div className="flex animate-pulse items-center gap-1 font-mono text-[8px] font-bold text-[#4F7CFF]">
              <Bell size={9} />
              <span>
                Target: &le; {activeAlert.targetPrice.toLocaleString("vi-VN")} ₫
              </span>
            </div>
          ) : (
            <div className="text-on-surface-variant flex items-center gap-1 font-mono text-[8px] font-medium">
              <Bell size={9} />
              <span>No alert set</span>
            </div>
          )}

          <button
            onClick={(e) => {
              e.stopPropagation();
              onOpenPriceAlert(laptop);
            }}
            className={`cursor-pointer rounded p-1 transition-[color,background-color,border-color] duration-200 ease-[var(--ease-out)] ${
              hasActiveAlert
                ? "border border-[#4F7CFF]/30 bg-[#4F7CFF]/20 text-[#4F7CFF] hover:bg-[#4F7CFF]/30"
                : "text-on-surface-variant hover:bg-surface-lowest hover:text-[#4F7CFF]"
            }`}
            title="Set Price Alert"
          >
            <Bell size={10} />
          </button>
        </div>
      </div>
    </div>
  );
}
