from langgraph.graph import StateGraph, END
from typing import TypedDict, List, Dict, Any
import sys
import os

# Add parent directory to path to import tools
sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
from tools.fetch_comparison_data import fetch_comparison_data
from prompts.comparison_reasoning import COMPARISON_REASONING_PROMPT

class AgentState(TypedDict):
    session_id: str
    messages: List[Dict[str, Any]]
    context: Dict[str, Any]
    next_node: str

def tool_node(state: AgentState):
    """Placeholder tool node that executes fetch_comparison_data if requested."""
    # Simplified implementation for the sake of the task
    last_msg = state.get("messages", [{}])[-1]
    if last_msg.get("type") == "tool_call" and last_msg.get("name") == "fetch_comparison_data":
        args = last_msg.get("args", {})
        result = fetch_comparison_data(args.get("product_ids", []))
        new_msgs = state.get("messages", []).copy()
        new_msgs.append({"type": "tool_response", "content": result})
        return {"messages": new_msgs}
    return state

def reasoner_node(state: AgentState):
    # Inject comparison reasoning prompt if doing comparison
    return state

def build_graph():
    workflow = StateGraph(AgentState)
    
    # Add nodes (placeholders for now)
    workflow.add_node("planner", lambda state: state)
    workflow.add_node("reasoner", reasoner_node)
    workflow.add_node("tools", tool_node)
    
    workflow.set_entry_point("planner")
    workflow.add_edge("planner", "reasoner")
    workflow.add_edge("reasoner", "tools")
    workflow.add_edge("tools", END)
    
    return workflow.compile()

