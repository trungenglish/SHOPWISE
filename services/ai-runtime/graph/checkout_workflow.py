from typing import Dict, Any, Optional
from tools.checkout_tools import validate_checkout_readiness

def get_checkout_readiness_node(state: Dict[str, Any]) -> Dict[str, Any]:
    """
    Checkout Readiness workflow node.
    """
    session_id = state.get("session_id", "default-session")
    product_id = state.get("product_id", "default-product")
    
    readiness_state = validate_checkout_readiness(session_id, product_id)
    
    if readiness_state.get("status") == "READY":
        return {
            "ui": "CheckoutSummary",
            "state": readiness_state
        }
    
    if readiness_state.get("status") == "MISSING_INFORMATION":
        return {
            "ui": "CustomerInformationForm",
            "state": readiness_state
        }
        
    status = readiness_state.get("status")
    if status in ["OUT_OF_STOCK", "STORE_UNAVAILABLE", "INVENTORY_CHANGED"]:
        # Explain inventory changes to the user and suggest alternatives
        explanation = f"We're sorry, but the inventory status changed: {status}."
        return {
            "ui": "CheckoutSummary",
            "state": readiness_state,
            "explanation": explanation
        }
        
    if status == "REQUIRES_USER_ACTION":
        return {
            "ui": "ErrorStateUI",
            "state": readiness_state,
            "explanation": "Action is required before you can proceed."
        }
        
    if status == "PROMOTION_CHANGED":
        return {
            "ui": "PromotionSummary",
            "state": readiness_state,
            "explanation": "Your promotion has expired or changed."
        }
        
    return {
        "ui": "ErrorStateUI",
        "state": readiness_state
    }
