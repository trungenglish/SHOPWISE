import httpx

class ToolCaller:
    def __init__(self, base_url="http://localhost:8080/tools"):
        self.base_url = base_url

    def call_tool(self, tool_name: str, args: dict):
        # In a real implementation, this would make an HTTP request to the Go backend
        pass

