from fastapi.encoders import jsonable_encoder
from shared.schemas.schema import Event

def security_normalizer(securityPaylod: dict) -> Event: 
     payload =jsonable_encoder(securityPaylod) 
     return Event(
            payload=payload,
            type="security",
            category="security", 
            source="agent",
            tags=[ "Firewall","Antivirus","MaliciousProcesses", "SuspiciousFiles"]

                  )
  
