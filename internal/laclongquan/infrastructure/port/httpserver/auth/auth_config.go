package auth

type Config struct {
	ExpireDuration int    `mapstructure:"ExpireDuration"`
	Issuer         string `mapstructure:"Issuer"`
	RSAKeyName     string `mapstructure:"RSAKeyName"`
	RSAPubKeyPath  string `mapstructure:"RSAPubKeyPath"`
	RSAPriKeyPath  string `mapstructure:"RSAPriKeyPath"`
}
