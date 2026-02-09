import json
from  tools.RuleData import ALL_CONTENT_KEYWORDS  
def security_enrich(event): 
    payload = event.get("payload", {})
    malicious_processes = payload.get("malicious_processes", {}) 
    suspicious_files =  payload.get("suspicious_files", {}) 

    FILE_RISK_CATEGORIES = {
    "high_risk": [
        "exe", "msi", "bat", "cmd", "ps1", "vbs", "js", "scr", "com", "dll",
        "sh", "run", "bin", "elf",
        "app", "pkg", "dmg", "command",
        "py", "pl", "rb", "jar", "php"
    ],
    "medium_risk": [
        "doc", "docm", "xls", "xlsm", "ppt", "pdf",
        "zip", "rar", "7z", "tar", "tar.gz", "tar.xz", "iso",
        "deb", "rpm", "AppImage"
    ],
    "low_risk": [
        "txt", "md", "csv", "json", "xml", "log", "ini", "cfg",
        "jpg", "png", "gif", "svg",
        "mp3", "wav", "mp4", "mkv"
    ]
   } 

    HIGH_RISK_KEYWORDS = [
              "monitor",
              "collector",
              "service",
              "daemon",
              "updater",
              "installer",
              "loader",
              "inject",
              "hook",
              "proxy",
              "miner",
              "watcher",
              "scanner",
              "driver"
            ];

    risk_map = {
    tuple(FILE_RISK_CATEGORIES["high_risk"]):   "high_risk",
    tuple(FILE_RISK_CATEGORIES["medium_risk"]): "medium_risk",
    tuple(FILE_RISK_CATEGORIES["low_risk"]):    "low_risk",
    }

    for file_dict in suspicious_files:
        if not isinstance(file_dict,dict): 
            continue

        ext = file_dict.get("extension", "").lower()
        found = False 
        for exts, risk_level in risk_map.items(): 
            if ext in  exts: 
                file_dict["fileNameFisk"] = risk_level
                found = True 
                break 

            if not found: 
                file_dict["fileNameFisk"] = "unknown_risk" 


        file_name = file_dict.get("name", "").lower()
        if file_name in HIGH_RISK_KEYWORDS: 
            file_dict[file_name] = "risky file name"

        file_content = file_dict.get("content", "").lower()
        if file_content.startswith("#!/bin/bash"): 
            file_dict[file_name + "-content_Is"] = "file is executable"   
        else: 
            file_dict[file_name + "-content_Is"] = "file is NOT executable"   


        destructive_keywords = ALL_CONTENT_KEYWORDS.get("destructive",[])
        persistence_keywords = ALL_CONTENT_KEYWORDS.get("persistence",[])
        remote_exec_keywords= ALL_CONTENT_KEYWORDS.get("remote-exec",[])
        obfuscation_keywords =ALL_CONTENT_KEYWORDS.get("obfuscation",[])
        privilege_keywords = ALL_CONTENT_KEYWORDS.get("privilege", [])
        surveillance_keywords = ALL_CONTENT_KEYWORDS.get("surveillance",[])
        network_keywords = ALL_CONTENT_KEYWORDS.get("network",[])
        resource_abuse_keywords = ALL_CONTENT_KEYWORDS.get("resource_abuse", [])

        warnings = []
        if any(kw in file_content  for kw in destructive_keywords) and any(kw in file_content for kw in persistence_keywords): 
                warnings.append("destructive + persistence")

        if any(kw in file_content  for kw in remote_exec_keywords) and any(kw in file_content for kw in obfuscation_keywords): 
            warnings.append("remote-exec + obfuscation")

        if any(kw in file_content  for kw in privilege_keywords) and any(kw in file_content for kw in surveillance_keywords): 
            warnings.append("privilege-escalation + surveillance")

        if any(kw in file_content  for kw in network_keywords) and any(kw in file_content for kw in resource_abuse_keywords): 
            warnings.append("network + resource-abuse")

        if warnings: 
            file_dict["content file contains"]  = warnings 
            if len(warnings) > 2: 
                file_dict["file_risk"]  = "high_risk" 
        else: 
            file_dict["file contain"] = ["clean (basic keyword check)"]


# amke oabject with their name so you can get them from therer and put enrich for them for therer ok 
        PERMISSION_RISK_EXPLANATIONS = {
        "4755": "SUID root – runs as root. Only safe on real system tools, everywhere else = backdoor risk.",
        "4750": "SUID + group writable – attacker in group overwrites → instant root.",
        "4777": "SUID + world writable – attacker replaces root binary → game over.",
        "0777": "777 – anyone edits & runs it. Classic malware dropper.",
        "0666": "666 – anyone overwrites. Perfect for messing with configs/scripts.",
        "2777": "SGID + world writable – group takeover + anyone edits.",
        "2775": "SGID dir – files inherit group. Fine in teams, dangerous if group is pwned.",
        "2755": "SGID binary – runs as group. Suspicious unless known daemon.",
        "1777": "/tmp-style sticky – mostly ok, but watch for abuse.",
        "0644": "Normal readable file – usually safe.",
        "308": "Normal executable – standard, low risk.",
    }
    
        def to_file_mode_string(mode_dec):
            return format(int(mode_dec) & 0o7777, "04o")
        #
        #
        file_Mode = file_dict.get("mode", "")
        for key, value in PERMISSION_RISK_EXPLANATIONS.items(): 
         if len(str(file_Mode)) < 5: 
             mode_num = to_file_mode_string(file_Mode)
             if key == mode_num: 
                file_dict[file_name] = value
         elif file_Mode == key: 
                file_dict[file_name] = value
         else: 
            file_dict[file_name] = "not identifyed" 


         file_size =  file_dict.get("size")
         try: 
             if not file_size: 
                 file_dict["is_file_size_risky"] = "number not found"
                 continue
             file_size = int(file_size)
         except(ValueError, TypeError): 
           file_size = 0


         file_size_in_MB = file_size / 10_000_000
         file_dict["file_size_MB"] = f"{round(file_size_in_MB)}MB"

         if file_size > 10_000_000: 
              file_dict["size_category"] = "large_suspicious"
              file_dict["is_file_size_risky"] = True
              file_dict["size_note"] = "> 10 MB – possible dropper/payload"

         elif file_size > 50_00_00: 
              file_dict["size_category"] = "medium"
              file_dict["is_file_size_risky"] = False   
              file_dict["size_note"] = "500 KB – 10 MB – common range"
         else: 
              file_dict["size_category"] = "small_normal"
              file_dict["is_file_size_risky"] = False
              file_dict["size_note"] = "≤ 500 KB – typical for scripts/configs"

             


         for proc in malicious_processes:
            for k, v in proc.items():
                if k == "name" and len(v) > 20 and any(x in v for x in ["_", "$", "#", "@"]):
                    file_dict["proccess_name"] = "name contains nonsense and looks suspicious"
                else: 
                    file_dict["proccess_name"] = "name contains looks good"
        
                if k == "path" and v in ["/dev/shm", "/var/tmp", "/tmp"]:
                    file_dict["proccess_path"] = "process running from suspicious path"
                else: 
                    file_dict["proccess_path"] = "process running from normal path"
        
                if k == "username" and v not in ["init/systemd", "/init"]:
                    file_dict["proccess_username"] = "suspicious username"
                else: 
                    file_dict["proccess_username"] = "username is ok"






    # print("Clean payload going to pipeline:\n" + json.dumps(event, indent=2))
    return event

