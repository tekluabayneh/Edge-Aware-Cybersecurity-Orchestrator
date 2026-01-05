from fastapi import APIRouter
from  ..schemas.Telementory import RawTelemetryPayload
router = APIRouter(tags=["ingest"])


@router.post("/rawTelementory")
def get_activity_for_user(payload: RawTelemetryPayload):
    print("this is paylod man paylod",payload)
    print("get activity")
    return {"msg": "get data"}


