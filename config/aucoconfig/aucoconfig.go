package aucoconfig

import (
	"github.com/thanhpp/zola/config/shared"
	"github.com/thanhpp/zola/pkg/logger"
)

type MainConfig struct {
	Name       string                  `mapstructure:"Name"`
	Log        logger.LogConfig        `mapstructure:"Log"`
	HTTPServer shared.HTTPServerConfig `mapstructure:"HTTPServer"`
}

var cfg = new(MainConfig)

func Set(path string) error {
	return shared.ReadFromFile(path, cfg)
}

func Get() *MainConfig {
	return cfg
}
