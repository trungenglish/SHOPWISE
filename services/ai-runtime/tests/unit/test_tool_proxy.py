import httpx
import pytest

from src.tools.interfaces import ToolProxy


@pytest.mark.asyncio
async def test_catalog_search_returns_validated_products():
    def handler(request: httpx.Request) -> httpx.Response:
        assert request.url.path == "/api/v1/products"
        return httpx.Response(
            200,
            json={
                "items": [
                    {
                        "ID": "product-1",
                        "Name": "Catalog Laptop",
                        "Price": 37_475_000,
                        "Specifications": {},
                        "Metadata": {},
                    }
                ]
            },
        )

    proxy = ToolProxy(
        backend_url="http://backend.test",
        transport=httpx.MockTransport(handler),
    )

    products = await proxy.catalog_search()

    assert products[0].id == "product-1"
    assert products[0].price == 37_475_000


@pytest.mark.asyncio
async def test_catalog_search_does_not_fallback_when_backend_fails():
    def handler(_request: httpx.Request) -> httpx.Response:
        return httpx.Response(503, json={"error": "catalog unavailable"})

    proxy = ToolProxy(
        backend_url="http://backend.test",
        transport=httpx.MockTransport(handler),
    )

    with pytest.raises(httpx.HTTPStatusError):
        await proxy.catalog_search()
