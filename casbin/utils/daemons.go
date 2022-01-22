package utils

import "context"

// Daemon abstract a daemon.
type Daemon func()

// DaemonGenerator generates a Daemon
type DaemonGenerator func(ctx context.Context) (Daemon, error)
