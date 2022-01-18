package main

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/config/aucoconfig"
	"github.com/thanhpp/zola/internal/auco/infra/port/websocketserver"
	"github.com/thanhpp/zola/pkg/booting"
	"github.com/thanhpp/zola/pkg/logger"
)

func start(configPath string) {
	if err := aucoconfig.Set(configPath); err != nil {
		panic(errors.WithMessage(err, "set config"))
	}

	if err := logger.SetLog(&aucoconfig.Get().Log); err != nil {
		panic(errors.WithMessage(err, "set log"))
	}
	logger.Info("logger OK")

	wsServer := websocketserver.NewWebsocketServer(&aucoconfig.Get().HTTPServer)
	wsServerDaemon, err := wsServer.Start()
	if err != nil {
		panic(errors.WithMessage(err, "start websocket server"))
	}
	logger.Info("http server OK")

	mainCtx := context.Background()
	daemonMan := booting.NewDaemonManeger(mainCtx)

	logger.Info("starting daemons....")
	daemonMan.Start(wsServerDaemon)
	booting.WaitSignals(mainCtx)
	daemonMan.Stop()

	log.Println("shutdown", time.Now())
}
