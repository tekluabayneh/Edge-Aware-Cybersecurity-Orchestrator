package security

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Security(ch chan bool) {
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
				fmt.Println("test Security")
				ch <- true
				time.Sleep(5 * time.Second)
				wg.Done()
			}()
			wg.Wait()
		}
	}
}
