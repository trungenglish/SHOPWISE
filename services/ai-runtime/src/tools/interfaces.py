import httpx

from src.models.schemas import CatalogProduct


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
