from fastapi import APIRouter
from api.routers import activity, users, dashboard

api_router = APIRouter()

api_router.include_router(activity.router)
api_router.include_router(users.router)
api_router.include_router(dashboard.router)
