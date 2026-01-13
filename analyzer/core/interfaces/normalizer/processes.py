from fastapi.encoders import jsonable_encoder
from shared.schemas.schema import Event


def processes_normalizer(processesPaylod: dict) -> Event: 
     payload =jsonable_encoder(processesPaylod) 
     return Event(
            payload=payload,
            type="processes",
            category="processes", 
            source="agent",
            tags=["processes"]
            )
  
