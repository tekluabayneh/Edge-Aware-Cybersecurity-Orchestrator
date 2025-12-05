package security

import (
	"context"
	"fmt"
	"time"
)

func Security(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("test Security")
			time.Sleep(5 * time.Second)
		}
	}
}
