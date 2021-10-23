package logger

type LogConfig struct {
	Color      bool   `mapstructure:"Color"`
	LoggerName string `mapstructure:"LoggerName"`
	Level      string `mapstructure:"Level"`
}
