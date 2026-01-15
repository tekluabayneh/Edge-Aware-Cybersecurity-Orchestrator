from core.interfaces.normalizer import integrity, network, processes, security, system


def Input(payload):
     try: 
      result = {
        "network":   network.netwrok_normalizer(payload.network),
        "processes": processes.processes_normalizer(payload.processes),
        "system":    system.system_normalizer(payload.system),
        "integrity": integrity.intgerity_normalizer(payload.integrity),
        "security":  security.security_normalizer(payload.security),
        } 
    
      return result
     except AttributeError as e: 
        print("ERROR in Input():", str(e))
      
