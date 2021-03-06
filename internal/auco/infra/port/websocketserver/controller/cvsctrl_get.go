package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/auco/app"
	acdto "github.com/thanhpp/zola/internal/auco/infra/port/websocketserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl ConversationController) GetByPartnerID(c *gin.Context) {
	requestorID, err := getRequestorIDFromClaims(c)
	if err != nil {
		logger.Errorf("CvsCtrl - get claims %v", err)
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser.Code, "invalid user", nil)
		return
	}
	partnerID := c.Param("id")

	offset, limit := pagination(c)
	res, err := ctrl.conversationHandler.GetByPartnerID(c, requestorID, partnerID, offset, limit)
	if err != nil {
		switch err {
		case app.ErrBlocked:
			resp := new(acdto.GetConversationResp)
			resp.SetCode(responsevalue.ValueOK.Code)
			resp.SetMsg(responsevalue.MsgOK)
			resp.SetIsBlocked()
			c.JSON(http.StatusOK, resp)
			return
		}
		ginAbortInternalError(c, responsevalue.ValueUnknownError.Code, responsevalue.MsgUnknownError, err.Error())
		return
	}

	resp := new(acdto.GetConversationResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res)

	c.JSON(http.StatusOK, resp)
}

func (ctrl ConversationController) GetByRoomID(c *gin.Context) {
	requestorID, err := getRequestorIDFromClaims(c)
	if err != nil {
		logger.Errorf("CvsCtrl - get claims %v", err)
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser.Code, "invalid user", nil)
		return
	}
	roomID := c.Param("id")

	offset, limit := pagination(c)
	res, err := ctrl.conversationHandler.GetByRoomID(c, requestorID, roomID, offset, limit)
	if err != nil {
		switch err {
		case app.ErrBlocked:
			resp := new(acdto.GetConversationResp)
			resp.SetCode(responsevalue.ValueOK.Code)
			resp.SetMsg(responsevalue.MsgOK)
			resp.SetIsBlocked()
			c.JSON(http.StatusOK, resp)
			return
		}
		ginAbortInternalError(c, responsevalue.ValueUnknownError.Code, responsevalue.MsgUnknownError, err.Error())
		return
	}

	resp := new(acdto.GetConversationResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res)

	c.JSON(http.StatusOK, resp)
}
