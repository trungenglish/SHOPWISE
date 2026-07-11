from tools.caller import ToolCaller

def fetch_comparison_data(product_ids: list[str]) -> dict:
    """
    Retrieves comparison data for a list of products.
    """
    caller = ToolCaller()
    return caller.call_tool("fetch_comparison_data", {"product_ids": product_ids})
