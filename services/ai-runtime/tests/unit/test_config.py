from src.core.config import Settings


def test_settings_use_openai_key_and_model_environment_names(monkeypatch):
    monkeypatch.setenv("OPENAI_API_KEY", "test-key")
    monkeypatch.setenv("MODEL", "gpt-5.4-mini")

    configuration = Settings(_env_file=None)

    assert configuration.openai_api_key.get_secret_value() == "test-key"
    assert configuration.model == "gpt-5.4-mini"
