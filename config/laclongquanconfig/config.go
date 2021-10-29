package laclongquanconfig

import (
	"github.com/thanhpp/zola/config/shared"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
	"github.com/thanhpp/zola/pkg/logger"
)

var cfg = new(MainConfig)

type MainConfig struct {
	Name       string                  `mapstructure:"Name"`
	Log        logger.LogConfig        `mapstructure:"Log"`
	HTTPServer shared.HTTPServerConfig `mapstructure:"HTTPServer"`
	Database   shared.DatabaseConfig   `mapstructure:"Database"`
	JWT        auth.Config             `mapstructure:"JWT"`
}

func Set(path string) error {
	return shared.ReadFromFile(path, cfg)
}

func Get() *MainConfig {
	return cfg
}
