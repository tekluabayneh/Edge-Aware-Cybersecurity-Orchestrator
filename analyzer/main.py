from fastapi import FastAPI, Request
from api.api import api_router
from starlette.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse
from fastapi.exceptions import RequestValidationError



@api_router.get("/")
def check():
    print("check")
    return {"msg": "work"}

app = FastAPI(title="Analyzer API")
app.include_router(api_router, prefix="/api/v1")

@app.exception_handler(RequestValidationError)
async def validation_exception_handler(request: Request, res: RequestValidationError):
    print(res.body)
    return JSONResponse(
        status_code=422,
        content={
            "detail": res.errors(),
            "body": res.body,
        },
    )


app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )





