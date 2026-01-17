from fastapi.encoders import jsonable_encoder
from shared.schemas.schema import Event

def security_normalizer(securityPaylod): 
     security_normalizer_result = {
             "payload":securityPaylod,
             "type":"security",
             "category":"security", 
            "source":"agent",
            "tags":[ "Firewall","Antivirus","MaliciousProcesses", "SuspiciousFiles"]
            }  
     return security_normalizer_result 
  
