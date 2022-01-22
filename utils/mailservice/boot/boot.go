package boot

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vfluxus/mailservice/webserver/middlewares"

	"github.com/vfluxus/dvergr/logger"
	"github.com/vfluxus/mailservice/core"
	"github.com/vfluxus/mailservice/repository"
	"github.com/vfluxus/mailservice/webserver"
)

func Boot(configPath string) {
	var ctx = context.Background()

	// Config
	if err := core.SetConfig(configPath); err != nil {
		log.Fatalf("Get config error: %v \n", err)
	}

	// Logger
	logConf := core.GetConfig().Log
	if err := logger.Set("ZAP", core.GetConfig().Name, "DEVELOPMENT", logConf.Level, logConf.Color); err != nil {
		log.Fatalf("Set logger error: %v", err)
	}

	// Database
	dbConf := core.GetConfig().DB
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbConf.Username, dbConf.Password, dbConf.Database, dbConf.Host, dbConf.Port)
	if err := repository.GetDAO().SetDBConnection(dsn, "INFO"); err != nil {
		log.Fatalf("Set DB connection error: %v", err)
	}
	if err := repository.GetDAO().AutoMigrate(ctx); err != nil {
		log.Fatalf("DB auto migrate error: %v", err)
	}

	// Web
	// -- Auth middleware
	authConf := core.GetConfig().Auth
	middlewares.InitAuthN(&authConf)

	// -- Server
	serverConf := core.GetConfig().Web
	shutdown, err := webserver.StartHTTP(ctx, serverConf.Host, serverConf.Port)
	if err != nil {
		log.Fatalf("Start HTTP server error: %v", err)
	}

	// handle shutdown signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	// after os signal
	<-sigs
	shutdown()
}
