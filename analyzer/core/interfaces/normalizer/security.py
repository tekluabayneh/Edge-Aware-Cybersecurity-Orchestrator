

from shared.schemas.schema import Event


def security_normalizer(securityPaylod: dict) -> Event: 
    
     return Event(
            payload=securityPaylod,
            type="security",
            category="security", 
            source="agent",
            tags=[ "Firewall","Antivirus","MaliciousProcesses", "SuspiciousFiles"]

                  )
  
