package booting

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pinezapple/LibraryProject20201/skeleton/model"
)

// BootstrapDaemons start daemons and handling os signal.
func BootstrapDaemons(ctx context.Context, daemonGenerators ...model.DaemonGenerator) {
	// os signal handling
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	// create daemon manager
	manager := newDaemonManager(ctx)
	manager.registerDaemons(daemonGenerators...)

	// wait os signal
	select {
	case <-sigs:
	case <-manager.ctx.Done():
	}

	// stop daemons
	manager.stop()
}

type daemonManager struct {
	ctx          context.Context
	daemonCtx    []context.Context
	daemonCancel []context.CancelFunc
	terminated   chan struct{}
}

func newDaemonManager(ctx context.Context) (d *daemonManager) {
	d = &daemonManager{
		ctx:        ctx,
		terminated: make(chan struct{}, 8),
	}
	return
}

func (d *daemonManager) registerDaemon(g model.DaemonGenerator) {
	if g != nil {
		ctx, cancel := context.WithCancel(context.Background())
		d.daemonCtx, d.daemonCancel = append(d.daemonCtx, ctx), append(d.daemonCancel, cancel)

		daemon, err := g(ctx)
		if err != nil {
			panic(err)
		}

		// start daemon
		go func() {
			daemon()
			d.terminated <- struct{}{}
		}()
	}
}

func (d *daemonManager) registerDaemons(gs ...model.DaemonGenerator) {
	for _, g := range gs {
		d.registerDaemon(g)
	}
}

func (d *daemonManager) stop() {
	for i := len(d.daemonCtx) - 1; i >= 0; i-- {
		d.daemonCancel[i]()
		<-d.terminated // wait termination
	}
}
