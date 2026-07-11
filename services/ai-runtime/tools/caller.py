import httpx

class ToolCaller:
    def __init__(self, base_url="http://localhost:8080/tools"):
        self.base_url = base_url

    def call_tool(self, tool_name: str, args: dict):
        response = httpx.post(f"{self.base_url}/{tool_name}", json=args)
        response.raise_for_status()
        return response.json()
