import { X, ShoppingBag } from "lucide-react";
import { Order, OrderStatus } from "../types";

interface OrderDetailModalProps {
  isOpen: boolean;
  onClose: () => void;
  order: Order | null;
}

const getStatusColor = (status: OrderStatus) => {
  switch (status) {
    case "Confirmed":
      return "bg-blue-500/10 text-blue-500 border-blue-500/20";
    case "Processing":
      return "bg-amber-500/10 text-amber-500 border-amber-500/20";
    case "Delivered":
      return "bg-green-500/10 text-green-500 border-green-500/20";
    default:
      return "bg-gray-500/10 text-gray-400 border-gray-500/20";
  }
};

const formatPrice = (price: number) => {
  return new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(price);
};

const formatDate = (dateString: string) => {
  return new Intl.DateTimeFormat("vi-VN", {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  }).format(new Date(dateString));
};

export function OrderDetailModal({ isOpen, onClose, order }: OrderDetailModalProps) {
  if (!isOpen || !order) return null;

  return (
    <div className="bg-black/80 fixed inset-0 z-50 flex items-center justify-center p-4 select-none backdrop-blur-sm">
      <div
        className="bg-surface-low border-outline-variant/30 flex w-full max-w-2xl flex-col overflow-hidden rounded-xl border shadow-2xl"
        style={{ maxHeight: "calc(100vh - 2rem)" }}
      >
        {/* Header */}
        <div className="border-outline-variant/20 bg-surface flex items-center justify-between border-b p-4 md:px-6">
          <div className="flex items-center gap-3">
            <ShoppingBag size={20} className="text-[#4F7CFF]" />
            <div>
              <h2 className="text-on-surface text-lg font-bold">
                Order Details
              </h2>
              <p className="text-on-surface-variant font-mono text-sm">
                {order.id}
              </p>
            </div>
          </div>
          <button
            onClick={onClose}
            className="text-on-surface-variant hover:text-on-surface hover:bg-surface-highest cursor-pointer rounded-lg p-2 transition-colors"
            aria-label="Close modal"
          >
            <X size={20} />
          </button>
        </div>

        {/* Content */}
        <div className="flex-1 overflow-y-auto p-4 md:p-6 flex flex-col gap-6">
          
          {/* Order Info */}
          <div className="flex justify-between items-start bg-surface-high border-outline-variant/20 rounded-lg border p-4">
            <div className="flex flex-col gap-1">
              <span className="text-on-surface-variant text-xs uppercase tracking-wider font-semibold">Date Placed</span>
              <span className="text-on-surface font-medium">{formatDate(order.date)}</span>
            </div>
            <div className="flex flex-col items-end gap-1">
              <span className="text-on-surface-variant text-xs uppercase tracking-wider font-semibold">Status</span>
              <span
                className={`flex items-center rounded-full border px-2 py-0.5 text-xs font-semibold ${getStatusColor(
                  order.status
                )}`}
              >
                {order.status}
              </span>
            </div>
          </div>

          {/* Items List */}
          <div>
            <h3 className="text-on-surface text-sm font-bold tracking-wide uppercase mb-3">Items</h3>
            <div className="flex flex-col gap-3">
              {order.items.map((item) => (
                <div key={item.id} className="flex items-center gap-4 border-b border-outline-variant/10 pb-3 last:border-0 last:pb-0">
                  <div className="border-outline-variant/10 bg-surface h-16 w-16 shrink-0 overflow-hidden rounded-lg border">
                    <img
                      src={item.thumbnail}
                      alt={item.name}
                      className="h-full w-full object-cover"
                    />
                  </div>
                  <div className="flex flex-1 flex-col">
                    <span className="text-on-surface text-sm font-medium">{item.name}</span>
                    <span className="text-on-surface-variant text-xs mt-0.5">Qty: {item.quantity}</span>
                  </div>
                  <div className="text-on-surface font-semibold text-sm">
                    {formatPrice(item.price * item.quantity)}
                  </div>
                </div>
              ))}
            </div>
          </div>

          {/* Summary */}
          <div>
            <h3 className="text-on-surface text-sm font-bold tracking-wide uppercase mb-3">Summary</h3>
            <div className="bg-surface-high border-outline-variant/20 rounded-lg border p-4 flex flex-col gap-2 text-sm">
              <div className="flex justify-between">
                <span className="text-on-surface-variant">Subtotal</span>
                <span className="text-on-surface font-medium">{formatPrice(order.totalPrice)}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-on-surface-variant">Shipping</span>
                <span className="text-on-surface font-medium">Free</span>
              </div>
              <div className="border-t border-outline-variant/20 my-1"></div>
              <div className="flex justify-between">
                <span className="text-on-surface font-bold">Total</span>
                <span className="text-[#4F7CFF] font-bold text-base">{formatPrice(order.totalPrice)}</span>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>
  );
}
