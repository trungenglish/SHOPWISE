import { ShoppingBag, Loader2, AlertCircle, RefreshCw } from "lucide-react";
import { OrderCard } from "./OrderCard";
import { EmptyOrders } from "./EmptyOrders";
import { OrderDetailModal } from "./OrderDetailModal";
import { useOrders } from "../hooks/use-orders";
import { Order } from "../types";
import { useState } from "react";
import { Button } from "@shopwise/ui/components/button";

interface OrderHistoryPageProps {
  onContinueShopping: () => void;
}

export function OrderHistoryPage({ onContinueShopping }: OrderHistoryPageProps) {
  const { data: orders = [], isLoading, isError, refetch } = useOrders();
  const [selectedOrder, setSelectedOrder] = useState<Order | null>(null);

  return (
    <>
      <div className="z-10 flex flex-1 flex-col overflow-y-auto p-8 select-none w-full">
        <div className="mb-8">
          <h2 className="text-on-surface flex items-center gap-2 text-2xl font-extrabold tracking-tight">
            <ShoppingBag size={24} className="text-[#4F7CFF]" />
            Đơn hàng đã đặt
          </h2>
          <p className="text-on-surface-variant mt-1 text-xs font-medium">
            Xem lại các đơn hàng đã được xác nhận thông qua ShopWise.
          </p>
        </div>

        <div className="flex-1 max-w-4xl flex flex-col gap-4">
          {isLoading ? (
            <div className="flex items-center justify-center p-12 text-on-surface-variant">
              <Loader2 className="animate-spin" size={32} />
            </div>
          ) : isError ? (
            <div className="flex flex-col items-center justify-center p-12 text-center rounded-xl border border-outline-variant/30 bg-surface">
              <AlertCircle size={48} className="text-red-500 mb-4" />
              <h3 className="text-on-surface font-bold text-lg mb-2">Đã xảy ra lỗi</h3>
              <p className="text-on-surface-variant text-sm mb-6 max-w-md">
                Không thể tải danh sách đơn hàng. Vui lòng kiểm tra kết nối mạng và thử lại.
              </p>
              <Button onClick={() => refetch()} variant="outline" className="gap-2">
                <RefreshCw size={16} /> Thử lại
              </Button>
            </div>
          ) : orders.length === 0 ? (
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
