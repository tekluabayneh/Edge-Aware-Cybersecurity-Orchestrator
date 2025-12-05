package processes

import (
	"context"
	"fmt"
	"time"
)

func Processes(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("test Processes")
			time.Sleep(5 * time.Second)
		}
	}
}
