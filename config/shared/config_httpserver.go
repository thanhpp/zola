package shared

type HTTPServerConfig struct {
	Host   string `mapstructure:"Host"`
	Port   string `mapstructure:"Port"`
	Domain string `mapstructure:"Domain"`
}
