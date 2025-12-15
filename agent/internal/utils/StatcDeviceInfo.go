package utils

import (
	"github.com/google/uuid"
	"github.com/shirou/gopsutil/v3/host"
)

type StaticSysInfoType struct {
	HostName     string `bson:hostname`
	Platform     string `bson:platform`
	MachineID    string `json:"machine_id"`
	AgentVersion string `json:"agent_version"`
	Status       string `json:"online"`
	OS           string `bson:platform`
}

func StaticSysInfo() StaticSysInfoType {

	info, err := host.Info()

	if err != nil {
		panic(err)
	}

	// 3️⃣ Generate UUID
	uid := uuid.NewString()
	machien_id := "machine_id_" + info.HostID + "_" + uid

	return StaticSysInfoType{
		HostName:     info.Hostname,
		Platform:     info.Platform,
		MachineID:    machien_id,
		AgentVersion: info.PlatformVersion,
		Status:       "online",
		OS:           info.OS,
	}
}
