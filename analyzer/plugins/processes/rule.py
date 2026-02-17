from datetime import datetime, timezone
def set_proccess_rule(event, alert): 

    cpu_load_level = event.get("cpu_load_level")
    memory_load_level = event.get("memory_load_level")

    proccess_rules = []
    memory_rule = []
    cpu_rule = []

    # MEMORY RULES
    if memory_load_level in ["Chill", "Light"]:
        memory_rule.append({"status":"ok", "message":"Memory load is light."})
    elif memory_load_level in ["Normal", "Medium"]:
        memory_rule.append({"status":"warning", "message":"Memory load is normal."})
    elif memory_load_level in ["Heavy", "Intense", "Crazy"]:
        memory_rule.append({"status":"critical", "message":"Memory load is high!"})
    else:
        memory_rule.append({"status":"unknown", "message":"Memory load unknown."})

    # CPU RULES
    if cpu_load_level in ["Chill", "Light"]:
        cpu_rule.append({"status":"ok", "message":"CPU load is light."})
    elif cpu_load_level in ["Normal", "Medium"]:
        cpu_rule.append({"status":"warning", "message":"CPU load is normal."})
    elif cpu_load_level in ["Heavy", "Intense", "Crazy"]:
        cpu_rule.append({"status":"critical", "message":"CPU load is high!"})
    else:
        cpu_rule.append({"status":"unknown", "message":"CPU load unknown."})

    # Combine rules
    proccess_rules.extend(memory_rule)
    proccess_rules.extend(cpu_rule)


    if event.get("cpu_load_level") in ["Heavy", "Intense", "Crazy"] or event.get("memory_load_level") in ["Heavy", "Intense", "Crazy"]:

         alert.append({
            "alert_type": "system",
            "severity": "high",
            "message": f"High resource usage detected (CPU: {event.get('cpu_load_level')}, Memory: {event.get('memory_load_level')})",
            "raw_payload":{}, 
            "status": "open",
            "risk_level": "high",
            "summary": "HIGH_RESOURCE_USAGE",
            "performance":{}, 
            "network":{}, 
            "security":{}, 
            "created_at": datetime.now(timezone.utc).isoformat()
        })

    return event
