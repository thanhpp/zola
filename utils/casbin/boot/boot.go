package boot

import (
	"context"
	"log"

	"bitbucket.org/tysud/gt-casbin/core"
	"bitbucket.org/tysud/gt-casbin/webserver"
)

func Boot(configPath string) {
	var ctx = context.Background()

	// Config
	if err := core.SetConfig(configPath); err != nil {
		log.Fatalf("Get config error: %v \n", err)
	}

	// Logger
	core.InitLogger(core.GetConfig())
	BootstrapDaemons(ctx, webserver.StartHTTP)
}
