from datetime import datetime, timezone
from fastapi import APIRouter, Request, HTTPException
from core.engine.pipeline import piplinejob   
from core.context.analyzer_context import AnalyzerContext
from ..schemas.Telementory import RawTelemetryPayload
from shared.schemas.schema import Alert   
from core.interfaces.alert_handler import send_alert
import uuid
from typing import Dict, Any
import json

router = APIRouter(tags=["ingest"])

@router.post("/rawTelementory")
async def ingest_raw_telemetry(payload: RawTelemetryPayload, req: Request):
    """
    Ingest raw telemetry from agent → normalize → analyze → return alerts if any
    """ 

    header = req.headers.get("Authorization")
    if not header: 
        print("heaer is empty")
        return

    token = header.split(" ")[1]
    print(token)

    # 1. Prepare clean payload dictionary (good, you're already doing this)
    payload_dict: Dict[str, Any] = {
        "email": payload.email,
        "agent_id": payload.agent_id,
        "agent_token": payload.agent_token,
        "machine_id": payload.machine_id,
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
    context.start_time = getattr(req.state, "start_time", None)

    # Attach user/agent info (make it a proper dict or object if you want later)
    context.user = {
        "token": token,
        "email": payload.email,
        "agent_id": payload.agent_id,
        "agent_token": payload.agent_token,  
        "machine_id": payload.machine_id,  
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

        alert_to_send = []
        if context.alerts:
            for alert in context.alerts:
                    duplicate = any(
                        a["message"] == alert["message"] and
                        a["severity"] == alert["severity"] and
                        a["alert_type"] == alert["alert_type"] and
                        a["risk_level"] == alert["risk_level"]
                        for a in alert_to_send
                                )
                    if duplicate: 
                         continue 
                    else: 
                        alert["agent_id"] = payload.agent_token 
                        alert["agent_token"] = payload.agent_token
                        alert["machine_id"] = payload.machine_id
                        alert["email"] = payload.email 
                        alert_to_send.append(alert)

        for uniqe_alert in alert_to_send: 
            send_alert(uniqe_alert, token)

    except Exception as e:
        context.logger.exception("Pipeline failed during ingest")
        error_response = {
            "status": "error",
            "request_id": context.request_id,
            "message": "Analysis failed",
            "detail": str(e),
        }

        raise HTTPException(status_code=500, detail=error_response)

