package network

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
)

// Active sockets	Local/remote IP, port, PID, state	Detect suspicious processes and endpoints
// Network interfaces	Up/down, IP addresses	Detect connectivity and unusual interfaces
// Connection patterns	Frequency, volume	Detect abnormal behavior

// i need to gather all device info like:
//    for sockets: Local/remote IP, port, PID, state
//    for interfaces: Up/down, IP addresses
//    for Connection patterns: Frequency, volume	Detect abnormal behavior

type ConnectionMonitoringType struct {
	ActiveSockets []struct {
		LocalIP string `json:"localIp"`
		Port    uint32 `json:"port"`
		PID     int32  `json:"pid"`
		State   string `json:"state"`
		Process string `json:"process,omitempty"`
	} `json:"activeSockets"`

	NetworkInterfaces []struct {
		Name        string   `json:"name"`
		Up          bool     `json:"up"`
		Down        bool     `json:"down"`
		IPAddresses []string `json:"ipAddresses"`
	} `json:"networkInterfaces"`

	ConnectionPatterns []struct {
		RemoteIP  string `json:"remoteIp"`
		Frequency int    `json:"frequency"`
		Volume    int64  `json:"volume"`
	} `json:"connectionPatterns"`
}

func ConnectionsMonitoring() {
	sockets, err := net.ConnectionStat("")
	if err != nil {
		panic(err)
	}
	fmt.Println("test connectionsMonitoring work")
}
