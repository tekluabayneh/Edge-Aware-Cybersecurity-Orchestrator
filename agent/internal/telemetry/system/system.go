package system

import (
	"agent/internal/utils"
	"context"
	"math"
	"sync"
	"time"
)

type GetSysInfotype struct {
	Uptime  string    `json:"uptime"`
	Cpu     []float64 `json:"cpu"`
	Ram     float64   `json:"ram"`
	Disk    float64   `json:"disk"`
	Network int64     `json:"network"`
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
			time.Sleep(20 * time.Second)
		}
	}
}
