package aucoconfig

import (
	"os"

	"github.com/thanhpp/zola/config/shared"
	"github.com/thanhpp/zola/pkg/logger"
)

type MainConfig struct {
	Name               string                  `mapstructure:"Name"`
	Log                logger.LogConfig        `mapstructure:"Log"`
	HTTPServer         shared.HTTPServerConfig `mapstructure:"HTTPServer"`
	Database           shared.DatabaseConfig   `mapstructure:"Database"`
	LacLongQuanService struct {
		Host string `mapstructure:"Host"`
	} `mapstructure:"LacLongQuanService"`
}

var cfg = new(MainConfig)

func Set(path string) error {
	return shared.ReadFromFile(path, cfg)
}

func SetFromENV(path string) error {
	if err := Set(path); err != nil {
		return err
	}

	dbHost := os.Getenv("AC_DB_HOST")
	if dbHost != "" {
		cfg.Database.Host = dbHost
	}

	dbPort := os.Getenv("AC_DB_PORT")
	if dbPort != "" {
		cfg.Database.Port = dbPort
	}

	dbUser := os.Getenv("AC_DB_USER")
	if dbUser != "" {
		cfg.Database.User = dbUser
	}

	dbPassword := os.Getenv("AC_DB_PASSWORD")
	if dbPassword != "" {
		cfg.Database.Pass = dbPassword
	}

	return nil
}

func Get() *MainConfig {
	return cfg
}
