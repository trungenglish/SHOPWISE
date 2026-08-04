import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import type { ReactNode } from "react";
import { afterEach, expect, it, vi } from "vitest";

import type { Laptop } from "../types";
import AccessoriesModal from "./accessories-modal";

const laptops: Laptop[] = [
  {
    id: "11111111-1111-4111-8111-111111111111",
    name: "Laptop A",
    price: 1,
    image: "",
    matchScore: 90,
    specs: { gpu: "", ram: "", cooling: "" },
  },
  {
    id: "22222222-2222-4222-8222-222222222222",
    name: "Laptop B",
    price: 2,
    image: "",
    matchScore: 80,
    specs: { gpu: "", ram: "", cooling: "" },
  },
];

const response = {
  items: [
    {
      id: "a",
      name: "Phong Vu mouse",
      kind: "external",
      category: "mouse",
      brand: "Logitech",
      image_url: "",
      retailer_offer_id: "aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa",
      retailer_product_id: "PV-1",
      source_url: "https://phongvu.vn/p/mouse",
      price: 900000,
      original_price: 0,
      in_stock: true,
      fetched_at: "2026-08-05T12:00:00Z",
      offer_state: "fresh",
      checkout_available: true,
      compatible_product_ids: [laptops[0].id],
      compatibility_reason: "Gaming fit",
      compatibility_status: "category",
    },
    {
      id: "b",
      name: "Phong Vu dock",
      kind: "external",
      category: "dock",
      brand: "Ugreen",
      image_url: "",
      retailer_offer_id: "bbbbbbbb-bbbb-4bbb-8bbb-bbbbbbbbbbbb",
      retailer_product_id: "PV-2",
      source_url: "https://phongvu.vn/p/dock",
      price: 1200000,
      original_price: 0,
      in_stock: true,
      fetched_at: "2026-08-05T10:00:00Z",
      offer_state: "stale",
      checkout_available: false,
      compatible_product_ids: [laptops[1].id],
      compatibility_reason: "Business fit",
      compatibility_status: "category",
    },
  ],
};

const paginatedResponse = {
  items: Array.from({ length: 13 }, (_, index) => ({
    ...response.items[0],
    id: `item-${index + 1}`,
    name: `Accessory ${index + 1}`,
  })),
};

const wrapper = ({ children }: { children: ReactNode }) => (
  <QueryClientProvider
    client={new QueryClient({ defaultOptions: { queries: { retry: false } } })}
  >
    {children}
  </QueryClientProvider>
);

afterEach(() => vi.unstubAllGlobals());

it("loads all recommendations then filters by laptop", async () => {
  vi.stubGlobal(
    "fetch",
    vi.fn().mockResolvedValue({ ok: true, json: async () => response })
  );
  render(
    <AccessoriesModal
      isOpen
      onClose={() => undefined}
      products={laptops}
      onAddToCheckout={() => undefined}
    />,
    { wrapper }
  );

  expect(await screen.findByText("Phong Vu mouse")).toBeInTheDocument();
  expect(screen.getByText("Phong Vu dock")).toBeInTheDocument();
  fireEvent.change(screen.getByLabelText("Lọc theo laptop"), {
    target: { value: laptops[0].id },
  });

  await waitFor(() =>
    expect(screen.queryByText("Phong Vu dock")).not.toBeInTheDocument()
  );
  expect(screen.getByText("Phong Vu mouse")).toBeInTheDocument();
});

it("disables checkout for stale offers", async () => {
  vi.stubGlobal(
    "fetch",
    vi.fn().mockResolvedValue({ ok: true, json: async () => response })
  );
  render(
    <AccessoriesModal
      isOpen
      onClose={() => undefined}
      products={laptops}
      onAddToCheckout={() => undefined}
    />,
    { wrapper }
  );

  expect(await screen.findByText("Phong Vu dock")).toBeInTheDocument();
  expect(
    screen.getAllByRole("button", { name: "Thêm vào checkout" })[1]
  ).toBeDisabled();
  expect(screen.getByText("Dữ liệu giá đã cũ")).toBeInTheDocument();
});

it("renders twelve accessories per page and navigates with previous and next", async () => {
  vi.stubGlobal(
    "fetch",
    vi.fn().mockResolvedValue({ ok: true, json: async () => paginatedResponse })
  );
  render(
    <AccessoriesModal
      isOpen
      onClose={() => undefined}
      products={laptops}
      onAddToCheckout={() => undefined}
    />,
    { wrapper }
  );

  expect(await screen.findByText("Accessory 1")).toBeInTheDocument();
  expect(screen.getByText("Accessory 12")).toBeInTheDocument();
  expect(screen.queryByText("Accessory 13")).not.toBeInTheDocument();
  expect(screen.getByText("Trang 1/2")).toBeInTheDocument();

  fireEvent.click(screen.getByRole("button", { name: "Sau" }));

  expect(screen.getByText("Accessory 13")).toBeInTheDocument();
  expect(screen.queryByText("Accessory 1")).not.toBeInTheDocument();
  expect(screen.getByText("Trang 2/2")).toBeInTheDocument();
  expect(screen.getByRole("button", { name: "Sau" })).toBeDisabled();
});

it("returns to the first page and scrolls to the top when a filter changes", async () => {
  vi.stubGlobal(
    "fetch",
    vi.fn().mockResolvedValue({ ok: true, json: async () => paginatedResponse })
  );
  render(
    <AccessoriesModal
      isOpen
      onClose={() => undefined}
      products={laptops}
      onAddToCheckout={() => undefined}
    />,
    { wrapper }
  );

  expect(await screen.findByText("Accessory 1")).toBeInTheDocument();
  fireEvent.click(screen.getByRole("button", { name: "Sau" }));
  const navigation = screen.getByRole("navigation", {
    name: "Phân trang phụ kiện",
  });
  const list = navigation.previousElementSibling as HTMLDivElement;
  list.scrollTop = 100;

  fireEvent.change(screen.getByLabelText("Lọc theo laptop"), {
    target: { value: laptops[0].id },
  });

  expect(screen.getByText("Trang 1/2")).toBeInTheDocument();
  expect(screen.getByText("Accessory 1")).toBeInTheDocument();
  expect(list.scrollTop).toBe(0);
});
