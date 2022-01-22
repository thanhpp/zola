package main

import (
	"flag"

	"bitbucket.org/tysud/gt-casbin/boot"
)

var (
	configPath = flag.String("config", "dev.yml", "config file")
)

func main() {
	flag.Parse()
	boot.Boot(*configPath)
}
