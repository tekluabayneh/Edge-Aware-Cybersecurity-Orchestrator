import math

def proccess_enrich(event): 
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

        if "Memory" in payload and "rss" in payload["Memory"]:
            rss_bytes = payload["Memory"]["rss"]
            rss_mb = rss_bytes / (1024 * 1024)  
            if rss_mb <= 50:
                level = "Chill"
            elif rss_mb <= 200:
                level = "Light"
            elif rss_mb <= 500:
                level = "Normal"
            elif rss_mb <= 1000:
                level = "Medium"
            elif rss_mb <= 2000:
                level = "Heavy"
            elif rss_mb <= 4000:
                level = "Intense"
            else:
                level = "Crazy"
            payload["memory_load_level"] = level  

 
            return payload
