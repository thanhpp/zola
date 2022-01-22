package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/auco/infra/port/websocketserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl ConversationController) DeleteByConversationID(c *gin.Context) {
	requestorID, err := getRequestorIDFromClaims(c)
	if err != nil {
		logger.Errorf("CvsCtrl - get claims %v", err)
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser.Code, "invalid user", nil)
		return
	}

	conversationID := c.Param("id")

	err = ctrl.conversationHandler.DeleteByConversationID(c, requestorID, conversationID)
	if err != nil {
		logger.Errorf("CvsCtrl - delete conversation %v", err)
		ginAbortInternalError(c, responsevalue.ValueUnknownError.Code, responsevalue.ValueOK.Message, err.Error())
		return
	}

	resp := new(dto.DefaultResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.ValueOK.Message)

	c.JSON(http.StatusOK, resp)
}

func (ctrl ConversationController) DeleteMessage(c *gin.Context) {
	requestorID, err := getRequestorIDFromClaims(c)
	if err != nil {
		logger.Errorf("CvsCtrl - get claims %v", err)
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser.Code, "invalid user", nil)
		return
	}

	messageID := c.Param("id")

	if err := ctrl.conversationHandler.DeleteMessage(c, requestorID, messageID); err != nil {
		logger.Errorf("CvsCtrl - delete message %v", err)
		ginAbortInternalError(c, responsevalue.ValueUnknownError.Code, responsevalue.ValueOK.Message, err.Error())
		return
	}

	resp := new(dto.DefaultResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.ValueOK.Message)
	c.JSON(http.StatusOK, resp)
}
