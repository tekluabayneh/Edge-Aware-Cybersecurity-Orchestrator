package processes

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"sync"
	"time"
)

type ProcInfo struct {
	PID        int32
	Name       string
	CPUPercent float64
	Memory     *process.MemoryInfoStat
}

func Processes(ch chan []ProcInfo) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	for {
		select {
		case <-ctx.Done():
			return
		default:
			wg.Add(1)
			go func() {
				var ProcssessData []ProcInfo
				processes, err := process.Processes()
				if err != nil {
					fmt.Println("Error:", err)
				}
				for _, p := range processes {
					name, _ := p.Name()
					cpuPercent, _ := p.CPUPercent()
					memInfo, _ := p.MemoryInfo()
					payload := ProcInfo{
						PID:        p.Pid,
						Name:       name,
						CPUPercent: cpuPercent,
						Memory:     memInfo,
					}
					ProcssessData = append(ProcssessData, payload)
				}
				ch <- ProcssessData
				time.Sleep(20 * time.Second)
				wg.Done()
			}()
			wg.Wait()
		}
	}
}
