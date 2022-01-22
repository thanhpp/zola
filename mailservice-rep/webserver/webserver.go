package webserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vfluxus/dvergr/logger"
	"github.com/vfluxus/mailservice/webserver/router"
)

func StartHTTP(ctx context.Context, host string, port string) (shutdown func(), err error) {
	var (
		srvAddr = fmt.Sprintf("%s:%s", host, port)
	)

	// init server
	server := &http.Server{
		Addr:    srvAddr,
		Handler: router.NewRouter(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Get().Errorf("Can not open HTTP Server. Error: %v", err)
			panic(err)
		}
	}()

	shutdown = func() {
		if err := server.Shutdown(ctx); err != nil {
			logger.Get().Errorf("Shutdown server error: %v", err)
			return
		}

		logger.Get().Info("Gracefully shutdown server")
	}

	return shutdown, err
}
