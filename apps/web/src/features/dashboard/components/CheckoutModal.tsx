import { useState } from "react";
import {
  X,
  ShoppingBag,
  CreditCard,
  Sparkles,
  Check,
  Gift,
} from "lucide-react";
import { Laptop } from "../types";

interface CheckoutModalProps {
  isOpen: boolean;
  onClose: () => void;
  product: Laptop | null;
  discountRate: number; // calculated from connected retail accounts
}

export default function CheckoutModal({
  isOpen,
  onClose,
  product,
  discountRate,
}: CheckoutModalProps) {
  const [promoCode, setPromoCode] = useState("");
  const [promoApplied, setPromoApplied] = useState(false);
  const [paymentSuccess, setPaymentSuccess] = useState(false);

  if (!isOpen) return null;
  if (!product) return null;

  const basePrice = product.price;
  const retailDiscount = basePrice * discountRate;
  const promoDiscount = promoApplied ? basePrice * 0.05 : 0; // extra 5% for promo "SHOPWISE5"
  const shipping = basePrice > 1500 ? 0 : 25;
  const tax = (basePrice - retailDiscount - promoDiscount) * 0.08;
  const finalTotal =
    basePrice - retailDiscount - promoDiscount + shipping + tax;

  const handleApplyPromo = () => {
    if (promoCode.toUpperCase() === "SHOPWISE5") {
      setPromoApplied(true);
      setPromoCode("");
    } else {
      alert("Invalid promo code. Try using code: SHOPWISE5");
    }
  };

  const handlePayment = () => {
    setPaymentSuccess(true);
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 p-4 backdrop-blur-md select-none">
      <div className="bg-surface-high border-outline-variant/30 flex max-h-[85vh] w-full max-w-md flex-col overflow-hidden rounded-2xl border shadow-2xl">
        {/* Header */}
        <div className="border-outline-variant/15 bg-surface-lowest flex items-center justify-between border-b p-5">
          <div className="flex items-center gap-2">
            <ShoppingBag size={18} className="text-[#4F7CFF]" />
            <div>
              <h3 className="text-on-surface font-display text-lg font-bold">
                Order Checkout
              </h3>
              <p className="text-on-surface-variant mt-1 text-xs font-medium">
                Complete the transaction for your optimal hardware product
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
          {paymentSuccess ? (
            <div className="flex flex-col items-center gap-4 py-12 text-center">
              <div className="flex h-16 w-16 items-center justify-center rounded-full border border-green-500/30 bg-green-500/10 text-green-400">
                <Check size={32} />
              </div>
              <div>
                <h4 className="font-display text-on-surface text-lg font-black">
                  Order Successful!
                </h4>
                <p className="text-on-surface-variant mx-auto mt-1.5 max-w-xs font-sans text-xs leading-relaxed">
                  Hardware configuration reconciliation transaction for{" "}
                  <strong>{product.name}</strong> is complete. The retailer will
                  contact you regarding inventory check within 15 minutes.
                </p>
              </div>
              <button
                onClick={() => {
                  setPaymentSuccess(false);
                  onClose();
                }}
                className="mt-4 cursor-pointer rounded-lg bg-[#4F7CFF] px-6 py-2 font-mono text-xs font-bold text-white transition-colors hover:bg-[#4F7CFF]/90"
              >
                Return to Workspace
              </button>
            </div>
          ) : (
            <>
              {/* Product overview card */}
              <div className="bg-surface-low border-outline-variant/15 flex items-center gap-3 rounded-xl border p-3">
                <img
                  src={product.image}
                  alt={product.name}
                  className="border-outline-variant/10 h-12 w-16 rounded-lg border object-cover"
                />
                <div className="min-w-0 flex-1">
                  <h4 className="font-display text-on-surface truncate text-xs font-bold">
                    {product.name}
                  </h4>
                  <p className="mt-0.5 font-mono text-[10px] text-[#4F7CFF]">
                    {product.specs.gpu}
                  </p>
                </div>
                <span className="text-on-surface shrink-0 font-mono text-xs font-bold">
                  ${product.price.toLocaleString("en-US")}
                </span>
              </div>

              {/* Promo code input */}
              <div className="flex gap-2">
                <input
                  value={promoCode}
                  onChange={(e) => setPromoCode(e.target.value)}
                  placeholder="Enter promo code (SHOPWISE5)..."
                  className="bg-surface-low border-outline-variant/30 text-on-surface placeholder:text-on-surface-variant/40 flex-1 rounded-lg border px-2.5 py-1.5 font-mono text-xs outline-none focus:border-[#4F7CFF]/60 focus:ring-0"
                />
                <button
                  onClick={handleApplyPromo}
                  className="bg-surface-low border-outline-variant/30 text-on-surface shrink-0 cursor-pointer rounded-lg border px-4 font-mono text-[11px] font-bold transition-colors hover:text-[#4F7CFF]"
                >
                  Apply
                </button>
              </div>

              {/* Pricing details */}
              <div className="bg-surface-low border-outline-variant/15 flex flex-col gap-2.5 rounded-xl border p-4">
                <div className="text-on-surface-variant flex justify-between text-xs">
                  <span>Base product price</span>
                  <span className="text-on-surface font-mono">
                    ${basePrice.toLocaleString("en-US")}
                  </span>
                </div>

                {retailDiscount > 0 && (
                  <div className="flex justify-between text-xs text-green-400">
                    <span>
                      Account integration discount ({discountRate * 100}%)
                    </span>
                    <span className="font-mono">
                      -${retailDiscount.toLocaleString("en-US")}
                    </span>
                  </div>
                )}

                {promoApplied && (
                  <div className="flex justify-between text-xs text-green-400">
                    <span>Exclusive promo code (SHOPWISE5)</span>
                    <span className="font-mono">
                      -${promoDiscount.toLocaleString("en-US")}
                    </span>
                  </div>
                )}

                <div className="text-on-surface-variant flex justify-between text-xs">
                  <span>Special shockproof shipping fee</span>
                  <span className="text-on-surface font-mono">
                    {shipping === 0 ? "FREE" : `$${shipping}`}
                  </span>
                </div>

                <div className="text-on-surface-variant flex justify-between text-xs">
                  <span>Customs clearance tax (8%)</span>
                  <span className="text-on-surface font-mono">
                    ${tax.toLocaleString("en-US", { maximumFractionDigits: 1 })}
                  </span>
                </div>

                <div className="border-outline-variant/15 text-on-surface flex justify-between border-t pt-2.5 text-sm font-bold">
                  <span className="flex items-center gap-1.5">
                    <Sparkles size={14} className="text-[#4F7CFF]" /> Total
                    payment
                  </span>
                  <span className="font-mono text-base text-[#4F7CFF]">
                    $
                    {finalTotal.toLocaleString("en-US", {
                      maximumFractionDigits: 1,
                    })}
                  </span>
                </div>
              </div>

              {/* Secure Payment confirmation */}
              <button
                onClick={handlePayment}
                className="mt-2 flex w-full cursor-pointer items-center justify-center gap-2 rounded-xl bg-[#4F7CFF] py-3 font-mono text-xs font-bold text-white shadow-lg transition-colors hover:bg-[#4F7CFF]/90"
              >
                <CreditCard size={15} /> CONFIRM SECURE PAYMENT
              </button>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
