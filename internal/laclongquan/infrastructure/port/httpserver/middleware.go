package httpserver

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (s HTTPServer) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := s.getBearerToken(c)
		if err != nil {
			logger.Errorf("get token %v", err)
			resp := new(dto.DefaultResp)
			resp.SetCode(responsevalue.ValueInvalidToken.Code)
			resp.SetMsg(responsevalue.MsgUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		claims, err := s.auth.NewClaimsFromToken(c, token)
		if err != nil {
			logger.Debugf("error token %s", token)
			logger.Errorf("get claims %v", err)
			resp := new(dto.DefaultResp)
			resp.SetCode(responsevalue.ValueInvalidToken.Code)
			resp.SetMsg(responsevalue.MsgUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		c.Set("claims", *claims)

		c.Next()
	}
}

const (
	authorization = "Authorization"
	bearer        = "Bearer "
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

func (s HTTPServer) getBearerToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader(authorization)
	if len(authHeader) > len(bearer) && authHeader[:len(bearer)] == bearer {
		return authHeader[len(bearer):], nil
	}

	return "", ErrInvalidToken
}

func (s HTTPServer) validateInternal(c *gin.Context) {
	token, err := s.getBearerToken(c)
	if err != nil {
		logger.Errorf("get token %v", err)
		resp := new(dto.DefaultResp)
		resp.SetCode(responsevalue.ValueInvalidToken.Code)
		resp.SetMsg(responsevalue.MsgUnauthorized)
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	claims, err := s.auth.NewClaimsFromToken(c, token)
	if err != nil {
		logger.Debugf("error token %s", token)
		logger.Errorf("get claims %v", err)
		resp := new(dto.DefaultResp)
		resp.SetCode(responsevalue.ValueInvalidToken.Code)
		resp.SetMsg(responsevalue.MsgUnauthorized)
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	c.JSON(http.StatusOK, claims)
}
