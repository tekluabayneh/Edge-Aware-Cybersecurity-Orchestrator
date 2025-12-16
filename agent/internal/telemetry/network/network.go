package network

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Network(ch chan bool) {
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
				susIp := FilterSusIp()
				fmt.Println(susIp)
				fmt.Println("network second")
				time.Sleep(5 * time.Second)
				ch <- true
				defer wg.Done()
			}()
			wg.Wait()
		}
	}
}
