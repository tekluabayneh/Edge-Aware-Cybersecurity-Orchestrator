from fastapi.encoders import jsonable_encoder
from shared.schemas.schema import Event


def system_normalizer(systemPaylod: dict) -> Event: 
     payload =jsonable_encoder(systemPaylod) 
     return Event(
            payload=payload,
            type="system",
            category="system", 
            source="agent",
            tags=["system"]
            )
  
