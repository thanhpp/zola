package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (u UserController) InternalGetUser(c *gin.Context) {
	userID, err := getUserUUIDFromParam(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, err.Error())
		return
	}

	user, err := u.handler.InternalGetUser(c, userID.String())
	if err != nil {
		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	resp := new(dto.InternalGetUserResp)
	resp.SetData(user, u.formUserMediaUrlFn)
	c.JSON(http.StatusOK, resp)
}

func (u UserController) InternalIsBlock(c *gin.Context) {
	userAID := c.Query("usera")

	userBID := c.Query("userb")

	isBlock, err := u.handler.InternalIsBlock(c, userAID, userBID)
	if err != nil {
		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	resp := new(dto.InternalIsBlockResp)
	resp.SetData(isBlock)
	c.JSON(http.StatusOK, resp)
}
