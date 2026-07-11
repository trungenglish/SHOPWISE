import React from "react";
import { X, HelpCircle, Check, ArrowRight } from "lucide-react";

interface Accessory {
  name: string;
  price: number;
  category: string;
  reason: string;
}

interface AccessoriesModalProps {
  isOpen: boolean;
  onClose: () => void;
  accessories: Accessory[];
}

export default function AccessoriesModal({
  isOpen,
  onClose,
  accessories,
}: AccessoriesModalProps) {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 p-4 backdrop-blur-md select-none">
      <div className="bg-surface-high border-outline-variant/30 flex max-h-[85vh] w-full max-w-xl flex-col overflow-hidden rounded-2xl border shadow-2xl">
        {/* Header */}
        <div className="border-outline-variant/15 bg-surface-lowest flex items-center justify-between border-b p-5">
          <div>
            <h3 className="text-on-surface font-display text-lg font-bold">
              Recommended Accessories
            </h3>
            <p className="text-on-surface-variant mt-1 text-xs font-medium">
              Accessories to maximize your hardware capabilities
            </p>
          </div>
          <button
            onClick={onClose}
            className="text-on-surface-variant hover:text-on-surface hover:bg-surface-highest cursor-pointer rounded-lg p-1.5 transition-colors"
          >
            <X size={18} />
          </button>
        </div>

        {/* Content */}
        <div className="flex flex-1 flex-col gap-4 overflow-auto p-5">
          {accessories && accessories.length > 0 ? (
            accessories.map((acc, index) => (
              <div
                key={index}
                className="bg-surface-low border-outline-variant/15 flex items-start gap-3 rounded-xl border p-4 transition-colors hover:border-[#4F7CFF]/40"
              >
                <div className="shrink-0 rounded-lg bg-[#4F7CFF]/10 p-2.5 text-[#4F7CFF]">
                  <Check size={16} />
                </div>
                <div className="min-w-0 flex-1">
                  <div className="mb-1 flex items-center justify-between">
                    <span className="rounded bg-[#4F7CFF]/10 px-2 py-0.5 font-mono text-[10px] font-bold text-[#4F7CFF] uppercase">
                      {acc.category}
                    </span>
                    <span className="text-on-surface font-mono text-sm font-bold">
                      ${acc.price}
                    </span>
                  </div>
                  <h4 className="font-display text-on-surface mb-1 truncate text-sm font-bold">
                    {acc.name}
                  </h4>
                  <p className="text-on-surface-variant font-sans text-xs leading-relaxed">
                    {acc.reason}
                  </p>
                </div>
              </div>
            ))
          ) : (
            <div className="text-on-surface-variant py-8 text-center font-medium">
              Searching for optimal accessories for this hardware...
            </div>
          )}
        </div>

        {/* Footer */}
        <div className="border-outline-variant/15 bg-surface-lowest flex justify-end border-t p-4">
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
