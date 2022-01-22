package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/skeleton/libs"
)

// Daemon abstract a daemon.
type Daemon func()

// DaemonGenerator generates a Daemon
type DaemonGenerator func(ctx context.Context) (Daemon, error)

// RPCErrCode rpc error code
type RPCErrCode int32

const (
	// RPCSuccess request success without any error
	RPCSuccess RPCErrCode = iota
	// RPCCustomErr custom error
	RPCCustomErr
	// RPCInternalServerErr internal server error
	RPCInternalServerErr
)

// codes for lead status update service response
const (
	CustomRPCErr RPCErrCode = -1

	UpdateSuccess RPCErrCode = 0

	InvalidLeadIDErr RPCErrCode = 1

	NotClientLeadErr RPCErrCode = 2

	InvalidLeadInfoErr RPCErrCode = 3
)

// ExecTransaction template to execute transaction
func ExecTransaction(ctx context.Context, tx *sql.Tx, exec func(ctx context.Context, tx *sql.Tx) error) (err error) {
	if err = exec(ctx, tx); err != nil {
		_ = tx.Rollback()
		return
	}

	return tx.Commit()
}

// LogFormat log pattern for whole system
type LogFormat struct {
	ServiceName string      `json:"srv"`
	Source      string      `json:"src,omitempty"`
	Action      string      `json:"act,omitempty"`
	Data        interface{} `json:"dat,omitempty"`   // data
	Err         interface{} `json:"err,omitempty"`   // error
	Success     interface{} `json:"suc,omitempty"`   // success
	Stack       interface{} `json:"stack,omitempty"` // stack trace
	Message     string      `json:"msg,omitempty"`
}

// ToMapStringItf ...
func (lg *LogFormat) ToMapStringItf() map[string]interface{} {
	return map[string]interface{}{
		"srv":   lg.ServiceName,
		"src":   lg.Source,
		"act":   lg.Action,
		"dat":   lg.Data,
		"err":   lg.Err,
		"suc":   lg.Success,
		"stack": lg.Stack,
	}
}

// ShardAddress address of shard, point to lead service
type ShardAddress struct {
	ID      uint64 `json:"id"`
	Address string `json:"address"`
}

//---------------------------------------- portal model ----------------------------------------------------------
// Claim jwt claim
type Claim struct {
	Username   string
	Group      []string
	Permission [][]string
	jwt.StandardClaims
}

// SecureCookieConfig secure cookie middleware configuration
type SecureCookieConfig struct {
	HashKey    []byte
	BlockKey   []byte
	CookieName string
	ContextKey string
}

type cookieValidator struct {
	secureCookie *securecookie.SecureCookie
	config       *SecureCookieConfig
}

func (c cookieValidator) MakeSecureCookie(val string) (*http.Cookie, error) {
	if c.secureCookie == nil || c.config == nil {
		return nil, fmt.Errorf("CookieValidator not initialized")
	}

	encoded, err := c.secureCookie.Encode(c.config.CookieName, val)
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:  c.config.CookieName,
		Value: encoded,
	}, nil
}

func (c cookieValidator) ExpireSecureCookie() (*http.Cookie, error) {
	if c.secureCookie == nil || c.config == nil {
		return nil, fmt.Errorf("CookieValidator not initialized")
	}

	return &http.Cookie{
		Name:   c.config.CookieName,
		MaxAge: -1,
	}, nil
}

// CookieValidator ...
var CookieValidator = cookieValidator{}

func readSecureCookie(secureCookie *securecookie.SecureCookie, c echo.Context, cookieName string) (value string, err error) {
	_cookie, err := c.Cookie(cookieName)
	if err != nil {
		return "", err
	}
	cookie := _cookie.Value

	var val string
	err = secureCookie.Decode(cookieName, cookie, &val)
	value = val

	return
}

// NewSecureCookieMW new secure cookie middleware
func NewSecureCookieMW(config SecureCookieConfig) echo.MiddlewareFunc {
	CookieValidator.secureCookie = securecookie.New(config.HashKey, config.BlockKey)

	if len(config.ContextKey) == 0 {
		config.ContextKey = "USER"
	}

	if len(config.CookieName) == 0 {
		config.CookieName = "auth"
	}

	CookieValidator.config = &config

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if cookie, err := readSecureCookie(CookieValidator.secureCookie, c, CookieValidator.config.CookieName); err == nil {
				c.Set(config.ContextKey, cookie)
			} else {
				return echo.ErrUnauthorized
			}

			// Continue the chain of middleware
			return next(c)
		}
	}
}

// User model user
type User struct {
	ID          uint64    `json:"id_user" db:"id_user"`
	Username    string    `json:"username" db:"username"`
	Name        string    `json:"name" db:"name"`
	Role        string    `json:"role" db:"role"`
	Dob         string    `json:"dob" db:"dob"`
	Sex         string    `json:"sex" db:"sex"`
	PhoneNumber string    `json:"phonenumber" db:"phonenumber"`
	Status      byte      `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Checksum    uint64    `json:"checksum" db:"checksum"`
}

// Sum calculate sip hash sum
func (c User) Sum(k0, k1 uint64) uint64 {
	sum := libs.ConcatCopyPreAllocate([][]byte{
		[]byte(strconv.FormatUint(c.ID, 10)),
		{c.Status},
		[]byte(c.Username),
	})

	return uint64(libs.SipHash48(k0, k1, []byte(sum)))
}

// ValidateChecksum validate record checksum
func (c User) ValidateChecksum(k0, k1 uint64) bool {
	return c.Sum(k0, k1) == c.Checksum
}

// UserSecurity user security table
type UserSecurity struct {
	Username  string
	Gr        uint64
	Role      uint64
	Password  []byte
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Checksum  uint64
}

// Sum do sum on user security
func (c *UserSecurity) Sum(k0, k1 uint64) uint64 {
	sum := libs.ConcatCopyPreAllocate([][]byte{
		[]byte(c.Username),
		libs.Uint64ToBytes(c.Gr),
		libs.Uint64ToBytes(c.Role),
		c.Password,
	})

	return uint64(libs.SipHash48(k0, k1, []byte(sum)))
}

// ValidateChecksum validate checksum
func (c *UserSecurity) ValidateChecksum(k0, k1 uint64) bool {
	return c.Sum(k0, k1) == c.Checksum
}
