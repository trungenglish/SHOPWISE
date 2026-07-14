import React, { useState, useMemo } from "react";
import { X, HelpCircle, Check, ArrowRight } from "lucide-react";

interface Accessory {
  name: string;
  price: number;
  category: string;
  reason: string;
  image?: string;
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
  const [activeFilter, setActiveFilter] = useState("All");

  const categories = useMemo(() => {
    if (!accessories) return ["All"];
    const cats = new Set(accessories.map((a) => a.category));
    return ["All", ...Array.from(cats)];
  }, [accessories]);

  const filteredAccessories = useMemo(() => {
    if (!accessories) return [];
    if (activeFilter === "All") return accessories;
    return accessories.filter((a) => a.category === activeFilter);
  }, [accessories, activeFilter]);

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 p-4 backdrop-blur-md select-none">
      <div className="bg-surface-high border-outline-variant/30 flex max-h-[85vh] w-full max-w-xl flex-col overflow-hidden rounded-2xl border shadow-2xl">
        {/* Header */}
        <div className="border-outline-variant/15 bg-surface-lowest flex flex-col border-b">
          <div className="flex items-center justify-between p-5 pb-3">
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
          
          {/* Filters */}
          {categories.length > 1 && (
            <div className="flex gap-2 overflow-x-auto px-5 pb-4 scrollbar-hide">
              {categories.map((cat) => (
                <button
                  key={cat}
                  onClick={() => setActiveFilter(cat)}
                  className={`shrink-0 cursor-pointer rounded-full px-4 py-1.5 font-mono text-xs font-bold transition-[color,background-color,border-color] duration-200 ease-[var(--ease-out)] ${
                    activeFilter === cat
                      ? "bg-[#4F7CFF] text-white shadow-md shadow-[#4F7CFF]/20 border border-[#4F7CFF]"
                      : "bg-surface border-outline-variant/30 text-on-surface-variant hover:bg-surface-highest hover:text-on-surface border"
                  }`}
                >
                  {cat}
                </button>
              ))}
            </div>
          )}
        </div>

        {/* Content */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 flex-1 overflow-auto p-5 content-start">
          {filteredAccessories.length > 0 ? (
            filteredAccessories.map((acc, index) => (
              <div
                key={index}
                className="bg-surface-low border-outline-variant/15 group flex cursor-pointer flex-col overflow-hidden rounded-xl border transition-colors hover:border-[#4F7CFF]/40"
              >
                {/* Image */}
                {acc.image ? (
                  <div className="border-outline-variant/10 h-32 w-full overflow-hidden border-b">
                    <img
                      src={acc.image}
                      alt={acc.name}
                      className="h-full w-full object-cover transition-transform duration-500 group-hover:scale-105"
                    />
                  </div>
                ) : (
                  <div className="border-outline-variant/10 bg-[#4F7CFF]/5 text-[#4F7CFF] flex h-32 w-full items-center justify-center border-b">
                    <Check size={24} />
                  </div>
                )}
                {/* Details */}
                <div className="flex flex-1 flex-col p-4">
                  <div className="mb-2 flex items-center justify-between">
                    <span className="bg-[#4F7CFF]/10 text-[#4F7CFF] rounded px-2 py-0.5 font-mono text-[10px] font-bold uppercase">
                      {acc.category}
                    </span>
                    <span className="text-on-surface font-mono text-sm font-bold">
                      {acc.price.toLocaleString("vi-VN")} ₫
                    </span>
                  </div>
                  <h4 className="font-display text-on-surface mb-1 text-sm font-bold">
                    {acc.name}
                  </h4>
                  <p className="text-on-surface-variant line-clamp-2 font-sans text-xs leading-relaxed">
                    {acc.reason}
                  </p>
                </div>
              </div>
            ))
          ) : (
            <div className="text-on-surface-variant col-span-full py-8 text-center font-medium">
              No accessories found for this category.
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
