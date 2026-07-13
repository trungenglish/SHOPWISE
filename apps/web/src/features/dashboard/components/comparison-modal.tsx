import React from "react";
import { X, Check, Star, ShieldCheck } from "lucide-react";
import { Laptop } from "../types";

interface ComparisonModalProps {
  isOpen: boolean;
  onClose: () => void;
  products: Laptop[];
}

export default function ComparisonModal({
  isOpen,
  onClose,
  products,
}: ComparisonModalProps) {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 p-4 backdrop-blur-md select-none">
      <div className="bg-surface-high border-outline-variant/30 flex max-h-[85vh] w-full max-w-4xl flex-col overflow-hidden rounded-2xl border shadow-2xl">
        {/* Header */}
        <div className="border-outline-variant/15 bg-surface-lowest flex items-center justify-between border-b p-5">
          <div>
            <h3 className="text-on-surface font-display text-lg font-bold">
              Hardware Comparison Matrix
            </h3>
            <p className="text-on-surface-variant mt-1 text-xs font-medium">
              Analysis of real-world performance aspects based on reliable
              benchmark sources
            </p>
          </div>
          <button
            onClick={onClose}
            className="text-on-surface-variant hover:text-on-surface hover:bg-surface-highest cursor-pointer rounded-lg p-1.5 transition-colors"
          >
            <X size={18} />
          </button>
        </div>

        {/* Content Table */}
        <div className="flex-1 overflow-auto p-6">
          <div className="grid min-w-[700px] grid-cols-4 gap-4">
            {/* Row Header */}
            <div className="text-on-surface-variant col-span-1 flex flex-col justify-end pb-3 font-mono text-xs font-bold tracking-wider uppercase">
              Comparison Metrics
            </div>

            {/* Laptop Headers */}
            {products.map((p) => (
              <div
                key={p.id}
                className="bg-surface-low border-outline-variant/15 relative col-span-1 flex flex-col items-center rounded-xl border p-4 text-center"
              >
                <div className="absolute top-2 right-2 rounded bg-[#4F7CFF]/15 px-2 py-0.5 font-mono text-[9px] font-bold text-[#4F7CFF]">
                  {p.matchScore}% Match
                </div>
                <img
                  src={p.image}
                  alt={p.name}
                  className="border-outline-variant/10 mb-2 h-12 w-16 rounded-md border object-cover"
                />
                <h4 className="font-display text-on-surface w-full truncate text-xs font-bold">
                  {p.name}
                </h4>
                <span className="mt-1 font-mono text-sm font-bold text-[#4F7CFF]">
                  {p.price.toLocaleString("vi-VN")} ₫
                </span>
              </div>
            ))}

            {/* Price row */}
            <div className="border-outline-variant/10 text-on-surface-variant col-span-1 border-b py-3 font-sans text-xs font-bold">
              Suggested Retail Price
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="border-outline-variant/10 text-on-surface col-span-1 border-b py-3 text-center font-mono text-xs font-bold"
              >
                {p.price.toLocaleString("vi-VN")} ₫
              </div>
            ))}

            {/* GPU specs row */}
            <div className="border-outline-variant/10 text-on-surface-variant col-span-1 border-b py-3 font-sans text-xs font-bold">
              Graphics Card (GPU)
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="border-outline-variant/10 text-on-surface col-span-1 border-b py-3 text-center font-sans text-xs font-medium"
              >
                {p.specs.gpu}
              </div>
            ))}

            {/* RAM specs row */}
            <div className="border-outline-variant/10 text-on-surface-variant col-span-1 border-b py-3 font-sans text-xs font-bold">
              RAM / Memory
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="border-outline-variant/10 text-on-surface col-span-1 border-b py-3 text-center font-sans text-xs font-medium"
              >
                {p.specs.ram}
              </div>
            ))}

            {/* Cooling technology row */}
            <div className="border-outline-variant/10 text-on-surface-variant col-span-1 border-b py-3 font-sans text-xs font-bold">
              Cooling Technology
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="border-outline-variant/10 text-on-surface col-span-1 border-b py-3 text-center font-sans text-xs font-medium"
              >
                {p.specs.cooling}
              </div>
            ))}

            {/* Warranty row */}
            <div className="border-outline-variant/10 text-on-surface-variant col-span-1 border-b py-3 font-sans text-xs font-bold">
              Warranty
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="border-outline-variant/10 text-on-surface col-span-1 border-b py-3 text-center font-sans text-xs font-medium"
              >
                {p.specs.warranty || "Standard 1-Year"}
              </div>
            ))}

            {/* CPU & Screen Display */}
            <div className="border-outline-variant/10 text-on-surface-variant col-span-1 border-b py-3 font-sans text-xs font-bold">
              Processor (CPU)
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="border-outline-variant/10 text-on-surface col-span-1 border-b py-3 text-center font-sans text-xs font-medium"
              >
                {p.specs.cpu || "Optimizing..."}
              </div>
            ))}

            <div className="border-outline-variant/10 text-on-surface-variant col-span-1 border-b py-3 font-sans text-xs font-bold">
              Display Screen
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="border-outline-variant/10 text-on-surface col-span-1 border-b py-3 text-center font-sans text-xs font-medium"
              >
                {p.specs.screen || "Default Panel"}
              </div>
            ))}

            {/* AI Performance Score row */}
            <div className="border-outline-variant/10 text-on-surface-variant col-span-1 border-b py-3 font-sans text-xs font-bold">
              AI Processing Performance
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="border-outline-variant/10 col-span-1 border-b py-3 text-center font-mono text-xs font-bold text-[#4F7CFF]"
              >
                {p.aiPerf} / 100
              </div>
            ))}

            {/* Rendering capability row */}
            <div className="border-outline-variant/10 text-on-surface-variant col-span-1 border-b py-3 font-sans text-xs font-bold">
              3D Rendering / Graphics Performance
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="border-outline-variant/10 col-span-1 border-b py-3 text-center font-mono text-xs font-bold text-[#4F7CFF]"
              >
                {p.rendering} / 100
              </div>
            ))}

            {/* Thermal/cooling efficiency row */}
            <div className="border-outline-variant/10 text-on-surface-variant col-span-1 border-b py-3 font-sans text-xs font-bold">
              Thermal Efficiency
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="border-outline-variant/10 col-span-1 border-b py-3 text-center font-mono text-xs font-bold text-[#4F7CFF]"
              >
                {p.thermals} / 100
              </div>
            ))}

            {/* Key advantage */}
            <div className="text-on-surface-variant col-span-1 py-3 font-sans text-xs font-bold">
              Core Evaluation
            </div>
            {products.map((p) => (
              <div
                key={p.id}
                className="text-on-surface col-span-1 py-3 text-center font-sans text-[11px] leading-relaxed font-medium"
              >
                {p.matchExplanation}
              </div>
            ))}
          </div>
        </div>

        {/* Footer */}
        <div className="border-outline-variant/15 bg-surface-lowest flex items-center justify-end border-t p-5">
          <button
            onClick={onClose}
            className="cursor-pointer rounded-lg bg-[#4F7CFF] px-6 py-2.5 font-mono text-xs font-bold text-white shadow-lg transition-colors hover:bg-[#4F7CFF]/90"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  );
}
