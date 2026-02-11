from dataclasses import dataclass, field
from datetime import datetime
from typing import Any, Dict, List, Optional


class Event:
    def __init__(
        self,
        category: str,
        type: str,
        source: str,
        payload: dict,
        severity: str = "info",
        tags: list[str] = None,
    ):
        self.category = category
        self.type = type
        self.source = source
        self.payload = payload
        self.severity = severity
        self.tags = tags or []


@dataclass
class Alert:
    """
    Represents a detected issue, anomaly, or security finding during analysis.
    Created by rules or enrich plugins and collected in AnalyzerContext.alerts.
    """
    # Required / core fields
    category: str          
    severity: str         
    message: str         

    rule_id: Optional[str] = None   
    rule_name: Optional[str] = None
    timestamp: datetime = field(default_factory=datetime.utcnow)

    evidence: Dict[str, Any] = field(default_factory=dict)   
    enriched_data: Dict[str, Any] = field(default_factory=dict)

    alert_id: Optional[str] = None   
    related_alerts: List['Alert'] = field(default_factory=list)  

    def __post_init__(self):
        """Auto-fixes / defaults after init"""
        if not self.message:
            self.message = f"Alert in category '{self.category}' (no message provided)"

        # Optional: normalize severity to lowercase
        if self.severity:
            self.severity = self.severity.lower()

    def add_evidence(self, key: str, value: Any) -> None:
        """Add supporting data that triggered the alert"""
        self.evidence[key] = value

    def add_enrichment(self, key: str, value: Any) -> None:
        """Add context from enrichment plugins (geoip, whois, reputation, etc.)"""
        self.enriched_data[key] = value

    def is_high_severity(self) -> bool:
        return self.severity in ("high", "critical")

    def to_dict(self) -> Dict[str, Any]:
        """Useful for JSON serialization / API response / storage"""
        data = {
            "category": self.category,
            "severity": self.severity,
            "message": self.message,
            "timestamp": self.timestamp.isoformat(),
        }

        if self.rule_id:
            data["rule_id"] = self.rule_id
        if self.rule_name:
            data["rule_name"] = self.rule_name
        if self.source:
            data["source"] = self.source
        if self.destination:
            data["destination"] = self.destination
        if self.ioc:
            data["ioc"] = self.ioc
        if self.evidence:
            data["evidence"] = self.evidence
        if self.enriched_data:
            data["enriched_data"] = self.enriched_data

        return data

    def __str__(self) -> str:
        return (
            f"[{self.severity.upper()}] {self.category} - "
            f"{self.message} "
            f"(rule: {self.rule_name or self.rule_id or 'n/a'}) "
            f"@ {self.timestamp.isoformat(timespec='seconds')}"
        )


class User: 
    def __init__(self, email:str, agent_id: str, agent_token: str):
        self.email=  email
        self.agent_id= agent_id
        self.agent_token= agent_token

