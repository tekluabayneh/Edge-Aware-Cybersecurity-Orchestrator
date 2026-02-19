from fastapi import FastAPI, Request
from api.api import api_router
from starlette.middleware.cors import CORSMiddleware
from dotenv import load_dotenv
from api.routers.middleware import JWTAuthMiddleware
load_dotenv()   


@api_router.get("/")
def check():
    return {"status check": "Analyzer status is ok"}

app = FastAPI(title="Analyzer API")
app.include_router(api_router, prefix="/api/v1")
app.add_middleware(JWTAuthMiddleware)
app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )





