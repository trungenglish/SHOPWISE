# AI Services

## Local setup

```bash
python -m venv .venv

# unix shell
source .venv/Scripts/activate
# cmd
.\.venv\Scripts\activate.bat
# PowerShell
.\.venv\Scripts\Activate.ps1

python -m pip install -r requirements.txt
```

- Option, update lib into requirment.txt

```bash
python -m pip freeze > requirment.txt
```

## Run API

Starts FastAPI on port **8081** and the **gRPC `ai.v1.AiService`** server on **`grpc_port`** (default **50051**, same port Node uses when `AI_GRPC_HOST` points at this service).

```bash
fastapi dev --port 8081
```

- Options, run app with infisical. Before running, setup secret at <https://app.infisical.com/>

```bash
infisical run --env=dev -- fastapi dev --port 8081
```

Imports such as `configs` and `core` resolve from the `src/` directory (see `src/__init__.py`). If you run Uvicorn directly, set `PYTHONPATH=src` (or `src` on `sys.path`) the same way pytest does in `pyproject.toml`.

## Run worker

Requires Redis (monorepo: `pnpm db:start` exposes Redis on **6381**). Set `REDIS_URL=redis://localhost:6381/0` if not using the default.

```bash
python -m src.worker
```

## Run tests

```bash
pytest
```

## Docker
