from shared.schemas.schema import Event
from fastapi.encoders import jsonable_encoder

def netwrok_normalizer(networkPaylod):
      netwrok_normalizer_result ={
              "payload":networkPaylod,
            "type":"integrity",
            "category":"integrity", 
            "source":"agent",
            "tags":["ConnectionMonitoring,	AbuseIPDBResponse"]
            } 
      return netwrok_normalizer_result
 
