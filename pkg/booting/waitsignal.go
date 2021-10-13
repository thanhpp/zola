package booting

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func WaitSignals(ctx context.Context) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	select {
	case <-sigChan:
		return
	case <-ctx.Done():
		return
	}
}
