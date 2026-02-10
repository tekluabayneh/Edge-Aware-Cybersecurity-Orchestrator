def set_integrity_rule(event): 
    payload = event.get("payload", {})
    kernel_status  =  payload.get("kernel_status")
    patch_status  = payload.get("patch_status")
    integrity_status = payload.get("integrity_status")
    overall_integrity = payload.get("overall_integrity")


    Rules = []

    if kernel_status is not "Stable":
        Rules.append({"kernel_rule":"kernel is not update to you must update your kernel"})
    else: 
        Rules.append({"kernel_rule":"kernel is up to date"})
    if patch_status is not "Fully Patched":
        Rules.append({"patch_rule":"patch is not update to you must update your patch"})
    else: 
        Rules.append({"patch_rule":"Patch is update to date"})

    if integrity_status is not "Clean": 
        Rules.append({"integrity_rule":"integrity is not in Clean state you must update your machine"})
    else: 
        Rules.append({"integrity_rule":"integrity is clean and  update to date"})

    if overall_integrity is not "Healthy": 
        Rules.append({"overall_Rule":"overall integrity is not Healthy so make sure to update any software you have and free up unused risky file" })
    else: 
        Rules.append({"overall_Rule":"overall integrity is Healthy"})

    payload["Rules"] = Rules
      
    return event

