package network

import (
	"encoding/json"
	"log"

	"github.com/shirou/gopsutil/v3/net"
)

// Active sockets	Local/remote IP, port, PID, state	Detect suspicious processes and endpoints
// Network interfaces	Up/down, IP addresses	Detect connectivity and unusual interfaces
// Connection patterns	Frequency, volume	Detect abnormal behavior
// i need to gather all device info like:
//	for sockets: Local/remote IP, port, PID, state
//	for interfaces: Up/down, IP addresses
//	for Connection patterns: Frequency, volume	Detect abnormal behavior

type ActiveSocket []struct {
	Fd        uint32 `json:"fd"`
	Family    uint32 `json:"family"`
	Type      uint32 `json:"type"`
	LocalAddr struct {
		IP   string `json:"ip"`
		Port uint32 `json:"port"`
	} `json:"localaddr"`
	RemoteAddr struct {
		IP   string `json:"ip"`
		Port uint32 `json:"port"`
	} `json:"remoteaddr"`
	Status string  `json:"status"`
	Uids   []int32 `json:"uids"`
	Pid    int32   `json:"pid"`
}

type networkInterfaces []struct {
	Name        string              `json:"name"`
	Up          string              `json:"up"`
	Down        string              `json:"down"`
	IPAddresses []net.InterfaceAddr `json:"ipAddresses"`
}

type ConnectionPatterns []struct {
	RemoteIP  string `json:"remoteIp"`
	Frequency int    `json:"frequency"`
	Volume    uint32 `json:"volume"`
}

type ConnectionMonitoringType struct {
	ActiveSockets      []ActiveSocket
	NetworkInterfaces  []networkInterfaces
	ConnectionPatterns []ConnectionPatterns
}

func ConnectionsMonitoring() ConnectionMonitoringType {
	var ConnectionMonitoringData ConnectionMonitoringType
	var ActiveSocketsData ActiveSocket
	var networkInterfacesData networkInterfaces
	var ConnectionPatternsData ConnectionPatterns

	sockets, err := net.Connections("inet")

	if err != nil {
		log.Panicln("error occur while retriving sockets")
	}

	netwroks, err := net.Interfaces()
	if err != nil {
		log.Panicln("error occur while retriving netwroks")
	}

	ips := []net.InterfaceAddr{}
	addr := netwroks[0].Addrs
	for _, ip := range addr {
		ips = append(ips, ip)
	}

	singleInterface := struct {
		Name        string              `json:"name"`
		Up          string              `json:"up"`
		Down        string              `json:"down"`
		IPAddresses []net.InterfaceAddr `json:"ipAddresses"`
	}{
		Name:        netwroks[0].Name,
		Up:          netwroks[0].Flags[0],
		Down:        netwroks[0].Flags[0],
		IPAddresses: ips,
	}

	networkInterfacesData = append(networkInterfacesData, singleInterface)

	/// socket info
	b, _ := json.Marshal(sockets)
	err = json.Unmarshal(b, &ActiveSocketsData)
	if err != nil {
		log.Panicln("error occur while Unmarshal bytes")
	}

	/// collect connection patterns
	for _, v := range ActiveSocketsData {
		singleConnectionPatter := struct {
			RemoteIP  string `json:"remoteIp"`
			Frequency int    `json:"frequency"`
			Volume    uint32 `json:"volume"`
		}{
			RemoteIP:  v.RemoteAddr.IP,
			Frequency: int(v.RemoteAddr.Port),
			Volume:    v.LocalAddr.Port,
		}
		ConnectionPatternsData = append(ConnectionPatternsData, singleConnectionPatter)
	}

	ConnectionMonitoringData.ActiveSockets = append(ConnectionMonitoringData.ActiveSockets, ActiveSocketsData)
	ConnectionMonitoringData.ConnectionPatterns = append(ConnectionMonitoringData.ConnectionPatterns, ConnectionPatternsData)
	ConnectionMonitoringData.NetworkInterfaces = append(ConnectionMonitoringData.NetworkInterfaces, networkInterfacesData)

	return ConnectionMonitoringData
}
