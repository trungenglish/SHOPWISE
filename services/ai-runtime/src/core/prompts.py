class PromptManager:
    """
    Manages system prompts, context assembly, and templates.
    """
    
    BASE_SYSTEM_PROMPT = """You are the SHOPWISE AI retail assistant.
Your goal is to help users compare products, answer questions, and assist with checkout.
Do not invent information. Follow instructions carefully.
You MUST output your response matching the structured DecisionResponse format.
Provide products, reasoning, agents involved in the analysis, audit logs, and recommended accessories.
"""

    @classmethod
    def get_system_prompt(cls, context: dict | None = None) -> str:
        prompt = cls.BASE_SYSTEM_PROMPT
        if context:
            prompt += "\n\nAdditional Context:\n"
            for k, v in context.items():
                prompt += f"{k}: {v}\n"
        return prompt
