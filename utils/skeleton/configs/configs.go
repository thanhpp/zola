package configs

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/pinezapple/LibraryProject20201/skeleton/libs"
)

// ErrOnlySupportMysql indicate that we only support mysql
var ErrOnlySupportMysql = fmt.Errorf("Support mysql only")

// ErrOnlySupportClickhouse indicate that we only support clickhouse
var ErrOnlySupportClickhouse = fmt.Errorf("Support clickhouse only")

// GRPCServerConfig gRPC server configuration
type GRPCServerConfig struct {
	PublicIP          string
	Port              int
	ClientCA          string
	Cert              string
	Key               string
	MaxConnectionIdle time.Duration
}

// MysqlConnConfig mysql database connection configuration
type MysqlConnConfig struct {
	Type        string
	DB          string
	Username    string
	Password    string
	Masters     []string
	Slaves      []string
	CaCert      string
	ClientCert  string
	ClientKey   string
	TLS         string
	Args        string
	IsWsrep     bool
	MaxIdleConn int
	MaxOpenConn int
}

// MysqlTLSRegister registers a custom tls.Config to be used with sql.Open.
// Use the key as a value in the DSN where tls=value.
type MysqlTLSRegister func(key string, config *tls.Config) error

// Construct database connection
func (c *MysqlConnConfig) Construct(register MysqlTLSRegister) (err error) {
	if c.Type != "mysql" {
		err = ErrOnlySupportMysql
		return
	}

	if len(c.TLS) > 0 {
		if strings.HasSuffix(c.Args, "&") {
			c.Args += "tls=" + c.TLS
		} else {
			c.Args += "&tls=" + c.TLS
		}

		// Load client cert
		clientCert, err := tls.LoadX509KeyPair(c.ClientCert, c.ClientKey)
		if err != nil {
			return err
		}

		rootCertPool, err := libs.LoadCACertPool([]string{c.CaCert})
		if err != nil {
			return err
		}

		err = register(c.TLS, &tls.Config{
			RootCAs:                  rootCertPool,
			Certificates:             []tls.Certificate{clientCert},
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: false,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
			InsecureSkipVerify: true,
		})

		if err != nil {
			return err
		}
	}

	if c.Masters != nil && len(c.Masters) > 0 {
		for i := range c.Masters {
			c.Masters[i] = fmt.Sprintf("%s:%s@(%s)/%s?%s", c.Username, c.Password, c.Masters[i], c.DB, c.Args)
		}
	}

	if c.Slaves != nil && len(c.Slaves) > 0 {
		for i := range c.Slaves {
			c.Slaves[i] = fmt.Sprintf("%s:%s@(%s)/%s?%s", c.Username, c.Password, c.Slaves[i], c.DB, c.Args)
		}
	}

	return
}

// Equal check two db configuration are equal
func (c *MysqlConnConfig) Equal(x *MysqlConnConfig) bool {
	if x == nil {
		return false
	}

	// check masters equal
	if len(c.Masters) != len(x.Masters) {
		return false
	}
	for i := range c.Masters {
		if c.Masters[i] != x.Masters[i] {
			return false
		}
	}

	// check slave equal
	if len(c.Slaves) != len(x.Slaves) {
		return false
	}
	for i := range c.Slaves {
		if c.Slaves[i] != x.Slaves[i] {
			return false
		}
	}

	tmp := c.IsWsrep == x.IsWsrep && c.DB == x.DB && c.Type == x.Type && c.Username == x.Username && c.Password == x.Password
	if !tmp {
		return false
	}

	return c.TLS == x.TLS && c.CaCert == x.CaCert && c.ClientCert == x.ClientCert && c.ClientKey == x.ClientKey && c.Args == x.Args
}

// ClickhouseConnConfig clickhouse database connection configuration
type ClickhouseConnConfig struct {
	Type        string
	DB          string
	Username    string
	Password    string
	Masters     []string
	Slaves      []string
	CaCert      string
	ClientCert  string
	ClientKey   string
	Args        string
	IsWsrep     bool
	MaxIdleConn int
	MaxOpenConn int
}

// Construct database connection
func (c *ClickhouseConnConfig) Construct() (err error) {
	if c.Type != "clickhouse" {
		err = ErrOnlySupportClickhouse
		return
	}

	if c.Masters != nil && len(c.Masters) > 0 {
		for i := range c.Masters {
			c.Masters[i] = fmt.Sprintf("tcp://%s?debug=false&username=%s&password=%s&database=%s&%s", c.Masters[i], c.Username, c.Password, c.DB, c.Args)
		}
	}

	if c.Slaves != nil && len(c.Slaves) > 0 {
		for i := range c.Slaves {
			c.Slaves[i] = fmt.Sprintf("tcp://%s?debug=false&username=%s&password=%s&database=%s&%s", c.Slaves[i], c.Username, c.Password, c.DB, c.Args)
		}
	}

	return
}

// Equal check two db configuration are equal
func (c *ClickhouseConnConfig) Equal(x *ClickhouseConnConfig) bool {
	if x == nil {
		return false
	}

	// check masters equal
	if len(c.Masters) != len(x.Masters) {
		return false
	}
	for i := range c.Masters {
		if c.Masters[i] != x.Masters[i] {
			return false
		}
	}

	// check slave equal
	if len(c.Slaves) != len(x.Slaves) {
		return false
	}
	for i := range c.Slaves {
		if c.Slaves[i] != x.Slaves[i] {
			return false
		}
	}

	tmp := c.IsWsrep == x.IsWsrep && c.DB == x.DB && c.Type == x.Type && c.Username == x.Username && c.Password == x.Password
	if !tmp {
		return false
	}

	// tmp = c.TLS == x.TLS && c.CaCert == x.CaCert && c.ClientCert == x.ClientCert && c.ClientKey == x.ClientKey && c.Args == x.Args
	return c.Args == x.Args
}

// HTTPClientConf  ...
type HTTPClientConf struct {
	ClientCert string
	ClientKey  string
	ServerCAs  string // csv list of trusted CAs
}

// HTTPServerConf binding configuration for webserver
type HTTPServerConf struct {
	PublicIP  string
	Port      int
	Cert      string
	Key       string
	ClientCAs string // csv list of trusted CAs
}

// KafkaConnection ...
type KafkaConnection struct {
	Brokers []string
}

// RedisConnection ...
type RedisConnection struct {
	Address  string
	Password string
	DB       int
}
