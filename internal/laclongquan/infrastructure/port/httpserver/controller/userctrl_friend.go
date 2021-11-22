package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) NewFriendRequest(c *gin.Context) {
	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from ctx error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", err)
		return
	}

	requesteeID, err := getUserUUIDFromParam(c)
	if err != nil {
		logger.Errorf("get user id from param error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", err)
		return
	}

	err = ctrl.handler.NewFriendRequest(c, requestorID.String(), requesteeID.String())
	if err != nil {
		logger.Errorf("new friend request error: %v", err)
		switch err {
		case application.ErrRelationExisted:
			ginAbortNotAcceptable(c, responsevalue.CodeActionHasBeenDone, "relation existed", nil)
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user not exist", nil)
			return

		case entity.ErrSameUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "same id request", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "locked user", nil)
			return
		}
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
}
