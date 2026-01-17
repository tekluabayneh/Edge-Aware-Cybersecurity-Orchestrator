import json
import math

def proccess_enrich(event): 
    # print("Clean payload going to pipeline:\n" + json.dumps(event, indent=2))
    for payload in event.get("payload", []):
        if "CPUPercent" not in payload:
            continue
        cpu = math.floor(payload["CPUPercent"])
        if cpu <= 10:
            level = "Chill"
        elif cpu <= 25:
            level = "Light"
        elif cpu <= 40:
            level = "Normal"
        elif cpu <= 55:
            level = "Medium"
        elif cpu <= 70:
            level = "Heavy"
        elif cpu <= 90:
            level = "Intense"
        else:
            level = "Crazy"
        payload["cpu_load_level"] = level            



        for key, value in payload.items():
            if key == "Memory": 
              for key, value in value.items():
                    ram = value 
                    if ram <= 10:
                        level = "Chill"
                    elif ram <= 25:
                        level = "Light"
                    elif ram <= 40:
                        level = "Normal"
                    elif ram <= 55:
                        level = "Medium"
                    elif ram <= 70:
                        level = "Heavy"
                    elif ram <= 90:
                        level = "Intense"
                    else:
                        level = "Crazy"
                    payload["load_load_level"] = level            



# is_system_process → True if under system folders
#
# is_suspicious → based on name, path, or hash
#
# risk_score → numeric score for potential threat
#
# hash → md5/sha256 hash of executable
#
# parent_process → PID or name of parent
#
# is_running → True/False if currently active
#
# privilege_level → admin, user, system
#
# tags → e.g., "malware_suspected", "trusted"
#
# last_seen → timestamp of last execution
#
# resource_usage → CPU/memory if available
