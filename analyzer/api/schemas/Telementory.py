from pydantic import BaseModel, Field, ConfigDict
from typing import List, Optional, Dict, Any

class Memory(BaseModel):
    rss: int
    vms: int
    hwm: int = 0
    data: int = 0
    stack: int = 0
    locked: int = 0
    swap: int = 0

class Addr(BaseModel):
    ip: str
    port: int

class ActiveSocket(BaseModel):
    model_config = ConfigDict(populate_by_name=True)  # Allows fd → Fd
    fd: int = Field(alias='Fd')
    family: int = Field(alias='Family')
    type: int = Field(alias='Type')
    localaddr: Addr = Field(alias='LocalAddr')
    remoteaddr: Addr = Field(alias='RemoteAddr')
    status: str = Field(alias='Status')
    uids: List[int] = Field(alias='Uids')
    pid: int = Field(alias='Pid')

class IPAddrItem(BaseModel):
    addr: str

class NetworkInterface(BaseModel):
    model_config = ConfigDict(populate_by_name=True)
    name: str = Field(alias='Name')
    up: str = Field(alias='Up')
    down: str = Field(alias='Down')
    ipAddresses: List[IPAddrItem] = Field(alias='IPAddresses')

class ConnectionPattern(BaseModel):
    model_config = ConfigDict(populate_by_name=True)
    remoteIp: str = Field(alias='RemoteIP')
    frequency: int = Field(alias='Frequency')
    volume: int = Field(alias='Volume')

class ConnectionMonitoring(BaseModel):
    ActiveSockets: List[List[ActiveSocket]]
    NetworkInterfaces: List[List[NetworkInterface]]
    ConnectionPatterns: List[List[ConnectionPattern]]

class Network(BaseModel):
    ConnectionMonitoring: ConnectionMonitoring
    AbuseIPDBResponse: Optional[Any] = None

class SuspProcess(BaseModel):
    PID: int
    Name: str
    CPUPercent: float
    Memory: Memory

class CriticalFile(BaseModel):
    path: str
    extension: str
    name: str
    size: int
    content: str

class Integrity(BaseModel):
    model_config = ConfigDict(populate_by_name=True)
    _os: str
    kernel_version: str
    patch_level: str
    critical_files: Dict[str, str]
    collected_at: int

# Top-level (add missing as Optional if needed)
class RawTelemetryPayload(BaseModel):
    email:str
    agent_id:str
    agent_token:str
    network: Network
    processes: List[SuspProcess]
    integrity: Integrity
    agent_id: Optional[str] = None