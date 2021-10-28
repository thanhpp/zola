package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"errors"

	"github.com/thanhpp/zola/config/shared"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
	"github.com/thanhpp/zola/pkg/booting"
	"github.com/thanhpp/zola/pkg/logger"
)

type HTTPServer struct {
	cfg  *shared.HTTPServerConfig
	app  application.Application
	auth *auth.AuthService
}

func NewHTTPServer(cfg *shared.HTTPServerConfig, app application.Application, authSrv *auth.AuthService) *HTTPServer {
	return &HTTPServer{
		cfg:  cfg,
		app:  app,
		auth: authSrv,
	}
}

func (s *HTTPServer) Start() (booting.Daemon, error) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port),
		Handler: s.newRouter(),
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
