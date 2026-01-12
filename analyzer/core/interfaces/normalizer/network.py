
from shared.schemas.schema import Event


def netwrok_normalizer(networkPaylod) -> Event: 
     return Event(
            payload=networkPaylod,
            type="integrity",
            category="integrity", 
            source="agent",
            tags=["ConnectionMonitoring,	AbuseIPDBResponse"]
            )
 
