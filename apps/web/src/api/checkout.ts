export interface CheckoutItemRequest {
  product_id: string;
  quantity: number;
}

export interface CheckoutRequest {
  customer_id: string;
  items: CheckoutItemRequest[];
  fulfillment_method: "DELIVERY" | "STORE_PICKUP";
  shipping_address?: string;
  coupon_code?: string;
}

export interface OrderItemResponse {
  product_id: string;
  quantity: number;
  unit_price: number;
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

export async function checkout(request: CheckoutRequest): Promise<OrderResponse> {
  const token = localStorage.getItem("token");
  const headers: Record<string, string> = {
    "Content-Type": "application/json",
  };
  
  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }

  const response = await fetch("/api/v1/checkout", {
    method: "POST",
    headers,
    body: JSON.stringify(request),
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => null);
    const errorMessage = errorData?.detail || errorData?.title || "Checkout failed";
    throw new Error(errorMessage);
  }

  return response.json();
}
