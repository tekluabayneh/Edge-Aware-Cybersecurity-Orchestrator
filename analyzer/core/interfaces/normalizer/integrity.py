from shared.schemas.schema import Event
from fastapi.encoders import jsonable_encoder

def intgerity_normalizer(intgerityPaylod): 
       integrity_normalizer_result = { 
            "payload":intgerityPaylod,
            "type":"integrity",
            "category":"integrity", 
            "source":"agent",
            "tags":["integrity"]
            }
       return integrity_normalizer_result

    
