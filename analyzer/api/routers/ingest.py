from core.interfaces.input import Input
from fastapi import APIRouter, Request
from  ..schemas.Telementory import RawTelemetryPayload

router = APIRouter(tags=["ingest"])


@router.post("/rawTelementory")
async def get_activity_for_user(payload: RawTelemetryPayload, req:Request):
            Input(payload)
           
          


