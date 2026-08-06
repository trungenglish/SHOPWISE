import { Order, OrderStatus } from "../types";

interface OrderCardProps {
  order: Order;
  onViewDetails: (order: Order) => void;
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

export function OrderCard({ order, onViewDetails }: OrderCardProps) {
  // We'll show the first item in the order for the thumbnail
  const firstItem = order.items[0];

  return (
    <div className="border-outline-variant/20 bg-surface-low flex w-full flex-col rounded-xl border p-5 sm:flex-row sm:items-center sm:justify-between">
      <div className="flex items-start gap-4 sm:items-center">
        {firstItem && (
          <div className="border-outline-variant/10 bg-surface-high h-16 w-16 shrink-0 overflow-hidden rounded-lg border">
            <img
              src={firstItem.thumbnail}
              alt={firstItem.name}
              className="h-full w-full object-cover"
            />
          </div>
        )}
        <div className="flex flex-col gap-1">
          <div className="flex items-center gap-2">
            <span className="text-on-surface font-mono text-sm font-semibold">
              {order.id}
            </span>
            <span className="text-on-surface-variant text-xs">
              • {formatDate(order.date)}
            </span>
          </div>
          {firstItem && (
            <p className="text-on-surface line-clamp-1 text-sm font-medium">
              {firstItem.name}
              {order.items.length > 1 && (
                <span className="text-on-surface-variant font-normal">
                  {" "}
                  and {order.items.length - 1} other item{order.items.length > 2 ? 's' : ''}
                </span>
              )}
            </p>
          )}
        </div>
      </div>
      
        <div className="mt-4 flex sm:mt-0 sm:shrink-0 flex-col items-end gap-2">
          <span
            className={`flex items-center rounded-full border px-3 py-1 text-xs font-semibold ${getStatusColor(
              order.status
            )}`}
          >
            {order.status}
          </span>
          <p className="text-on-surface-variant text-xs mt-1">
            Total: <span className="font-semibold text-on-surface">{formatPrice(order.totalPrice)}</span>
          </p>
          <button
            onClick={() => onViewDetails(order)}
            className="text-[#4F7CFF] hover:text-[#4F7CFF]/80 text-xs font-semibold transition-colors cursor-pointer"
          >
            View Details
          </button>
        </div>
    </div>
  );
}
