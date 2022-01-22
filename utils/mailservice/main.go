package main

import (
	"flag"

	"github.com/vfluxus/mailservice/boot"
)

var (
	configPath = flag.String("config", "dev.yml", "config file")
)

func main() {
	flag.Parse()
	boot.Boot(*configPath)
}
