from shared.schemas.schema import Event

def intgerity_normalizer(intgerityPaylod) -> Event: 
    return Event(
            payload=intgerityPaylod,
            type="integrity",
            category="integrity", 
            source="agent",
            tags=["integrity"]
            )
    
