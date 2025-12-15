package utils

import (
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type SysInfo struct {
	Hostname string                 `bson:hostname`
	Platform string                 `bson:platform`
	CPU      []float64              `bson:cpu`
	RAM      *mem.VirtualMemoryStat `bson:ram`
	Disk     *disk.UsageStat        `bson:disk`
	Network  int64                  `bson:network`
}

func GetUptime() string {
	uptime, err := host.Uptime()
	if err != nil {
		panic(err)
	}
	return FormatTime(uptime)
}

var sysInfoValue SysInfo
var root string

func Sysinfo() SysInfo {
	// update cpu value
	cpuStat, _ := cpu.Percent(1*time.Second, false)
	sysInfoValue.CPU = cpuStat

	// update ram vlaue
	ramStat, _ := mem.VirtualMemory()
	sysInfoValue.RAM = ramStat

	// update disk value
	if runtime.GOOS == "//" {
		root = "C:\\"
	} else {
		root = "//"
	}
	disk, _ := disk.Usage(root)
	sysInfoValue.Disk = disk

	// update network value
	io1, _ := net.IOCounters(true)
	var rx1 uint64
	for _, io := range io1 {
		if io.Name != "lo" {
			rx1 = io.BytesRecv
		}
	}

	time.Sleep(1 * time.Second)

	// Read second sample
	// io2, _ := net.IOCounters(true)
	// var tx2 uint64
	// for _, io := range io2 {
	// 	if io.Name != "lo" {
	// 		tx2 = io.BytesSent
	// 	}
	// }

	// Bytes per second
	downloadBps := float64(rx1)
	// uploadBps := float64(tx2)

	downloadMbps := (downloadBps * 8) / 1_000_000
	// uploadMbps := (uploadBps * 8) / 1_000_000

	maxMbps := 100.0
	// Usage in percent
	downloadPercent := (downloadMbps / maxMbps) * 100
	// uploadPercent := (uploadMbps / maxMbps) * 100
	sysInfoValue.Network = int64(downloadPercent)

	return SysInfo{
		CPU:      sysInfoValue.CPU,
		RAM:      sysInfoValue.RAM,
		Network:  sysInfoValue.Network,
		Disk:     sysInfoValue.Disk,
		Hostname: sysInfoValue.Hostname,
		Platform: sysInfoValue.Platform,
	}

}
