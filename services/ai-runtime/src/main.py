import time
from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse
from src.api.chat import router as chat_router
from src.models.errors import AIErrorState
from src.core.observability import logger

app = FastAPI(title="AI Runtime API")

@app.middleware("http")
async def add_process_time_header(request: Request, call_next):
    start_time = time.time()
    response = await call_next(request)
    process_time = time.time() - start_time
    
    # Log metadata only, no PII
    logger.info(
        "request_processed",
        extra={
            "path": request.url.path,
            "method": request.method,
            "status_code": response.status_code,
            "latency_ms": round(process_time * 1000, 2)
        }
    )
    
    response.headers["X-Process-Time"] = str(process_time)
    return response

@app.exception_handler(Exception)
async def global_exception_handler(request: Request, exc: Exception):
    error_response = AIErrorState(
        type="internal_error",
        message="An unexpected error occurred processing the request.",
        retryable=False
    )
    return JSONResponse(
        status_code=500,
        content=error_response.model_dump()
    )

app.include_router(chat_router, prefix="/api/v1")

@app.get("/health")
def health():
    return {"status": "ok"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("src.main:app", host="0.0.0.0", port=8000, reload=True)
