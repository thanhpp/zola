package websocketserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/thanhpp/zola/config/shared"
	"github.com/thanhpp/zola/pkg/booting"
	"github.com/thanhpp/zola/pkg/logger"
)

type WebsocketServer struct {
	cfg *shared.HTTPServerConfig
}

func NewWebsocketServer(cfg *shared.HTTPServerConfig) *WebsocketServer {
	return &WebsocketServer{
		cfg: cfg,
	}
}

func (s *WebsocketServer) Start() (booting.Daemon, error) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port),
		Handler: newRouter(),
	}

	return func(ctx context.Context) (start func() error, cleanup func()) {
		start = func() error {
			err := server.ListenAndServe()
			if err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					return nil
				}

				return err
			}

			return nil
		}

		cleanup = func() {
			shutdownCtx, cancel := context.WithTimeout(
				ctx,
				time.Second*5,
			)
			defer cancel()

			if err := server.Shutdown(shutdownCtx); err != nil {
				logger.Errorf("cleanup httpserver %v", err)
				return
			}
		}

		return start, cleanup
	}, nil
}
