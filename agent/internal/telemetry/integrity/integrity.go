package integrity

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

type IntegritySnapshot struct {
	OS            string         `json:"_os"`
	KernelVersion string         `json:"kernel_version"`
	PatchLevel    string         `json:"patch_level"`
	CriticalFiles map[string]any `json:"critical_files"`
	CollectedAt   int64          `json:"collected_at"`
}

func Integrity(ch chan IntegritySnapshot) {
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
				info, err := host.Info()
				if err != nil {
					fmt.Println("error readin host info", err)
				}
				criticalFiles := map[string]any{}
				var filesToCheck []string

				switch runtime.GOOS {
				case "linux":
					filesToCheck = []string{
						"/etc/shadow",
						"/bin/bash",
						"/usr/bin/sudo",
					}
				case "windows":
					filesToCheck = []string{
						`C:\Windows\System32\ntoskrnl.exe`,
						`C:\Windows\System32\cmd.exe`,
						`C:\Windows\System32\drivers\`,
					}
				case "darwin":
					filesToCheck = []string{
						"/System/Library/Kernels/kernel",
						"/usr/bin/sudo",
						"/bin/zsh",
					}
				default:
					fmt.Println("Unsupported OS:", runtime.GOOS)
					return
				}

				for _, path := range filesToCheck {
					fi, err := os.Stat(path)
					if err != nil {
						continue
					}
					modTime := fi.ModTime()
					criticalFiles[path] = modTime
				}
				payload := IntegritySnapshot{
					OS:            info.OS,
					KernelVersion: info.KernelVersion,
					PatchLevel:    info.PlatformVersion,
					CriticalFiles: criticalFiles,
					CollectedAt:   time.Now().Unix(),
				}
				ch <- payload
				time.Sleep(20 * time.Second)
				defer wg.Done()
			}()
			wg.Wait()
		}
	}
}
