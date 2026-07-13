from fastapi import APIRouter, HTTPException, Depends
from sse_starlette.sse import EventSourceResponse
from pydantic import BaseModel, Field

from src.core.config import settings
from src.llm.openai_provider import OpenAIProvider
from src.core.prompts import PromptManager
from src.api.streaming import generate_sse
from src.graph.workflow import create_workflow

router = APIRouter()

class ChatRequest(BaseModel):
    session_id: str = Field(..., description="Unique session identifier")
    message: str = Field(..., description="User message content")
    # Runtime Model Configuration
    model: str | None = Field(default=None, description="Model override")
    temperature: float | None = Field(default=None, description="Temperature override")
    max_tokens: int | None = Field(default=None, description="Max tokens override")

class ChatResponse(BaseModel):
    response: str

def get_provider():
    if not settings.llm_api_key:
        raise HTTPException(status_code=500, detail="LLM API key not configured")
    
    if settings.llm_provider == "openai":
        return OpenAIProvider(
            api_key=settings.llm_api_key.get_secret_value(),
            base_url=settings.llm_base_url
        )
    raise HTTPException(status_code=500, detail=f"Provider {settings.llm_provider} not supported")

@router.post("/chat", response_model=ChatResponse)
async def chat_endpoint(req: ChatRequest, provider: OpenAIProvider = Depends(get_provider)):
    # Merge runtime params with defaults
    model = req.model or "gpt-4o-mini"
    temperature = req.temperature if req.temperature is not None else 0.7
    max_tokens = req.max_tokens
    
    system_prompt = PromptManager.get_system_prompt()
    messages = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": req.message}
    ]
    
    try:
        # Wire LangGraph workflow
        workflow = create_workflow(provider)
        state = {
            "messages": messages,
            "session_id": req.session_id,
            "retry_count": 0,
            "final_response": None,
            "error": None
        }
        
        result = await workflow.ainvoke(state)
        
        if result.get("error"):
            raise HTTPException(status_code=500, detail=f"Workflow error: {result['error']}")
            
        return ChatResponse(response=result.get("final_response", ""))
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

from src.models.schemas import DecisionResponse

@router.post("/chat/stream")
async def chat_stream_endpoint(req: ChatRequest, provider: OpenAIProvider = Depends(get_provider)):
    # Merge runtime params with defaults
    model = req.model or "gpt-4o-mini"
    temperature = req.temperature if req.temperature is not None else 0.7
    max_tokens = req.max_tokens
    
    system_prompt = PromptManager.get_system_prompt()
    messages = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": req.message}
    ]
    
    schema_dict = DecisionResponse.model_json_schema()
    response_format = {
        "type": "json_schema",
        "json_schema": {
            "name": "DecisionResponse",
            "schema": schema_dict,
            "strict": False
        }
    }
    
    try:
        generator = provider.stream_chat_completion(
            messages=messages,
            model=model,
            temperature=temperature,
            max_tokens=max_tokens,
            session_id=req.session_id,
            response_format=response_format
        )
        return EventSourceResponse(generate_sse(generator))
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
