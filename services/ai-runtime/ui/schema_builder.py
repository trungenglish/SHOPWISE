class UISchemaBuilder:
    def build_carousel(self, products: list, rationale: dict):
        return {
            "type": "carousel",
            "items": []
        }
    
    def build_comparison_table(self, products: list, features: list):
        return {
            "type": "comparison_table",
            "products": [],
            "features": []
        }

