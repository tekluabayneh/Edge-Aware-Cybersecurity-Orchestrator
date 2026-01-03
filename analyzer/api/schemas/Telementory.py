from pydantic import BaseModel
from typing import List, Optional
from datetime import datetime

# -------------------- Process --------------------
class ProcessInfo(BaseModel):
    pid: int
    name: str
    cpu: float
    memory: int
    parent_pid: Optional[int] = None

# -------------------- Networking --------------------
class Address(BaseModel):
    IP: str
    Port: int

class ActiveSocket(BaseModel):
    Fd: int
    Family: int
    Type: int
    LocalAddr: Address
    RemoteAddr: Address
    Status: str
    Uids: List[int]
    Pid: int

class NetworkInterface(BaseModel):
    Name: str
    Up: bool
    Down: bool
    IPAddresses: List[str]

class ConnectionPattern(BaseModel):
    RemoteIP: str
    Frequency: int
    Volume: int

class AbuseIPDBData(BaseModel):
    IPAddress: str
    AbuseConfidenceScore: int
    TotalReports: int
    IsWhitelisted: bool

class AbuseIPDBResponse(BaseModel):
    data: AbuseIPDBData

class ConnectionMonitoringType(BaseModel):
    ActiveSockets: List[ActiveSocket]
    NetworkInterfaces: List[NetworkInterface]
    ConnectionPatterns: List[ConnectionPattern]

class NetworkInfo(BaseModel):
    ConnectionMonitoring: ConnectionMonitoringType
    AbuseIPDBResponse: AbuseIPDBResponse

# -------------------- System --------------------
class SystemInfo(BaseModel):
    Uptime: str
    Cpu: List[float]
    Ram: float
    Disk: float
    Network: int

# -------------------- Security --------------------
class AntivirusStatus(BaseModel):
    Running: bool
    Name: str
    Detected: Optional[str] = None

class SuspiciousProcess(BaseModel):
    PID: int
    Name: str
    CPUPercent: float
    Memory: Optional[int] = None

class SuspiciousFiletype(BaseModel):
    Path: str
    Extension: str
    Name: str
    Size: int
    Mode: Optional[str] = None
    Content: str

class SecurityInfo(BaseModel):
    Firewall: bool
    Antivirus: AntivirusStatus
    MaliciousProcesses: List[SuspiciousProcess]
    SuspiciousFiles: List[SuspiciousFiletype]

# -------------------- Integrity --------------------


class criticalFiles(): 
    criticalFiles :str

class IntegrityInfo(BaseModel):
    OS :           str
    KernelVersion :str
    PatchLevel    :str
    CriticalFiles : criticalFiles
    CollectedAt   :int

# -------------------- Top-level Payload --------------------
class TelemetryPayload(BaseModel):
    agent_id: str
    agent_token: str
    SystemInfo: SystemInfo
    Security: SecurityInfo
    Network: NetworkInfo
    Processes: List[ProcessInfo]
    Integrity: IntegrityInfo


