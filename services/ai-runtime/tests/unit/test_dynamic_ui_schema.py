import json
from pathlib import Path

from jsonschema import Draft202012Validator

REPOSITORY_ROOT = Path(__file__).parents[4]
SCHEMA_PATH = REPOSITORY_ROOT / "packages/schemas/dynamic-ui-v1.schema.json"
FIXTURE_PATH = REPOSITORY_ROOT / "packages/schemas/fixtures/dynamic-ui-all-components.json"


def test_shared_dynamic_ui_fixture_matches_schema() -> None:
    schema = json.loads(SCHEMA_PATH.read_text(encoding="utf-8"))
    fixture = json.loads(FIXTURE_PATH.read_text(encoding="utf-8"))

    Draft202012Validator(schema).validate(fixture)


def test_dynamic_ui_schema_rejects_embedded_script_props() -> None:
    schema = json.loads(SCHEMA_PATH.read_text(encoding="utf-8"))
    fixture = json.loads(FIXTURE_PATH.read_text(encoding="utf-8"))
    fixture["ui_operations"][0]["component"]["children"][0]["props"] = {"script": "alert(1)"}

    errors = list(Draft202012Validator(schema).iter_errors(fixture))

    assert errors
