package main

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/config/aucoconfig"
	"github.com/thanhpp/zola/internal/auco/app"
	"github.com/thanhpp/zola/internal/auco/infra/adapter/gormdb"
	"github.com/thanhpp/zola/internal/auco/infra/port/websocketserver"
	"github.com/thanhpp/zola/pkg/booting"
	"github.com/thanhpp/zola/pkg/logger"
)

func start(configPath string) {
	if err := aucoconfig.SetFromENV(configPath); err != nil {
		panic(errors.WithMessage(err, "set config"))
	}

	if err := logger.SetLog(&aucoconfig.Get().Log); err != nil {
		panic(errors.WithMessage(err, "set log"))
	}
	logger.Info("logger OK")

	err := gormdb.InitConnection(aucoconfig.Get().Database.DSN(), aucoconfig.Get().Log.Level, aucoconfig.Get().Log.Color)
	if err != nil {
		panic(errors.WithMessage(err, "init connection"))
	}
	logger.Info("dbao OK")

	wmManager := app.NewWsManager(gormdb.NewGormDB(), gormdb.NewGormDB())

	wsServer := websocketserver.NewWebsocketServer(&aucoconfig.Get().HTTPServer, wmManager)
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
