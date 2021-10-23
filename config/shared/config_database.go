package shared

import "fmt"

type DatabaseConfig struct {
	Host     string `mapstructure:"Host"`
	Port     string `mapstructure:"Port"`
	DBName   string `mapstructure:"DBName"`
	User     string `mapstructure:"User"`
	Pass     string `mapstructure:"Pass"`
	LogLevel string `mapstructure:"LogLevel"`
	LogColor bool   `mapstructure:"LogColor"`
}

func (cfg DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User,
		cfg.Pass,
		cfg.DBName,
		cfg.Host,
		cfg.Port,
	)
}
