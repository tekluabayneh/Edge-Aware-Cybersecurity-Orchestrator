import json
import time


def intgerity_enrich(event): 
    payload = event.get("payload",{})
    kernel_version = payload.get("kernel_version")

    if kernel_version:
      try: 

        major_kernel = int(kernel_version.split(".")[0])
        if major_kernel >= 6:
            event["kernel_status"] = "Stable"
        elif major_kernel == 5:
            event["kernel_status"] = "Acceptable"
        else:
            event["kernel_status"] = "Outdated"
      except ValueError:
            event["kernel_status"] = "Unknown"

      patch_level = payload.get("patch_level")
      if patch_level:
        try: 
            major_patch = int(patch_level.split(".")[0])

            if major_patch >= 24:
                event["patch_status"] = "Fully Patched"
            elif major_patch >= 22:
                event["patch_status"] = "Mostly Patched"
            elif major_patch >= 20:
                event["patch_status"] = "Partially Patched"
            else:
                event["patch_status"] = "End-of-Life"
        except ValueError:
            event["patch_status"] = "Unknown"



    critical_files = payload.get("critical_files", {})
    now = time.time()
    recent_changes = 0 

    for key, timestamp in critical_files.items(): 

        try:
            file_time = time.mktime( 
                time.strptime(timestamp.split(".")[0], "%Y-%m-%dT%H:%M:%S")
                ) 

            if (now - file_time) < (30 * 24 * 60 * 60):  
                recent_changes += 1
        except Exception: 
           continue


    if recent_changes == 0:
        event["integrity_status"] = "Clean"
    elif recent_changes == 1:
        event["integrity_status"] = "Expected Changes"
    elif recent_changes <= 3:
        event["integrity_status"] = "Suspicious"
    else:
        event["integrity_status"] = "Compromised"

   
    statuses = [
        event.get("kernel_status"),
        event.get("patch_status"),
        event.get("integrity_status"),
        ]   
 


    if "Compromised" in statuses or "Outdated" in statuses:
        event["overall_integrity"] = "High Risk"
    elif "Suspicious" in statuses:
        event["overall_integrity"] = "Warning"
    elif "Acceptable" in statuses or "Mostly Patched" in statuses:
        event["overall_integrity"] = "Attention Needed"
    else:
        event["overall_integrity"] = "Healthy"

    print("Clean payload going to pipeline:\n" + json.dumps(event, indent=2))
    return event



