from datetime import datetime, timezone
def set_integrity_rule(event, alert): 
    payload = event.get("payload", {})
    kernel_status  =  payload.get("kernel_status")
    patch_status  = payload.get("patch_status")
    integrity_status = payload.get("integrity_status")
    overall_integrity = payload.get("overall_integrity")


    Rules = []

    if kernel_status != "Stable":
        Rules.append({"kernel_rule":"kernel is not update to you must update your kernel"})
    else: 
        Rules.append({"kernel_rule":"kernel is up to date"})
    if patch_status != "Fully Patched":
        Rules.append({"patch_rule":"patch is not update to you must update your patch"})
    else: 
        Rules.append({"patch_rule":"Patch is update to date"})

    if integrity_status != "Clean": 
        Rules.append({"integrity_rule":"integrity is not in Clean state you must update your machine"})
    else: 
        Rules.append({"integrity_rule":"integrity is clean and  update to date"})

    if overall_integrity != "Healthy": 
        Rules.append({"overall_Rule":"overall integrity is not Healthy so make sure to update any software you have and free up unused risky file" })
    else: 
        Rules.append({"overall_Rule":"overall integrity is Healthy"})

    payload["Rules"] = Rules

    if event.get("overall_integrity") in ["High Risk","Warning","Attention Needed"]: 
        alert.append({
            "alert_type": "integrity",
            "severity": "high",
            "message": "System integrity is degraded. Kernel, patches, or system state require attention.",
            "raw_payload":"" ,
            "status": "open",
            "risk_level": "high",
            "summary": "SYSTEM_INTEGRITY_ISSUE",
            "performance":"" ,
            "network":"" ,
            "security":"", 
            "created_at": datetime.now(timezone.utc).isoformat()
        })

      
    return event


