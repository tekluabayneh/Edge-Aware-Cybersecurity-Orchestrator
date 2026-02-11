# from core.engine.pipeline import piplinejob
# from fastapi import APIRouter, Request
# from  ..schemas.Telementory import RawTelemetryPayload
# from core.context.analyzer_context import AnalyzerContext
# import json
# router = APIRouter(tags=["ingest"])
#
# @router.post("/rawTelementory")
# async def get_activity_for_user(payload: RawTelemetryPayload, req:Request):
#         payload_dict = {
#         "email": payload.email,
#         "agent_id": payload.agent_id,
#         "agent_token": payload.agent_token,
#         "processes": [
#             proc.model_dump(mode='json')   
#             for proc in payload.processes
#         ],
#         "network":   payload.network.model_dump(mode='json'),
#         "system":    payload.system.model_dump(mode='json'),
#         "integrity": payload.integrity.model_dump(mode='json'),
#         "security":  payload.security.model_dump(mode='json'),
#          }
#         context = AnalyzerContext()
#         user = {
#                 "email":payload.get("email"),
#                 "agent_id":payload.get("agent_id"),
#                 "agent_token":payload.get("agent_token"),
#                 }
#
#         context.user =  user
#
#         # print("Clean payload going to pipeline:\n" + json.dumps(payload_dict, indent=2))
#
#         piplinejob(payload_dict)

from datetime import datetime, timezone
from fastapi import APIRouter, Request, HTTPException
from core.engine.pipeline import piplinejob   
from core.context.analyzer_context import AnalyzerContext
from ..schemas.Telementory import RawTelemetryPayload
from shared.schemas.schema import Alert   
import uuid
from typing import Dict, Any
import json

router = APIRouter(tags=["ingest"])

@router.post("/rawTelementory")
async def ingest_raw_telemetry(payload: RawTelemetryPayload, req: Request):
    """
    Ingest raw telemetry from agent → normalize → analyze → return alerts if any
    """
    # 1. Prepare clean payload dictionary (good, you're already doing this)
    payload_dict: Dict[str, Any] = {
        "email": payload.email,
        "agent_id": payload.agent_id,
        "agent_token": payload.agent_token,
        "processes": [proc.model_dump(mode='json') for proc in payload.processes],
        "network": payload.network.model_dump(mode='json') if payload.network else {},
        "system": payload.system.model_dump(mode='json') if payload.system else {},
        "integrity": payload.integrity.model_dump(mode='json') if payload.integrity else {},
        "security": payload.security.model_dump(mode='json') if payload.security else {},
    }

    # 2. Create and populate AnalyzerContext
    context = AnalyzerContext()

    # Better request tracing
    context.request_id = str(uuid.uuid4())
    # context.start_time = req.state.start_time if hasattr(req.state, 'start_time') else None 
    context.start_time = getattr(req.state, "start_time", None)

    # Attach user/agent info (make it a proper dict or object if you want later)
    context.user = {
        "email": payload.email,
        "agent_id": payload.agent_id,
        "agent_token": payload.agent_token,  
        "ip": req.client.host,              
    }

    # Put the actual telemetry data into context (pipeline will use this)
    context.input_data = payload_dict
    context.raw_payload = payload.model_dump()  

    try:
        # 3. Run the pipeline — passing context instead of just dict
        piplinejob(context) 

        # 4. Prepare response based on what happened in the pipeline
        response = {
            "status": "success",
            "request_id": context.request_id,
            "agent_id": payload.agent_id,
            "email": payload.email,
            "processed_at": context.end_time.isoformat() if hasattr(context, 'end_time') else None,
            "alert_count": len(context.alerts),
            "has_alerts": len(context.alerts) > 0,
        }

        if context.alerts:
            response["alerts"] = [alert.to_dict() for alert in context.alerts]
            response["severity_summary"] = {
                "critical": sum(1 for a in context.alerts if a.severity == "critical"),
                "high": sum(1 for a in context.alerts if a.severity == "high"),
                "medium": sum(1 for a in context.alerts if a.severity == "medium"),
                "low": sum(1 for a in context.alerts if a.severity == "low"),
            }
        return response

    except Exception as e:
        context.logger.exception("Pipeline failed during ingest")
        error_response = {
            "status": "error",
            "request_id": context.request_id,
            "message": "Analysis failed",
            "detail": str(e),
        }

        raise HTTPException(status_code=500, detail=error_response)

