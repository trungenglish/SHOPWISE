import { ShoppingBag } from "lucide-react";

interface EmptyOrdersProps {
  onContinueShopping: () => void;
}

export function EmptyOrders({ onContinueShopping }: EmptyOrdersProps) {
  return (
    <div className="border-outline-variant/15 bg-surface-low/30 mx-auto flex w-full max-w-2xl flex-1 flex-col items-center justify-center rounded-2xl border border-dashed p-12 text-center h-full">
      <div className="bg-surface-high mb-4 flex h-16 w-16 items-center justify-center rounded-full">
        <ShoppingBag size={32} className="text-on-surface-variant/60" />
      </div>
      <h3 className="text-on-surface text-xl font-bold mb-2">
        No orders yet
      </h3>
      <p className="text-on-surface-variant max-w-sm text-sm leading-relaxed mb-8">
        After completing checkout, your orders will appear here so you can easily track them.
      </p>
      <button
        onClick={onContinueShopping}
        className="bg-[#4F7CFF] text-white hover:bg-[#4F7CFF]/90 flex cursor-pointer items-center justify-center rounded-lg px-6 py-3 text-sm font-semibold transition-colors"
      >
        Continue shopping
      </button>
    </div>
  );
}
