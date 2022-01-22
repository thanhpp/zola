package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func ginRespOK(c *gin.Context, value responsevalue.ResponseValue, data interface{}) {
	c.JSON(
		http.StatusOK,
		dto.NewDefaultResp(value.Code, value.Message, data),
	)
}

func ginAbortNotAcceptable(c *gin.Context, value responsevalue.ResponseValue, data interface{}) {
	c.AbortWithStatusJSON(
		http.StatusNotAcceptable,
		dto.NewDefaultResp(value.Code, value.Message, data),
	)
}

func ginAbortInternalError(c *gin.Context, value responsevalue.ResponseValue, data interface{}) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		dto.NewDefaultResp(value.Code, value.Message, data),
	)
}

func ginAbortUnauthorized(c *gin.Context, value responsevalue.ResponseValue, data interface{}) {
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		dto.NewDefaultResp(value.Code, value.Message, data),
	)
}
