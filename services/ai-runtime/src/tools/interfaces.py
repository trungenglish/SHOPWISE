import httpx

from src.models.schemas import CatalogProduct, OfferComparison


class ToolProxy:
    def __init__(
        self,
        backend_url: str,
        transport: httpx.AsyncBaseTransport | None = None,
    ) -> None:
        self.backend_url = backend_url
        self.transport = transport

    async def catalog_search(self) -> list[CatalogProduct]:
        # ponytail: loads the demo catalog each turn; add filtering or caching when it grows.
        async with httpx.AsyncClient(
            base_url=self.backend_url,
            transport=self.transport,
            timeout=10,
        ) as client:
            response = await client.get("/api/v1/products")
            response.raise_for_status()

        payload = response.json()
        items = payload.get("items")
        if not isinstance(items, list):
            raise ValueError("catalog response must contain an items list")
        return [CatalogProduct.model_validate(item) for item in items]

    async def offer_comparison(self, product_id: str) -> OfferComparison:
        async with httpx.AsyncClient(
            base_url=self.backend_url,
            transport=self.transport,
            timeout=10,
        ) as client:
            response = await client.get(f"/api/v1/products/{product_id}/offer-comparison")
            response.raise_for_status()
        return OfferComparison.model_validate(response.json())
