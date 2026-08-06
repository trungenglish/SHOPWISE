export interface CheckoutItemRequest {
  product_id?: string;
  retailer_offer_id?: string;
  quantity: number;
}

export interface CheckoutRequest {
  customer_id: string;
  items: CheckoutItemRequest[];
  fulfillment_method: "DELIVERY" | "STORE_PICKUP";
  shipping_address?: string;
  coupon_code?: string;
  idempotency_key?: string;
}

export interface OrderItemResponse {
  product_id?: string;
  retailer_offer_id?: string;
  name?: string;
  source_url?: string;
  verified_at?: string;
  quantity: number;
  unit_price: number;
}

interface CheckoutErrorPayload {
  code?: string;
  detail?: string;
  title?: string;
  offer?: {
    retailer_offer_id: string;
    price: number;
    in_stock: boolean;
    fetched_at: string;
  };
}

export class CheckoutError extends Error {
  readonly code?: string;
  readonly offer?: CheckoutErrorPayload["offer"];

  constructor(payload: CheckoutErrorPayload | null) {
    super(payload?.detail ?? payload?.title ?? "Checkout failed");
    this.name = "CheckoutError";
    this.code = payload?.code;
    this.offer = payload?.offer;
  }
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
	estimated_delivery_from: string;
	estimated_delivery_to: string;
	confirmation_email_status: "queued" | "failed";
}

export async function checkout(
  request: CheckoutRequest
): Promise<OrderResponse> {
  const token = localStorage.getItem("token");
  const headers: Record<string, string> = {
    "Content-Type": "application/json",
  };

  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }

  const { idempotency_key: idempotencyKey, ...payload } = request;
  if (idempotencyKey) {
    headers["Idempotency-Key"] = idempotencyKey;
  }

  const response = await fetch("/api/v1/checkout", {
    method: "POST",
    headers,
    body: JSON.stringify(payload),
  });

  if (!response.ok) {
    const errorData = (await response
      .json()
      .catch(() => null)) as CheckoutErrorPayload | null;
    throw new CheckoutError(errorData);
  }

  return response.json();
}
