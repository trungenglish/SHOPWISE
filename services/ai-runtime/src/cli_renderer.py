import json
import sys
from typing import Any


def render_envelope(envelope: dict[str, Any]) -> str:
    """Render the canonical envelope as readable text for CLI consumers."""
    lines = [str(envelope.get("message", ""))]
    for operation in envelope.get("ui_operations", []):
        component = operation.get("component", {})
        _render_component(component, lines)
    return "\n".join(line for line in lines if line)


def _render_component(component: dict[str, Any], lines: list[str]) -> None:
    props = component.get("props", {})
    title = props.get("title") or props.get("message")
    if title:
        lines.append(str(title))
    for product in props.get("products", []):
        lines.append(f"- {product.get('name', 'Unknown')} — {product.get('price', 0):,} VND")
    for option in props.get("options", []):
        lines.append(f"[ ] {option.get('label', option.get('id', ''))}")
    for child in component.get("children", []):
        _render_component(child, lines)


def main() -> None:
    envelope = json.load(sys.stdin)
    print(render_envelope(envelope))


if __name__ == "__main__":
    main()
