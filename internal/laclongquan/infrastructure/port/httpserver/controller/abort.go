package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
)

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
