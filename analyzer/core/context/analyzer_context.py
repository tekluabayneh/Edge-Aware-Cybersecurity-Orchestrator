from typing import Any, Dict, List, Optional
from  api.schemas.Telementory import RawTelemetryPayload
from datetime import datetime, timezone
import logging
import time
from shared.schemas.schema import Alert
class AnalyzerContext:
    def __init__(self):
        self.request_id: str = None          
        self.user: Dict[str, Any] 
        self.config: dict = {}                 
        self.input_data: Any = None             
        self.normalized_data: dict = {}          
        self.enrichments: dict = {}               
        self.alerts: list[Alert] = []              
        self.rules_matched: list[Any] = []        
        self.start_time = time.time()
        self.logger = logging.getLogger(__name__)
        self.start_time = datetime.now()
        self.start_time = datetime.now(timezone.utc)
        self.end_time = datetime.now()
        self.raw_payload = RawTelemetryPayload

    def add_enrichment(self, category: str, key: str, value: Any):
        self.enrichments.setdefault(category, {})[key] = value

    def has_alerts(self) -> bool:
        return len(self.alerts) > 0

    # Sometimes also context managers
    def __enter__(self):
        # setup tracing / timing
        return self

    def __exit__(self, *args):
        # log duration, cleanup
        pass
