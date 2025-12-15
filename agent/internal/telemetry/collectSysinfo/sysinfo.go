package collectsysinfo

import (
	"fmt"
	"math"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetSysInfo(uptime string, cpu []float64, ram *mem.VirtualMemoryStat, diskinfo *disk.UsageStat, network int64) {
	fmt.Printf("cpu is  %.0f%%\n", math.Round(cpu[0]))
	fmt.Printf("ram is %d%%\n", uint64(ram.UsedPercent))
	fmt.Printf("disk is %v%%\n", math.Round(diskinfo.UsedPercent))
	fmt.Printf("network is %d%%\n", network)
	fmt.Printf("uptime is %v\n", uptime)
}
