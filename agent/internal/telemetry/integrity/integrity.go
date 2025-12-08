package integrity

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Integrity(ch chan bool) {
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
				ch <- true
				fmt.Println("integrity scan")
				time.Sleep(5 * time.Second)
			}()
			wg.Wait()
		}
	}
}
