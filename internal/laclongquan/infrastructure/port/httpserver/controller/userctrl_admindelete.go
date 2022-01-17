package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) AdminDeleteUser(c *gin.Context) {
	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		ginAbortUnauthorized(c, responsevalue.CodeInvalidateUser, "invalidate user", nil)
		return
	}

	requestedID, err := getUserUUIDFromParam(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid user ID", nil)
		return
	}

	if err := ctrl.handler.AdminDeleteUser(c, requestorID.String(), requestedID.String()); err != nil {
		logger.Errorf("admin %s delete user error: %v", requestorID, err)
		switch err {
		case application.ErrSelfDelete:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "self delete", nil)
			return

		case entity.ErrPermissionDenied:
			ginAbortUnauthorized(c, responsevalue.CodeInvalidAccess, "invalid access", nil)
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid user ID", nil)
			return

		default:
			ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, err.Error())
			return
		}
	}

	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
}
