from core.interfaces.normalizer import integrity, network, processes, security, system


def Input(payload):
     try: 
      result = {
        "network":   network.netwrok_normalizer(payload.get("network")),
        "processes": processes.processes_normalizer(payload.get("processes")),
        "system":    system.system_normalizer(payload.get("system")),
        "integrity": integrity.intgerity_normalizer(payload.get("integrity")),
        "security":  security.security_normalizer(payload.get("security")),
        } 
     
      return result
     except AttributeError as e: 
        print("ERROR in Input():", str(e))
     

