from fastapi import APIRouter
from  ..schemas.Telementory import TelemetryPayload
router = APIRouter(tags=["ingest"])


@router.post("/rawTelementory")
def get_activity_for_user(payload: TelemetryPayload):
    print(payload)
    print("get activity")
    return {"msg": "get data"}


