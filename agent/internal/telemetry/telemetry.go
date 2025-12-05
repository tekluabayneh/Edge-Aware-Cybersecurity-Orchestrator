package telemetry

import (
	"agent/internal/telemetry/integrity"
	"agent/internal/telemetry/network"
	"agent/internal/telemetry/processes"
	"agent/internal/telemetry/security"
	"agent/internal/telemetry/system"
	"context"
	"sync"
)

func Telemetry() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	wg.Add(5)
	go func() { defer wg.Done(); security.Security(ctx) }()

	go func() { defer wg.Done(); network.Network(ctx) }()

	go func() { defer wg.Done(); processes.Processes(ctx) }()

	go func() { defer wg.Done(); integrity.Integrity(ctx) }()

	go func() { defer wg.Done(); system.System(ctx) }()

	wg.Wait()
}
