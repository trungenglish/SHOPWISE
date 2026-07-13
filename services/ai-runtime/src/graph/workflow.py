from langgraph.graph import StateGraph, START, END
from typing import Any, TypedDict
import json

class GraphState(TypedDict):
    messages: list[dict[str, Any]]
    session_id: str
    retry_count: int
    final_response: str | None
    error: str | None

def create_workflow(provider: Any):
    workflow = StateGraph(GraphState)  # type: ignore
    
    async def llm_node(state: GraphState):
        from src.models.schemas import DecisionResponse
        schema_dict = DecisionResponse.model_json_schema()
        response_format = {
            "type": "json_schema",
            "json_schema": {
                "name": "DecisionResponse",
                "schema": schema_dict,
                "strict": False
            }
        }
        response = await provider.chat_completion(
            messages=state["messages"],
            model="gpt-4o-mini",
            session_id=state["session_id"],
            response_format=response_format
        )
        return {"final_response": response}
        
    async def parse_json_node(state: GraphState):
        try:
            if state["final_response"]:
                # Attempt to parse
                json.loads(state["final_response"])
                return {"error": None}
            return {"error": "empty_response"}
        except json.JSONDecodeError as e:
            if state.get("retry_count", 0) < 2:
                # Append error message back to LLM to self-correct
                return {"error": str(e), "retry_count": state.get("retry_count", 0) + 1}
            return {"error": "max_retries_exceeded"}
            
    workflow.add_node("llm", llm_node)
    workflow.add_node("parse_json", parse_json_node)
    
    workflow.add_edge(START, "llm")
    workflow.add_edge("llm", "parse_json")
    
    def routing_logic(state: GraphState):
        if state["error"] == "max_retries_exceeded":
            return END
        elif state["error"]:
            return "llm"
        return END
        
    workflow.add_conditional_edges("parse_json", routing_logic)
    
    return workflow.compile()
