import { Order } from "./types";

export const MOCK_ORDERS: Order[] = [
  {
    id: "ORD-7492-BX",
    date: "2026-07-10T14:30:00Z",
    status: "Confirmed",
    totalPrice: 38500000,
    items: [
      {
        id: "prod-1",
        name: "ROG Strix G16 (RTX 5057)",
        quantity: 1,
        price: 38500000,
        thumbnail: "https://images.unsplash.com/photo-1603302576837-37561b2e2302?w=600&q=80",
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
        thumbnail: "https://images.unsplash.com/photo-1603302576837-37561b2e2302?w=600&q=80",
      },
    ],
  },
];

// Set to true to test the empty state
export const TEST_EMPTY_STATE = false;

export const getMockOrders = (): Order[] => {
  return TEST_EMPTY_STATE ? [] : MOCK_ORDERS;
};
