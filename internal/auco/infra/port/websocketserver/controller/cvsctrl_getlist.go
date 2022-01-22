package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	acdto "github.com/thanhpp/zola/internal/auco/infra/port/websocketserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl ConversationController) GetList(c *gin.Context) {
	requestorID, err := getRequestorIDFromClaims(c)
	if err != nil {
		logger.Errorf("CvsCtrl - get claims %v", err)
		ginAbortUnauthorized(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}
	offset, limit := pagination(c)

	data, err := ctrl.conversationHandler.GetListConversation(c, requestorID, offset, limit)
	if err != nil {
		logger.Errorf("ConversationCtrl: get list conversation error: %v", err)
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, err.Error())
		return
	}

	resp := new(acdto.GetListConversationResp)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(data, requestorID)

	c.JSON(http.StatusOK, resp)
}
