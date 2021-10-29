package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
)

var (
	claimsKey = "claims"
)

var (
	ErrClaimsNotExist = errors.New("claims not exist")
	ErrNotClaims      = errors.New("not claims")
)

func getClaimsFromCtx(c *gin.Context) (*auth.Claims, error) {
	claimsItf, ok := c.Get(claimsKey)
	if !ok {
		return nil, ErrClaimsNotExist
	}

	claims, ok := claimsItf.(auth.Claims)
	if !ok {
		return nil, ErrNotClaims
	}

	return &claims, nil
}
