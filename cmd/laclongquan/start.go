package main

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/config/laclongquanconfig"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/accountcipher"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/adapter/esclient"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/adapter/gormdb"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
	"github.com/thanhpp/zola/pkg/booting"
	"github.com/thanhpp/zola/pkg/logger"
)

func start(configPath string) {
	if err := laclongquanconfig.SetFromENV(configPath); err != nil {
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

	accountCipher, err := accountcipher.New(laclongquanconfig.Get().AESKey)
	if err != nil {
		panic(errors.WithMessage(err, "new account cipher"))
	}
	logger.Info("account cipher OK")

	esClient := esclient.NewEsClient(laclongquanconfig.Get().ESClient.Host)
	app := application.NewApplication(
		accountCipher,
		dbao.User,
		dbao.Post, laclongquanconfig.Get().SaveDirectory,
		dbao.Report,
		dbao.Like,
		dbao.Relation,
		dbao.Comment,
		esClient,
	)
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

	// ------------- Migrate data from config -------------
	createAdminCtx := context.Background()
	for _, adminAcc := range laclongquanconfig.Get().Admins {
		if err := app.UserHandler.CreateAdminUser(createAdminCtx, adminAcc.Phone, adminAcc.Pass, "", ""); err != nil {
			logger.Errorf("can't create admin user %s - err: %v", adminAcc.Phone, err)
		}
	}

	// ---------------------- Sync all user to ES ----------------------
	if err := app.UserHandler.SyncAllUser(); err != nil {
		logger.Errorf("can't sync all user to ES - err: %v", err)
	}

	// ------------- Daemons ---------------
	mainCtx := context.Background()
	daemonMan := booting.NewDaemonManeger(mainCtx)

	logger.Info("starting daemons....")
	daemonMan.Start(httpDaemon, authSrv.DeleteExpiredDaemons())
	booting.WaitSignals(mainCtx)
	daemonMan.Stop()

	log.Println("shutdown", time.Now())
}
