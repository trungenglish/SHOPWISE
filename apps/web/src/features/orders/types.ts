export type OrderStatus = "Confirmed" | "Processing" | "Delivered";

export interface OrderItem {
  id: string;
  name: string;
  quantity: number;
  price: number;
  thumbnail: string;
}

export interface Order {
  id: string;
  date: string;
  status: OrderStatus;
  totalPrice: number;
  items: OrderItem[];
}
