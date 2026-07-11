import React, { useState, useEffect } from "react";
import {
  X,
  Bell,
  BellOff,
  DollarSign,
  Sparkles,
  Check,
  AlertCircle,
  Trash2,
} from "lucide-react";
import { Laptop, PriceAlert } from "../types";

interface PriceAlertModalProps {
  isOpen: boolean;
  onClose: () => void;
  product: Laptop | null;
  currentDiscountedPrice: number;
  activeAlert: PriceAlert | null;
  onSaveAlert: (productId: string, targetPrice: number) => void;
  onDeleteAlert: (productId: string) => void;
}

export default function PriceAlertModal({
  isOpen,
  onClose,
  product,
  currentDiscountedPrice,
  activeAlert,
  onSaveAlert,
  onDeleteAlert,
}: PriceAlertModalProps) {
  const [targetPrice, setTargetPrice] = useState<number>(0);

  useEffect(() => {
    if (product) {
      if (activeAlert) {
        setTargetPrice(activeAlert.targetPrice);
      } else {
        // Default to a 5% discount of current price as suggestion
        setTargetPrice(Math.round(currentDiscountedPrice * 0.95));
      }
    }
  }, [product, activeAlert, currentDiscountedPrice, isOpen]);

  if (!isOpen || !product) return null;

  const handleSave = () => {
    onSaveAlert(product.id, targetPrice);
    onClose();
  };

  const handleRemove = () => {
    onDeleteAlert(product.id);
    onClose();
  };

  // Percentage off
  const percentBelow = Math.round(
    ((product.price - targetPrice) / product.price) * 100
  );

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 p-4 backdrop-blur-md select-none">
      <div className="bg-surface-high border-outline-variant/30 flex max-h-[85vh] w-full max-w-md flex-col overflow-hidden rounded-2xl border shadow-2xl">
        {/* Header */}
        <div className="border-outline-variant/15 bg-surface-lowest flex items-center justify-between border-b p-5">
          <div className="flex items-center gap-2">
            <Bell size={18} className="animate-bounce text-[#4F7CFF]" />
            <div>
              <h3 className="text-on-surface font-display text-lg font-bold">
                Price Alert Setup
              </h3>
              <p className="text-on-surface-variant mt-1 text-xs font-medium">
                Track price changes and get notified automatically when price
                hits your target
              </p>
            </div>
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
          {/* Product card overview */}
          <div className="bg-surface-low border-outline-variant/15 flex items-center gap-3.5 rounded-xl border p-3.5">
            <img
              src={product.image}
              alt={product.name}
              className="border-outline-variant/10 h-12 w-16 rounded-lg border object-cover"
            />
            <div className="min-w-0 flex-1">
              <h4 className="font-display text-on-surface truncate text-xs font-bold">
                {product.name}
              </h4>
              <div className="mt-1 flex items-center gap-2">
                <span className="text-on-surface-variant font-mono text-xs line-through">
                  ${product.price.toLocaleString("en-US")}
                </span>
                <span className="font-mono text-xs font-bold text-emerald-400">
                  ${currentDiscountedPrice.toLocaleString("en-US")} (Current
                  price)
                </span>
              </div>
            </div>
          </div>

          {/* Current status info */}
          {activeAlert ? (
            <div className="flex items-center justify-between rounded-xl border border-[#4F7CFF]/20 bg-[#4F7CFF]/10 p-3">
              <div className="flex items-center gap-2">
                <Check size={14} className="text-[#4F7CFF]" />
                <span className="font-sans text-[11px] font-semibold text-[#4F7CFF]">
                  Active: Alert below{" "}
                  <strong>${activeAlert.targetPrice.toLocaleString()}</strong>
                </span>
              </div>
              <button
                onClick={handleRemove}
                className="flex cursor-pointer items-center gap-1 font-mono text-[10px] font-bold text-red-400 uppercase hover:text-red-300"
              >
                <Trash2 size={12} /> Cancel alert
              </button>
            </div>
          ) : (
            <div className="flex items-start gap-2 rounded-xl border border-amber-500/20 bg-amber-500/10 p-3">
              <AlertCircle
                size={14}
                className="mt-0.5 shrink-0 text-amber-400"
              />
              <p className="text-on-surface-variant font-sans text-[11px] leading-relaxed">
                When the price of this product drops below or equal to the
                target price (via retailer discount or direct price drop), the
                system will send an immediate notification.
              </p>
            </div>
          )}

          {/* Configuration interface */}
          <div className="border-outline-variant/15 bg-surface-low/50 rounded-xl border p-4">
            <label className="text-on-surface-variant mb-2 block font-mono text-xs font-bold uppercase">
              Target Price
            </label>

            <div className="relative flex items-center">
              <span className="text-on-surface-variant absolute left-3.5 font-mono text-sm">
                $
              </span>
              <input
                type="number"
                value={targetPrice}
                onChange={(e) => setTargetPrice(Number(e.target.value))}
                className="bg-surface-lowest border-outline-variant/30 text-on-surface w-full rounded-xl border py-3 pr-16 pl-8 font-mono text-sm outline-none focus:border-[#4F7CFF]/60 focus:ring-0"
                placeholder="Enter price..."
              />
              <span className="absolute right-3 rounded bg-red-400/10 px-2 py-0.5 font-mono text-[10px] font-bold text-red-400">
                -{percentBelow}%
              </span>
            </div>

            {/* Slider control */}
            <div className="mt-4">
              <input
                type="range"
                min={Math.round(product.price * 0.7)}
                max={product.price}
                value={targetPrice}
                onChange={(e) => setTargetPrice(Number(e.target.value))}
                className="bg-surface-lowest h-1.5 w-full cursor-pointer appearance-none rounded-lg accent-[#4F7CFF]"
              />
              <div className="text-on-surface-variant mt-1.5 flex justify-between font-mono text-[10px]">
                <span>30% Off (${Math.round(product.price * 0.7)})</span>
                <span>Original price (${product.price})</span>
              </div>
            </div>
          </div>

          {/* Quick tips */}
          <div className="bg-primary/5 border-outline-variant/10 flex gap-2 rounded-xl border p-3">
            <Sparkles
              size={14}
              className="mt-0.5 shrink-0 animate-pulse text-[#4F7CFF]"
            />
            <div className="text-on-surface-variant font-sans text-[10px] leading-relaxed">
              <strong>Price Optimization Tip:</strong> Connect your retailer
              account in the <strong>"Retail Account"</strong> tab on the left
              sidebar to automatically apply a 5% - 8% discount, instantly
              triggering the price alert!
            </div>
          </div>
        </div>

        {/* Footer */}
        <div className="border-outline-variant/15 bg-surface-lowest flex justify-between gap-3 border-t p-4">
          <button
            onClick={onClose}
            className="border-outline-variant/30 hover:bg-surface-low text-on-surface cursor-pointer rounded-lg border px-4 py-2.5 font-mono text-xs font-bold transition-colors"
          >
            Đóng
          </button>
          <button
            onClick={handleSave}
            disabled={targetPrice <= 0}
            className="flex cursor-pointer items-center gap-1.5 rounded-lg bg-[#4F7CFF] px-6 py-2.5 font-mono text-xs font-bold text-white shadow-lg transition-colors hover:bg-[#4F7CFF]/90 disabled:opacity-50"
          >
            <Check size={14} />{" "}
            {activeAlert ? "Cập Nhật Báo Giá" : "Thiết Lập Ngưỡng Giá"}
          </button>
        </div>
      </div>
    </div>
  );
}
