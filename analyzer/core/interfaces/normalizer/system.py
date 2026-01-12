
from shared.schemas.schema import Event


def system_normalizer(systemPaylod: dict) -> Event: 
     return Event(
            payload=systemPaylod,
            type="system",
            category="system", 
            source="agent",
            tags=["system"]
            )
  
