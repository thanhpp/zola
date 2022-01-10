package laclongquanconfig

import (
	"os"

	"github.com/thanhpp/zola/config/shared"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
	"github.com/thanhpp/zola/pkg/logger"
)

var cfg = new(MainConfig)

type MainConfig struct {
	Name          string                  `mapstructure:"Name"`
	Log           logger.LogConfig        `mapstructure:"Log"`
	HTTPServer    shared.HTTPServerConfig `mapstructure:"HTTPServer"`
	Database      shared.DatabaseConfig   `mapstructure:"Database"`
	JWT           auth.Config             `mapstructure:"JWT"`
	SaveDirectory string                  `mapstructure:"SaveDirectory"`
	AESKey        string                  `mapstructure:"AESKey"`
	Admins        []AdminAccount          `mapstructure:"Admins"`
	ESClient      ESClientConfig          `mapstructure:"ESClient"`
}

type AdminAccount struct {
	Phone string `mapstructure:"Phone"`
	Pass  string `mapstructure:"Pass"`
}

func Set(path string) error {
	return shared.ReadFromFile(path, cfg)
}

func SetFromENV(path string) error {
	if err := Set(path); err != nil {
		return err
	}

	dbHost := os.Getenv("LLQ_DB_HOST")
	if dbHost != "" {
		cfg.Database.Host = dbHost
	}

	dbPort := os.Getenv("LLQ_DB_PORT")
	if dbPort != "" {
		cfg.Database.Port = dbPort
	}

	dbUser := os.Getenv("LLQ_DB_USER")
	if dbUser != "" {
		cfg.Database.User = dbUser
	}

	dbPassword := os.Getenv("LLQ_DB_PASSWORD")
	if dbPassword != "" {
		cfg.Database.Pass = dbPassword
	}

	return nil
}

func Get() *MainConfig {
	return cfg
}
