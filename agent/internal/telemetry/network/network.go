package network

import (
	"context"
	"fmt"
	"time"
)

func Network(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("network second")
			time.Sleep(5 * time.Second)
		}
	}
}
