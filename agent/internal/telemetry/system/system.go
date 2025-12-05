package system

import (
	"context"
	"fmt"
	"time"
)

func System(ctx context.Context) {
	for {
		select {

		case <-ctx.Done():
			return
		default:
			fmt.Println("test system")
			time.Sleep(5 * time.Second)
		}
	}
}
