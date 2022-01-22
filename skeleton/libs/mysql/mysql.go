package mysql

import (
	"context"
	"net"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
)

// RegisterDial register tcp dial for mysql
func RegisterDial() {
	mysqlDriver.RegisterDialContext("tcp", func(ctx context.Context, addr string) (net.Conn, error) {
		if ctx == nil {
			ctx = context.Background()
		}

		var dialer net.Dialer
		conn, err := dialer.DialContext(ctx, "tcp", addr)
		if err != nil {
			return nil, err
		}

		type keepAliveSetter interface {
			SetKeepAlive(keepalive bool) error
			SetKeepAlivePeriod(d time.Duration) error
		}

		if setter, ok := conn.(keepAliveSetter); ok {
			if err = setter.SetKeepAlive(true); err == nil {
				_ = setter.SetKeepAlivePeriod(time.Second * 10)
			}
		}

		return conn, nil
	})
}
