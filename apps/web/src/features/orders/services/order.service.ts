import { Order, OrderItem, OrderStatus } from "../types";
import { initialProducts } from "../../dashboard/mock-data";

export interface OrderListResponse {
  items: OrderResponse[];
  limit: number;
  offset: number;
}

export interface OrderResponse {
  order_id: string;
  customer_id: string;
  customer_name: string;
  customer_email: string;
  customer_phone: string;
  fulfillment_method: string;
  shipping_address?: string;
  items: OrderItemResponse[];
  coupon_code?: string;
  subtotal_amount: number;
  discount_amount: number;
  shipping_amount: number;
  tax_amount: number;
  total_amount: number;
  status: string;
  created_at: string;
}

export interface OrderItemResponse {
  product_id: string;
  quantity: number;
  unit_price: number;
}

export class OrderService {
  private static readonly API_BASE = "http://localhost:8080/api/v1/orders";
  private static readonly CUSTOMER_ID = "11111111-1111-1111-1111-111111111111"; // Dev bypass ID

  static async getOrders(): Promise<Order[]> {
    const res = await fetch(`${this.API_BASE}?customer_id=${this.CUSTOMER_ID}`, {
      headers: {
        "x-checkout-auth-bypass": "true",
      },
    });

    if (!res.ok) {
      throw new Error(`Failed to fetch orders: ${res.statusText}`);
    }

    const data: OrderListResponse = await res.json();

    return data.items.map((order) => this.mapOrderResponse(order));
  }

  private static mapOrderResponse(order: OrderResponse): Order {
    return {
      id: order.order_id,
      date: order.created_at,
      status: this.mapStatus(order.status),
      totalPrice: order.total_amount,
      items: order.items.map((item) => this.mapOrderItem(item)),
    };
  }

  private static mapStatus(status: string): OrderStatus {
    switch (status) {
      case "PROCESSING":
        return "Processing";
      case "PENDING":
      default:
        return "Confirmed"; // Mapped from pending for UI
    }
  }

  private static mapOrderItem(item: OrderItemResponse): OrderItem {
    // Attempt to enrich from mock data since API only returns ID and price
    const mockProduct = initialProducts.find((p) => p.id === item.product_id);
    
    return {
      id: item.product_id,
      name: mockProduct?.name || "Sản phẩm ShopWise",
      quantity: item.quantity,
      price: item.unit_price,
      thumbnail: mockProduct?.image || "https://images.unsplash.com/photo-1603302576837-37561b2e2302?w=600&q=80",
    };
  }
}
