import { ShoppingBag } from "lucide-react";
import { OrderCard } from "./OrderCard";
import { EmptyOrders } from "./EmptyOrders";
import { OrderDetailModal } from "./OrderDetailModal";
import { getMockOrders } from "../mock";
import { Order } from "../types";
import { useState } from "react";

interface OrderHistoryPageProps {
  onContinueShopping: () => void;
}

export function OrderHistoryPage({ onContinueShopping }: OrderHistoryPageProps) {
  const orders = getMockOrders();
  const [selectedOrder, setSelectedOrder] = useState<Order | null>(null);

  return (
    <>
      <div className="z-10 flex flex-1 flex-col overflow-y-auto p-8 select-none w-full">
        <div className="mb-8">
        <h2 className="text-on-surface flex items-center gap-2 text-2xl font-extrabold tracking-tight">
          <ShoppingBag size={24} className="text-[#4F7CFF]" />
          Order History
        </h2>
        <p className="text-on-surface-variant mt-1 text-xs font-medium">
          Review confirmed orders placed through ShopWise.
        </p>
      </div>

        <div className="flex-1 max-w-4xl flex flex-col gap-4">
          {orders.length === 0 ? (
            <EmptyOrders onContinueShopping={onContinueShopping} />
          ) : (
            orders.map((order) => (
              <OrderCard 
                key={order.id} 
                order={order} 
                onViewDetails={(o) => setSelectedOrder(o)}
              />
            ))
          )}
        </div>
      </div>

      <OrderDetailModal 
        isOpen={!!selectedOrder} 
        onClose={() => setSelectedOrder(null)} 
        order={selectedOrder} 
      />
    </>
  );
}
