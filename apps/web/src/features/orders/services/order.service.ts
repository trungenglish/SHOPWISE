import { Order, OrderItem, OrderStatus } from "../types";

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
  product_id?: string;
  retailer_offer_id?: string;
  name?: string;
  source_url?: string;
  quantity: number;
  unit_price: number;
}

export class OrderService {
  private static readonly API_BASE = "/api/v1/orders";
  private static readonly CUSTOMER_ID = "11111111-1111-1111-1111-111111111111"; // Dev bypass ID

  static async getOrders(): Promise<Order[]> {
    const [ordersRes, productsRes] = await Promise.all([
      fetch(`${this.API_BASE}?customer_id=${this.CUSTOMER_ID}`, {
        headers: { "x-checkout-auth-bypass": "true" },
      }),
      fetch("/api/v1/products"),
    ]);

    if (!ordersRes.ok) {
      throw new Error(`Failed to fetch orders: ${ordersRes.statusText}`);
    }

    const data: OrderListResponse = await ordersRes.json();
    let products: any[] = [];
    if (productsRes.ok) {
      const prodData = await productsRes.json();
      products = prodData.items || [];
    }

    return data.items.map((order) => this.mapOrderResponse(order, products));
  }

  private static mapOrderResponse(
    order: OrderResponse,
    products: any[]
  ): Order {
    return {
      id: order.order_id,
      date: order.created_at,
      status: this.mapStatus(order.status),
      totalPrice: order.total_amount,
      items: order.items.map((item) => this.mapOrderItem(item, products)),
    };
  }

  private static mapStatus(status: string): OrderStatus {
    switch (status) {
      case "PROCESSING":
        return "Processing";
      case "PENDING_SUPPLIER_CONFIRMATION":
        return "Pending Supplier";
      case "PENDING":
      default:
        return "Confirmed"; // Mapped from pending for UI
    }
  }

  private static mapOrderItem(
    item: OrderItemResponse,
    products: any[]
  ): OrderItem {
    // Attempt to enrich from real product data
    const apiProduct = products.find(
      (p) => p.ID === item.product_id || p.id === item.product_id
    );
    let thumbnail =
      "https://images.unsplash.com/photo-1603302576837-37561b2e2302?w=600&q=80";
    if (
      apiProduct &&
      apiProduct.Metadata &&
      apiProduct.Metadata.images &&
      apiProduct.Metadata.images.length > 0
    ) {
      thumbnail = apiProduct.Metadata.images[0];
    }

    return {
      id: item.product_id ?? item.retailer_offer_id ?? crypto.randomUUID(),
      name:
        item.name ??
        apiProduct?.Name ??
        apiProduct?.name ??
        "Sản phẩm ShopWise",
      quantity: item.quantity,
      price: item.unit_price,
      thumbnail: thumbnail,
    };
  }
}
