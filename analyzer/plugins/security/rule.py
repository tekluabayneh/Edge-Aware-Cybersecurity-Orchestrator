from datetime import datetime, timezone
def set_security_rule(event, alert): 
    securityPaylod  = event.get("payload")
    malicious_processes = securityPaylod.get("malicious_processes", {}) 
    suspicious_files =  securityPaylod.get("suspicious_files", {}) 

    suspicious_file_rules = []
    malicious_processe_rules = []
    security_rules= []
    for suspicious_file in suspicious_files:
        # Rule 1: file name risk
        file_risk = suspicious_file.get("fileNameFisk", "unknown")
        if file_risk == "high_risk":
             suspicious_file_rules.append({
                "rule_name": "File Risk",
                "status": "critical",
                "message": "File name is high risk!"
            })
        else:
            suspicious_file_rules.append({
                "rule_name": "File Risk",
                "status": "normal",
                "message": "File name looks fine."
            })

        # Rule 2: executable file
        install_flag = suspicious_file.get("install.sh-content_Is", "")
        if "executable" in install_flag.lower():
            suspicious_file_rules.append({
                "rule_name": "Executable File",
                "status": "warning",
                "message": "This file is executable."
            })
        else:
            suspicious_file_rules.append({
                "rule_name": "Executable File",
                "status": "normal",
                "message": "File is not executable."
            })

        # Rule 3: file size risk
        if suspicious_file.get("is_file_size_risky", False):
            suspicious_file_rules.append({
                "rule_name": "File Size",
                "status": "warning",
                "message": f"File size is risky: {securityPaylod.get('file_size_MB')}"
            })
        else:
            suspicious_file_rules.append({
                "rule_name": "File Size",
                "status": "normal",
                "message": f"File size is {securityPaylod.get('file_size_MB')}, looks fine."
            })

        # Rule 4: process username
        username = suspicious_file.get("proccess_username", "")
        if "suspicious" in username.lower():
            suspicious_file_rules.append({
                "rule_name": "Process Username",
                "status": "warning",
                "message": "Process is running with a suspicious username!"
            })
        else:
            suspicious_file_rules.append({
                "rule_name": "Process Username",
                "status": "normal",
                "message": "Process username looks fine."
            })

        # Rule 5: process path
        path = suspicious_file.get("proccess_path", "")
        if "normal path" in path.lower():
            suspicious_file_rules.append({
                "rule_name": "Process Path",
                "status": "normal",
                "message": "Process path looks fine."
            })
        else:
            suspicious_file_rules.append({
                "rule_name": "Process Path",
                "status": "warning",
                "message": "Process is running from an unusual path!"
            }) 

        security_rules.append(suspicious_file_rules)


        for malicious_processe in malicious_processes:
            proc_name_status = malicious_processe.get("proccess_name", "")
            if proc_name_status == "name contains looks good":
                malicious_processe_rules.append("Process name format looks normal.")
            elif proc_name_status == "name contains nonsense and looks suspicious":
                malicious_processe_rules.append("Process name contains unusual or suspicious characters.")
            else:
                malicious_processe_rules.append("Process name could not be determined.")
        
            proc_path_status = malicious_processe.get("proccess_path", "")
            if proc_path_status == "process running from normal path":
                malicious_processe_rules.append("Process is running from a standard system or user directory.")
            elif proc_path_status == "process running from suspicious path":
                malicious_processe_rules.append("Process is running from a temporary or memory-backed path, which may be risky.")
            else:
                malicious_processe_rules.append("Process path could not be determined.")

            proc_user_status = malicious_processe.get("proccess_username", "")
            if proc_user_status == "username is ok":
                malicious_processe_rules.append("Process is running under a normal user account.")
            elif proc_user_status == "suspicious username":
                malicious_processe_rules.append("Process is running under an unexpected or suspicious user.")
            else:
                malicious_processe_rules.append("Process username could not be determined.")

    all_statuses = [] 

    for ruls_group in security_rules:
         for rule in ruls_group: 
             all_statuses.append(rule["status"])
    

    if "critical" in all_statuses:
        event["overall_security"] = "Critical"
    elif "warning" in all_statuses:
        event["overall_security"] = "Medium"
    else:
        event["overall_security"] = "Chill"


    if event.get("overall_security") in ["Medium", "Critical"]:
        alert.append({
            "alert_type": "security",
            "severity": event.get("overall_security"),
            "message": f"Security state is {event.get('overall_security')}",
            "raw_payload":{}, 
            "status": "open",
            "risk_level": "high",
            "summary": "OVERALL_SECURITY_RISK",
            "performance":{}, 
            "network":{} ,
            "security":{},
            "created_at": datetime.now(timezone.utc).isoformat()
        })


    return event
