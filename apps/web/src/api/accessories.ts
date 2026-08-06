export interface AccessoryRecommendation {
  id: string;
  name: string;
  kind: "external" | "upgrade";
  category: string;
  brand: string;
  image_url: string;
  retailer_offer_id: string;
  retailer_product_id: string;
  source_url: string;
  price: number;
  original_price: number;
  in_stock: boolean;
  fetched_at: string;
  offer_state: "fresh" | "stale" | "unavailable";
  checkout_available: boolean;
  compatible_product_ids: string[];
  compatibility_reason: string;
  compatibility_status: "category" | "verified_model";
}

interface AccessoryRecommendationResponse {
  items: AccessoryRecommendation[];
}

export const getAccessoryRecommendations = async (
  productIds: string[]
): Promise<AccessoryRecommendation[]> => {
  const parameters = new URLSearchParams();
  for (const productId of productIds) {
    parameters.append("product_ids", productId);
  }
  const response = await fetch(
    `/api/v1/accessories/recommendations?${parameters.toString()}`
  );
  if (!response.ok) {
    throw new Error("Không thể tải phụ kiện từ Phong Vũ.");
  }
  const payload = (await response.json()) as AccessoryRecommendationResponse;
  return payload.items;
};
