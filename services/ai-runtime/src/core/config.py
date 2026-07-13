from pydantic_settings import BaseSettings, SettingsConfigDict
from pydantic import Field, SecretStr
from typing import Literal

class Settings(BaseSettings):
    # LLM Provider Configuration
    llm_provider: Literal["openai", "azure", "anthropic", "gemini", "local"] = "openai"
    llm_api_key: SecretStr | None = Field(default=None, description="API key for the selected LLM provider")
    llm_base_url: str | None = Field(default=None, description="Optional base URL for the LLM provider (e.g., for Azure or proxy)")
    
    # Server Configuration
    host: str = "0.0.0.0"
    port: int = 8000
    environment: Literal["development", "staging", "production"] = "development"

    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        extra="ignore"
    )

settings = Settings()
