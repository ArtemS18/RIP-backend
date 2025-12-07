from fastapi import FastAPI
from app.schema import ResultRequest
from app.routers import router

app = FastAPI(root_path="/api")

@app.webhooks.put("send-results")
def send_results_webhook(body: ResultRequest):
    """Send available calculalation results"""

app.include_router(router)
