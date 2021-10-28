package main

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/config/laclongquanconfig"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/adapter/gormdb"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
	"github.com/thanhpp/zola/pkg/booting"
	"github.com/thanhpp/zola/pkg/logger"
)

func start(configPath string) {
	if err := laclongquanconfig.Set(configPath); err != nil {
		panic(errors.WithMessage(err, "set config"))
	}

	if err := logger.SetLog(&laclongquanconfig.Get().Log); err != nil {
		panic(errors.WithMessage(err, "set log"))
	}
	logger.Info("logger OK")

	dbao, err := gormdb.NewDBAO(&laclongquanconfig.Get().Database)
	if err != nil {
		panic(errors.WithMessage(err, "start database"))
	}
	logger.Info("dbao OK")

	app := application.NewApplication(dbao.User)
	logger.Info("application OK")

	authSrv, err := auth.NewAuthService(
		&laclongquanconfig.Get().JWT,
		dbao.Auth,
	)
	if err != nil {
		panic(errors.WithMessage(err, "start auth service"))
	}
	logger.Info("auth service OK")

	httpServer := httpserver.NewHTTPServer(
		&laclongquanconfig.Get().HTTPServer,
		app,
		authSrv,
	)
	httpDaemon, err := httpServer.Start()
	if err != nil {
		panic(errors.WithMessage(err, "http daemon"))
	}
	logger.Info("http server OK")

	mainCtx := context.Background()
	daemonMan := booting.NewDaemonManeger(mainCtx)

	logger.Info("starting daemons....")
	daemonMan.Start(httpDaemon)
	booting.WaitSignals(mainCtx)
	daemonMan.Stop()

	log.Println("shutdown", time.Now())
}
