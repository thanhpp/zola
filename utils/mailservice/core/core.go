package core

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

type MainConfig struct {
	Name        string          `mapstructure:"NAME"`
	Environment string          `mapstructure:"ENVIRONMENT"`
	KeyPath     string          `mapstructure:"ENCRYPTKEYPATH"`
	Key         string          `mapstructure:"KEY"`
	Log         logConfig       `mapstructure:"LOG"`
	Web         webServerConfig `mapstructure:"WEB"`
	DB          databaseConfig  `mapstructure:"DB"`
	Auth        AuthConfig      `mapstructure:"AUTH"`
}

type logConfig struct {
	Level string `mapstructure:"LEVEL"`
	Color bool   `mapstructure:"COLOR"`
}

type webServerConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type databaseConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	Database string `mapstructure:"Database"`
}

type AuthConfig struct {
	JWKSUrl   string   `mapstructure:"JWKS_URL"`
	Audiences []string `mapstructure:"AUDIENCES"`
	Issuers   []string `mapstructure:"ISSUERS"`
	Purpose   string   `mapstructure:"PURPOSE"`
}

var mainConfig = new(MainConfig)

// ------------------------------
// GetConfig Return config object
func GetConfig() *MainConfig {
	return mainConfig
}

// ------------------------------
// SetConfig read config from filepath
func SetConfig(configPath string) (err error) {
	if err = readConfigFromFile(configPath); err != nil {
		return err
	}
	key, err := ioutil.ReadFile(mainConfig.KeyPath)
	if err != nil {
		return err
	}
	mainConfig.Key = string(key)
	return nil
}

// ------------------------------
// readConfigFromFile read config by viper
func readConfigFromFile(configPath string) (err error) {
	v := viper.New()
	configPart := strings.Split(configPath, ".")
	if len(configPart) > 2 {
		return errors.New("Unacceptable file path format. Require ***.***")
	}
	v.SetConfigName(configPart[0])
	v.SetConfigType(configPart[1])

	// add config path
	v.AddConfigPath(".")
	v.AddConfigPath("../")
	v.AddConfigPath("../../")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.UnmarshalExact(mainConfig); err != nil {
		return err
	}

	return nil
}
