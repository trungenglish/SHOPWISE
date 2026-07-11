# AI Runtime

`ai-runtime` is the AI orchestration service for the SHOPWISE platform, built with **FastAPI**, **LangGraph**, and **LangChain**. It provides AI-driven workflows and tools, such as the checkout workflow, to power intelligent features within the application.

## Prerequisites

- **Python**: 3.11 or higher
- **Dependency Manager**: [uv](https://github.com/astral-sh/uv)

## Installation

This project uses `uv` for fast dependency management. To install the dependencies, run:

```bash
uv sync
```

## Running the Development Server

To start the FastAPI development server, run:

```bash
uv run uvicorn main:app --reload --host 0.0.0.0 --port 8000
```

Alternatively, you can just run `main.py` directly:

```bash
uv run python main.py
```

The server will be available at `http://localhost:8000`.
You can check its health at `http://localhost:8000/health`.

## Project Structure

- `main.py`: The entry point for the FastAPI application.
- `graph/`: Contains LangGraph workflows (e.g., `checkout_workflow.py`, `workflow.py`).
- `tools/`: Contains tools used by the AI workflows (e.g., `checkout_tools.py`, `fetch_comparison_data.py`).
- `engine/`: Core execution logic and model orchestration.
- `nodes/`: Custom LangGraph nodes used in the state graphs.
- `prompts/`: System prompts and prompt templates for the LLMs.
- `ui/`: Potential frontend components or test UI code.
