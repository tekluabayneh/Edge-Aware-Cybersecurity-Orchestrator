def set_network_rule(event): 
    payload = event.get("payload",{})
    ConnectionMonitoring = payload.get("ConnectionMonitoring",{})
    ActiveSockets =  ConnectionMonitoring.get("ActiveSockets", {})
    NetworkInterfaces =  ConnectionMonitoring.get("NetworkInterfaces", {})   
    ConnectionPatterns =  ConnectionMonitoring.get("ConnectionPatterns", {})   

    # set rule for activesockets
    activesockets = ActiveSockets[0]
    for actsocket in activesockets:
         socket_findings_rules = []
         connection_type = actsocket.get("connection_type")
         is_localhost = actsocket.get("is_localhost")
         suspicious = actsocket.get("suspicious")
         risk_level = actsocket.get("risk_level")

         if connection_type == "listening" and suspicious == True: 
             socket_findings_rules.append({"socket_listening":"A service is listening in a way that may expose the system."})
         elif suspicious == True  and risk_level in ["High", "High_risk"]:
             socket_findings_rules.append({"socket_listening":"High risk socket activity detected"})
         elif is_localhost == True and suspicious == False: 
             socket_findings_rules.append({"socket_listening":"Local listening socket (normal)"}) 
         else: 
             socket_findings_rules.append({"socket_listening": "No suspicious socket behavior detected" })

         actsocket["active_socker_rules"] = socket_findings_rules



      

         networkINterfaces = NetworkInterfaces[0]
         for  netINterface in networkINterfaces:
             networkInterface_rules = []
             is_up_and_running = netINterface.get("is_up_and_running")
             interface_type =netINterface.get("interface_type")
             internal_only = netINterface.get("internal_only")
             ip_count = netINterface.get("ip_count")
             has_public_ip = netINterface.get("has_public_ip")


             if interface_type == "loopback" or internal_only and not has_public_ip:
                 networkInterface_rules.append({
                      "interface_status": "normal",
                      "message": "Internal interface running normally",
                      "severity": "info"
                }) 

             elif interface_type != "loopback" and  has_public_ip and not internal_only:
                networkInterface_rules.append({
                  "interface_status": "external_exposure",
                  "message": "Network interface exposed to external network",
                  "severity": "warning"
                    })
             elif has_public_ip  and ip_count > 4000 and is_up_and_running: 
                networkInterface_rules.append({
                      "interface_status": "suspicious",
                      "message": "Suspicious network interface configuration detected",
                      "severity": "danger"
                        })
             else: 
                    networkInterface_rules.append({
                  "interface_status": "normal",
                  "severity": "info",
                  "message": "Internal interface running normally"
                  })

             netINterface["networkInterface_rules"] = networkInterface_rules


             connectionpattern =  ConnectionPatterns[0]
             connectionpatterns_rules = []
             for conn in connectionpattern:
                 pattern_type = conn.get("pattern_type")
                 is_suspicious_volume = conn.get("is_suspicious_volume")
                 traffic_category = conn.get("traffic_category")
                 potential_scan = conn.get("potential_scan")
                 remoteIp = conn.get("remoteIp", "Unknown IP")
                  
                 if pattern_type == "listening" and is_suspicious_volume != True: 
                     connectionpatterns_rules.append({
                    "connectionpattern_status": "",
                    "message": f"{remoteIp} is listening with normal volume."
                            })
                 else: 
                     connectionpatterns_rules.append({
                     "connectionpattern_status": "normal",
                     "message": f"{remoteIp} has suspicious pattern/volume."
                        })                  


                 if traffic_category == "None" and potential_scan != True: 
                    connectionpatterns_rules.append({
                            "connectionpattern_status": "",
                            "message": f"{remoteIp} traffic category normal and no scan detected."
                        })
                 else:
                        connectionpatterns_rules.append({
                            "connectionpattern_status": "normal",
                            "message": f"{remoteIp} traffic suspicious or potential scan detected."
                    })


                 conn["connectionpatterns_rules"] = connectionpatterns_rules






    
