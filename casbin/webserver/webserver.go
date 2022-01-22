package webserver

import (
	"context"
	"fmt"
	"net/http"

	"bitbucket.org/tysud/gt-casbin/core"
	"bitbucket.org/tysud/gt-casbin/utils"
	"bitbucket.org/tysud/gt-casbin/webserver/router"
)

var logger *core.LogFormat

func StartHTTP(ctx context.Context) (shutdown utils.Daemon, err error) {
	logger := core.GetLogger()
	serverConf := core.GetConfig().Web
	var (
		srvAddr = fmt.Sprintf("%s:%s", serverConf.Host, serverConf.Port)
	)

	// init server
	server := &http.Server{
		Addr:    srvAddr,
		Handler: router.NewRouter(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("Can not open HTTP Server. Error: %v", err)
			panic(err)
		}
	}()

	shutdown = func() {
		<-ctx.Done()

		if err := server.Shutdown(context.Background()); err != nil {
			logger.Errorf("Shutdown server error: %v", err)
			return
		}

		logger.Info("Gracefully shutdown server")
	}

	return shutdown, err
}
