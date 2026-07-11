from langgraph.graph import StateGraph, END
from typing import TypedDict, List, Dict, Any

class AgentState(TypedDict):
    session_id: str
    messages: List[Dict[str, Any]]
    context: Dict[str, Any]
    next_node: str

def build_graph():
    workflow = StateGraph(AgentState)
    
    # Add nodes (placeholders for now)
    workflow.add_node("planner", lambda state: state)
    workflow.add_node("reasoner", lambda state: state)
    
    workflow.set_entry_point("planner")
    workflow.add_edge("planner", "reasoner")
    workflow.add_edge("reasoner", END)
    
    return workflow.compile()

