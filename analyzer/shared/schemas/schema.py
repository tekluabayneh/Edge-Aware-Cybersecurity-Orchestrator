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
