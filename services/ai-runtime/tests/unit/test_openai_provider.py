import pytest
from unittest.mock import AsyncMock, MagicMock
from src.llm.openai_provider import OpenAIProvider

@pytest.mark.asyncio
async def test_openai_provider_chat_completion():
    mock_client = AsyncMock()
    mock_response = MagicMock()
    mock_response.choices = [MagicMock(message=MagicMock(content="Hello!"))]
    mock_client.chat.completions.create.return_value = mock_response
    
    provider = OpenAIProvider(api_key="test", client=mock_client)
    
    result = await provider.chat_completion(
        messages=[{"role": "user", "content": "Hi"}],
        model="gpt-4o-mini"
    )
    
    assert result == "Hello!"
    mock_client.chat.completions.create.assert_called_once()
