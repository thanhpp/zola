package controller

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func getUserUUIDFromCtx(c *gin.Context) (uuid.UUID, error) {
	claims, err := getClaimsFromCtx(c)
	if err != nil {
		return uuid.Nil, err
	}

	userUUID, err := uuid.Parse(claims.User.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return userUUID, nil
}

func getUserUUID(c *gin.Context) string {
	claims, err := getClaimsFromCtx(c)
	if err != nil {
		return ""
	}

	return claims.User.ID
}

const source = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func genRandomString(length int) string {
	seedRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var strB = new(strings.Builder)
	strB.Grow(length)
	for i := 0; i < length; i++ {
		strB.WriteByte(source[seedRand.Intn(len(source))])
	}

	return strB.String()
}
