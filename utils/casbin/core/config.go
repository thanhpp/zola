package core

import (
	"errors"
	"strings"

	"bitbucket.org/tysud/gt-casbin/utils"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MainConfig struct {
	Name             string          `mapstructure:"NAME"`
	Environment      string          `mapstructure:"ENVIRONMENT"`
	DBdns            string          `mapstructure:"DB_DNS"`
	CasbinConfigFile string          `mapstructure:"CASBIN_CONFIG"`
	FileMode         fileModeConfig  `mapstructure:"FILE_MODE"`
	PolicyCSVFile    string          `mapstructure:"POLICY_CSV_FILE"`
	DefaultRole      []string        `mapstructure:"DEFAULT_ROLES"`
	Log              logConfig       `mapstructure:"LOG"`
	Web              webServerConfig `mapstructure:"WEB"`
}

type fileModeConfig struct {
	ReadFromFile bool `mapstructure:"READ_FROM_FILE"`
	WriteToDB    bool `mapstructure:"WRITE_TO_DB"`
}

type webServerConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type logConfig struct {
	Level string `mapstructure:"LEVEL"`
	Color bool   `mapstructure:"COLOR"`
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

	return nil
}

func InitDB(connString string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connString), &gorm.Config{})
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

	if mainConfig.FileMode.ReadFromFile {
		// if read from file but not write to db (enable for fast testing)
		if !mainConfig.FileMode.WriteToDB {
			SetFileMode(true)
			e, err := casbin.NewEnforcer(mainConfig.CasbinConfigFile, mainConfig.PolicyCSVFile)
			if err != nil {
				return err
			}
			SetCasbinEnforcer(e)
		} else {
			SetFileMode(false)
			rawdata, err := utils.ReadCsvFile(mainConfig.PolicyCSVFile)
			if err != nil {
				return err
			}
			db, err := InitDB(mainConfig.DBdns)
			if err != nil {
				return err
			}

			// casbin
			a, err := gormadapter.NewAdapterByDBUseTableName(db, "", "casbin")
			if err != nil {
				return err
			}
			e, err := casbin.NewEnforcer(mainConfig.CasbinConfigFile, a)
			if err != nil {
				return err
			}
			SetCasbinEnforcer(e)

			err = insertCasbinRuleFromFile(rawdata)
			if err != nil {
				return err
			}
			/*
							namedPolicy := e.GetNamedPolicy("p")
							fmt.Println(namedPolicy)
							allNamedRoles := e.GetAllNamedRoles("g")
							fmt.Println(allNamedRoles)
				filteredGroupingPolicy := e.GetFilteredGroupingPolicy(1, "admin")
						fmt.Println(filteredGroupingPolicy)
			*/
		}
	} else {
		SetFileMode(false)
		db, err := InitDB(mainConfig.DBdns)
		if err != nil {
			return err
		}

		// casbin
		a, err := gormadapter.NewAdapterByDBUseTableName(db, "", "casbin")
		if err != nil {
			return err
		}
		e, err := casbin.NewEnforcer(mainConfig.CasbinConfigFile, a)
		if err != nil {
			return err
		}
		SetCasbinEnforcer(e)
	}
	return nil
}
