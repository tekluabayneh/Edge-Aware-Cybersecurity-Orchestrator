package integrity

import (
	"context"
	"fmt"
	"time"
)

func Integrity(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("integrity scan")
			time.Sleep(5 * time.Second)
		}
	}
}
