import { useQuery } from "@tanstack/react-query";
import { OrderService } from "../services/order.service";
import { Order } from "../types";

export function useOrders() {
  return useQuery<Order[], Error>({
    queryKey: ["orders"],
    queryFn: () => OrderService.getOrders(),
    staleTime: 1000 * 60 * 5, // 5 minutes
    retry: 2,
  });
}
