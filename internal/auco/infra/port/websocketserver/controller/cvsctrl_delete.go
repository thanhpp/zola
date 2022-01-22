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
		ginAbortUnauthorized(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	conversationID := c.Param("id")

	err = ctrl.conversationHandler.DeleteByConversationID(c, requestorID, conversationID)
	if err != nil {
		logger.Errorf("CvsCtrl - delete conversation %v", err)
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, err.Error())
		return
	}

	resp := new(dto.DefaultResp)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetMsg(responsevalue.MsgOK)

	c.JSON(http.StatusOK, resp)
}
