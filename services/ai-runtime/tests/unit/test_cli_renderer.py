import json
from pathlib import Path

from src.cli_renderer import render_envelope


def test_cli_renderer_uses_the_same_dynamic_ui_envelope() -> None:
    fixture_path = (
        Path(__file__).parents[4] / "packages/schemas/fixtures/dynamic-ui-all-components.json"
    )
    rendered = render_envelope(json.loads(fixture_path.read_text(encoding="utf-8")))

    assert "Dynamic UI fixture" in rendered
    assert "Promotion" in rendered
