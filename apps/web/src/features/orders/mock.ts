import { Order } from "./types";

export const MOCK_ORDERS: Order[] = [
  {
    id: "ORD-7492-BX",
    date: "2026-07-10T14:30:00Z",
    status: "Confirmed",
    totalPrice: 29990000,
    items: [
      {
        id: "prod-1",
        name: "MacBook Air M3 13-inch 16GB 512GB",
        quantity: 1,
        price: 29990000,
        thumbnail: "https://images.unsplash.com/photo-1517336714731-489689fd1ca8?auto=format&fit=crop&q=80&w=200",
      },
    ],
  },
  {
    id: "ORD-6124-KL",
    date: "2026-07-05T09:15:00Z",
    status: "Delivered",
    totalPrice: 8590000,
    items: [
      {
        id: "prod-2",
        name: "Màn hình Dell UltraSharp U2723QE 27 inch 4K",
        quantity: 1,
        price: 8590000,
        thumbnail: "https://images.unsplash.com/photo-1527443224154-c4a3942d4aff?auto=format&fit=crop&q=80&w=200",
      },
    ],
  },
];

// Set to true to test the empty state
export const TEST_EMPTY_STATE = false;

export const getMockOrders = (): Order[] => {
  return TEST_EMPTY_STATE ? [] : MOCK_ORDERS;
};
