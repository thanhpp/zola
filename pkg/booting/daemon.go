package booting

import (
	"context"

	"github.com/thanhpp/zola/pkg/logger"
)

type Daemon func(ctx context.Context) (start func() error, cleanup func())

type DaemonManager struct {
	MainCtx     context.Context
	DaemonCtrls []*daemonController
	Terminated  chan struct{}
}

func NewDaemonManeger(ctx context.Context) *DaemonManager {
	return &DaemonManager{
		MainCtx:    ctx,
		Terminated: make(chan struct{}),
	}
}

func (m *DaemonManager) Start(daemons ...Daemon) {
	if m == nil {
		panic("nil daemon manager")
	}

	for i := range daemons {
		ctx, cancel := context.WithCancel(m.MainCtx)
		start, cleanup := daemons[i](ctx)

		m.DaemonCtrls = append(m.DaemonCtrls, &daemonController{
			ctx:     ctx,
			cancel:  cancel,
			cleanup: cleanup,
		})

		go func() {
			if err := start(); err != nil {
				logger.FatalFmt(*logger.NewLogFormat("daemon manager").SetMsg("start daemon").SetError(err))
				m.Terminated <- struct{}{}
				return
			}
			m.Terminated <- struct{}{}
		}()
	}
}

func (m *DaemonManager) Stop() {
	if m == nil {
		panic("nil daemon manager")
	}

	for i := len(m.DaemonCtrls) - 1; i >= 0; i-- {
		m.DaemonCtrls[i].cancel()
		m.DaemonCtrls[i].cleanup()

		<-m.Terminated // wait for termination
	}
}
