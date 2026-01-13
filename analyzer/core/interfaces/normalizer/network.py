from shared.schemas.schema import Event
from fastapi.encoders import jsonable_encoder

def netwrok_normalizer(networkPaylod) -> Event: 
     payload = jsonable_encoder(networkPaylod) 
     return Event(
            payload=payload,
            type="integrity",
            category="integrity", 
            source="agent",
            tags=["ConnectionMonitoring,	AbuseIPDBResponse"]
            )
 
