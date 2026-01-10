from fastapi import FastAPI, Request
from api.api import api_router
from starlette.middleware.cors import CORSMiddleware


@api_router.get("/")
def check():
    print("check")
    return {"msg": "work"}

app = FastAPI(title="Analyzer API")
app.include_router(api_router, prefix="/api/v1")
app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )





