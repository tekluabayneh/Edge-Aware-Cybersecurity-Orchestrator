package system

import (
	"agent/internal/utils"
	"context"
	"math"
	"sync"
)

type GetSysInfotype struct {
	Uptime  string
	Cpu     []float64
	Ram     float64
	Disk    float64
	Network int64
}

func System(ch chan GetSysInfotype) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	for {
		select {
		case <-ctx.Done():
			return
		default:
			upTime := utils.GetUptime()
			cpuInfo := utils.Sysinfo().CPU
			ramInfo := utils.Sysinfo().RAM
			diskInfo := utils.Sysinfo().Disk
			netWorkInfo := utils.Sysinfo().Network

			payload := GetSysInfotype{
				Uptime:  upTime,
				Cpu:     cpuInfo,
				Ram:     ramInfo.UsedPercent,
				Disk:    math.Round(float64(diskInfo.UsedPercent)),
				Network: netWorkInfo,
			}
			ch <- payload
			wg.Wait()
		}
	}
}
