

from shared.schemas.schema import Event


def processes_normalizer(processesPaylod: dict) -> Event: 

     return Event(
            payload=processesPaylod,
            type="processes",
            category="processes", 
            source="agent",
            tags=["processes"]
            )
  
