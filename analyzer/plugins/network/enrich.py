import json

from shared.schemas.schema import Event


def network_enrich(event):
        payload = event.get("payload",{})
        ConnectionMonitoring = payload.get("ConnectionMonitoring",{})
        ActiveSockets =  ConnectionMonitoring.get("ActiveSockets", {})   
        NetworkInterfaces =  ConnectionMonitoring.get("NetworkInterfaces", {})   
        ConnectionPatterns =  ConnectionMonitoring.get("ConnectionPatterns", {})   

        suspicious_ports = {21, 23, 25, 445, 1433, 3389, 5900, 4444, 6667, 1337, 31337, 12345, 8080}  
        for socket_item in ActiveSockets[0]:  
            local_ip   = socket_item.get("localaddr", {}).get("ip", "")
            local_port = socket_item.get("localaddr", {}).get("port", 0)
            remote_ip  = socket_item.get("remoteaddr", {}).get("ip", "")
            remote_port = socket_item.get("remoteaddr", {}).get("port", 0)
            pid        = socket_item.get("pid", 0)
        
            # is_localhost
            is_loopback = local_ip.startswith("127.") or local_ip == "127.0.0.1" or local_ip == "::1"
            socket_item["is_localhost"] = is_loopback and (remote_ip == "0.0.0.0" or remote_ip.startswith("127.") or remote_ip == "")
        
            #  connection_type
            if remote_ip == "0.0.0.0" and remote_port == 0:  
                socket_item["connection_type"] = "listening"
            elif is_loopback:
                socket_item["connection_type"] = "loopback"
            elif local_ip in ["0.0.0.0", "::"] or not (local_ip.startswith(("192.168.", "10.", "172.")) or local_ip.startswith("127.")):
                socket_item["connection_type"] = "external"  
            else:
                socket_item["connection_type"] = "internal"   
        
            #  suspicious + risk
            exposed = not is_loopback and local_ip != "0.0.0.0"   
        
            is_susp_port = local_port in suspicious_ports or remote_port in suspicious_ports
        
            suspicious = False
            risk = "Low"
        
            if socket_item["connection_type"] == "listening" and exposed and is_susp_port:
                suspicious = True
                risk = "High"
            elif socket_item["connection_type"] == "external" and is_susp_port:
                suspicious = True
                risk = "High"
            elif pid == 0 and local_port > 1024:   
                suspicious = True
                risk = "Medium"

            socket_item["suspicious"] = suspicious
            socket_item["risk_level"] = risk
    
   
   # NetworkInterfaces 
        for  interface in  NetworkInterfaces[0]: 
                def is_private_ip(ip: str) -> bool:
                    ip = ip.split('/')[0].strip()  

                    if ip in ("::1", "127.0.0.1") or ip == "":
                        return True
                    if ':' in ip:  
                        return ip.startswith(("fc", "fd", "fe80:", "::1"))
                    if '.' not in ip:
                        return False
                    parts = ip.split('.')
                    if len(parts) != 4:
                        return False

                    try:
                        a, b, _, _ = map(int, parts)
                        return (a == 10 or (a == 172 and 16 <= b <= 31) or (a == 192 and b == 168) or a == 127)
                    except ValueError:
                        return False


                ip_addresses   = interface.get("ipAddresses", {})
            # Normalize to list of strings
                if isinstance(ip_addresses, dict):
                    ip_list = list(ip_addresses.values())
                elif isinstance(ip_addresses, list):
                    ip_list = [entry.get("addr", "") for entry in ip_addresses if isinstance(entry, dict)]
                else:
                    ip_list = []

                # Flags
                has_public_ip = False
                only_loopback = True
                ip_count = len(ip_list)

                for raw_ip in ip_list:
                    if not raw_ip:
                        continue
                    if not is_private_ip(raw_ip):
                        has_public_ip = True
                        only_loopback = False
                    elif not (raw_ip.startswith("127.") or raw_ip == "::1"):
                        only_loopback = False

                # Interface type (simple but better than before)
                name = interface.get("name", "").lower()
                if "lo" in name or only_loopback:
                    interface_type = "loopback"
                elif "wl" in name or "wifi" in name:
                    interface_type = "wifi"
                elif "en" in name or "eth" in name:
                    interface_type = "ethernet"
                else:
                    interface_type = "other"

                # Up / running
                is_up_and_running = interface.get("up") == "up"

                # Internal only = no public IP **and** not loopback-only
                internal_only = (not has_public_ip) and (not only_loopback)

                # Store results
                interface["interface_type"]   = interface_type
                interface["is_up_and_running"] = is_up_and_running
                interface["internal_only"]     = internal_only
                interface["has_public_ip"]     = has_public_ip
                interface["ip_count"]          = ip_count

               

           # ConnectionPatterns 
        for conn in ConnectionPatterns[0]:  
            remote_ip = conn.get("remoteIp", "unknown")
            frequency = conn.get("frequency", 0)
            volume    = conn.get("volume", 0)  
        
            # ── pattern_type 
            if remote_ip in ("0.0.0.0", "::"):
                conn["pattern_type"] = "listening"
            elif frequency == 0 and volume < 1000: 
                conn["pattern_type"] = "idle"
            else:
                conn["pattern_type"] = "active"      
        
            # ── is_suspicious_volume 
            conn["is_suspicious_volume"] = frequency > 40
        
            # ── traffic_category 
            if frequency == 0 and volume < 1000:
                conn["traffic_category"] = "None"
            elif frequency <= 4:
                conn["traffic_category"] = "Low"
            elif frequency <= 40:
                conn["traffic_category"] = "Medium"
            else:
                conn["traffic_category"] = "High"
        
            bytes_per_connection = volume / max(frequency, 1) 
            is_scan_like = ( frequency > 25 and bytes_per_connection < 2000 )
            conn["potential_scan"] = is_scan_like


        return event


