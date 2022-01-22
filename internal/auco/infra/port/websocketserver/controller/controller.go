package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/auco/infra/adapter/llqclient"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
)

const (
	paginationMinOffset    = 0
	paginationMinLimit     = 1
	paginationDefaultLimit = 20
	paginationMaxLimit     = 100
)

func pagination(c *gin.Context) (offset, limit int) {
	indexStr := c.Query("index")
	if len(indexStr) == 0 {
		return paginationMinOffset, paginationDefaultLimit
	}

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return paginationMinOffset, paginationDefaultLimit
	}

	if index < 1 {
		return paginationMinOffset, paginationDefaultLimit
	}

	limitStr := c.Query("count")
	if len(limitStr) == 0 {
		return (index - 1) * paginationDefaultLimit, paginationDefaultLimit
	}

	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return (index - 1) * paginationDefaultLimit, paginationDefaultLimit
	}

	if limit < paginationMinLimit || limit > paginationMaxLimit {
		return (index - 1) * paginationDefaultLimit, paginationDefaultLimit
	}

	return (index - 1) * limit, limit
}

func ginRespOK(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(
		http.StatusOK,
		dto.NewDefaultResp(code, msg, data),
	)
}

func ginAbortNotAcceptable(c *gin.Context, code int, msg string, data interface{}) {
	c.AbortWithStatusJSON(
		http.StatusNotAcceptable,
		dto.NewDefaultResp(code, msg, data),
	)
}

func ginAbortInternalError(c *gin.Context, code int, msg string, data interface{}) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		dto.NewDefaultResp(code, msg, data),
	)
}

func ginAbortUnauthorized(c *gin.Context, code int, msg string, data interface{}) {
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		dto.NewDefaultResp(code, msg, data),
	)
}

var (
	claimsKey = "claims"
)

var (
	ErrClaimsNotExist = errors.New("claims not exist")
	ErrNotClaims      = errors.New("not claims")
)

func getClaimsFromCtx(c *gin.Context) (*llqclient.ValidateTokenResp, error) {
	claimsItf, ok := c.Get(claimsKey)
	if !ok {
		return nil, ErrClaimsNotExist
	}

	claims, ok := claimsItf.(llqclient.ValidateTokenResp)
	if !ok {
		return nil, ErrNotClaims
	}

	return &claims, nil
}

func getRequestorIDFromClaims(c *gin.Context) (string, error) {
	claims, err := getClaimsFromCtx(c)
	if err != nil {
		return "", err
	}
	return claims.User.ID, nil
}
