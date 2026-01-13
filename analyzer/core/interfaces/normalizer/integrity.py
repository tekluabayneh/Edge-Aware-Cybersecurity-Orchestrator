from shared.schemas.schema import Event
from fastapi.encoders import jsonable_encoder

def intgerity_normalizer(intgerityPaylod) -> Event: 
    payload =jsonable_encoder(intgerityPaylod) 
    return Event(
            payload=payload,
            type="integrity",
            category="integrity", 
            source="agent",
            tags=["integrity"]
            )

    
