from core.engine.pipeline import piplinejob
from fastapi import APIRouter, Request
from  ..schemas.Telementory import RawTelemetryPayload

router = APIRouter(tags=["ingest"])


@router.post("/rawTelementory")
async def get_activity_for_user(payload: RawTelemetryPayload, req:Request):
            piplinejob(payload)
           

