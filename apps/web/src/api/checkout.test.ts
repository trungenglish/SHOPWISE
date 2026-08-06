import { afterEach, expect, it, vi } from "vitest";

import { checkout } from "./checkout";

afterEach(() => vi.unstubAllGlobals());

it("sends idempotency header and exposes changed offer metadata", async () => {
  const fetchMock = vi.fn().mockResolvedValue({
    ok: false,
    json: async () => ({
      code: "OFFER_CHANGED",
      detail: "retailer offer price changed",
      offer: {
        retailer_offer_id: "aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa",
        price: 950_000,
        in_stock: true,
        fetched_at: "2026-08-05T12:00:00Z",
      },
    }),
  });
  vi.stubGlobal("fetch", fetchMock);

  await expect(
    checkout({
      customer_id: "11111111-1111-4111-8111-111111111111",
      items: [
        {
          retailer_offer_id: "aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa",
          quantity: 1,
        },
      ],
      fulfillment_method: "STORE_PICKUP",
      idempotency_key: "stable-key",
    })
  ).rejects.toMatchObject({
    code: "OFFER_CHANGED",
    offer: { price: 950_000 },
  });

  expect(fetchMock).toHaveBeenCalledWith(
    "/api/v1/checkout",
    expect.objectContaining({
      headers: expect.objectContaining({ "Idempotency-Key": "stable-key" }),
    })
  );
});
