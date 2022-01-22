package middlewares

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uc-cdis/go-authutils/authutils"
	"github.com/vfluxus/dvergr/logger"
	"github.com/vfluxus/mailservice/core"
)

// jwt authutils
var (
	jwtApp   = new(authutils.JWTApplication)
	expected = new(authutils.Expected)
)

// authutils const
const (
	realmAuthenRequire = "Authentication required"
	realmAuthenPrefix  = "Authen realm="
	authorization      = "Authorization"
)

// errors
var (
	ErrNilContext   = errors.New("Nil Context")
	ErrEmptyToken   = errors.New("Empty authHeader")
	ErrUnauthorized = errors.New("Unauthorized")
)

// ------------------------------
// InitAuthN init JWT validator from config
func InitAuthN(config *core.AuthConfig) {
	jwtApp = authutils.NewJWTApplication(config.JWKSUrl)
	expected = &authutils.Expected{
		Audiences: config.Audiences,
		Issuers:   config.Issuers,
		Purpose:   &config.Purpose,
	}
}

// ------------------------------
// validateToken validate token using authutils from uc-cdis
func validateToken(c *gin.Context) (realm string, err error) {
	if c == nil {
		return "", ErrNilContext
	}

	realm = realmAuthenPrefix + strconv.Quote(realmAuthenRequire)

	authHeader := c.GetHeader(authorization)
	if authHeader == "" {
		logger.Get().Errorf("Authorization header: ", authHeader)
		return realm, ErrUnauthorized
	}

	if len(authHeader) < len("Bearer") {
		return "", ErrEmptyToken
	}

	jwtToken := authHeader[len("Bearer"):]
	claims, err := jwtApp.Decode(jwtToken)
	if err != nil {
		logger.Get().Errorf("JWT token: %s error %v", authHeader, err.Error())
		return realm, ErrUnauthorized
	}

	err = expected.Validate(claims)
	if err != nil {
		logger.Get().Errorf("Validate token result: %v", err.Error())
		return realm, ErrUnauthorized
	}

	return "", nil
}

// ------------------------------
// DecodeToken extracts user info in JWT token
func DecodeToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if core.GetConfig().Environment == "production" {
			message, err := validateToken(c)
			if err != nil {
				c.Header("WWW-Authenticate", message)
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		}

		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 6 {
			logger.Get().Errorf("Invalid Authorization length", authHeader)
			c.Header("WWW-Authenticate", "token error")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		jwtToken := authHeader[len("Bearer"):]
		claims, err := jwtApp.Decode(jwtToken)
		if err != nil {
			logger.Get().Errorf("Decode jwt token error: ", err)
			c.Header("WWW-Authenticate", "token error")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		cl := *claims

		claimContext, ok := cl["context"]
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		} else {
			mapClaim := claimContext.(map[string]interface{})
			user, ok := mapClaim["user"]
			if ok {
				name := user.(map[string]interface{})
				username, ok := name["name"]
				if ok {
					c.Set("UserName", username)
					c.Next()
					return
				}
			}

			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
