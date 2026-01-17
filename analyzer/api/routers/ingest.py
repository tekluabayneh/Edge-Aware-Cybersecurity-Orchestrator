from core.engine.pipeline import piplinejob
from fastapi import APIRouter, Request
from  ..schemas.Telementory import RawTelemetryPayload
import json

router = APIRouter(tags=["ingest"])


@router.post("/rawTelementory")
async def get_activity_for_user(payload: RawTelemetryPayload, req:Request):
        payload_dict = {
        "email": payload.email,
        "agent_id": payload.agent_id,
        "agent_token": payload.agent_token,
        "processes": [
            proc.model_dump(mode='json')   
            for proc in payload.processes
        ],
        "network":   payload.network.model_dump(mode='json'),
        "system":    payload.system.model_dump(mode='json'),
        "integrity": payload.integrity.model_dump(mode='json'),
        "security":  payload.security.model_dump(mode='json'),
         }

        # print("Clean payload going to pipeline:\n" + json.dumps(payload_dict, indent=2))

        piplinejob(payload_dict)
           

