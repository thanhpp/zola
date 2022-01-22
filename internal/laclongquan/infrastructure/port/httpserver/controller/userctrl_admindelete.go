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
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser, "invalidate user")
		return
	}

	requestedID, err := getUserUUIDFromParam(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid user ID")
		return
	}

	if err := ctrl.handler.AdminDeleteUser(c, requestorID.String(), requestedID.String()); err != nil {
		logger.Errorf("admin %s delete user error: %v", requestorID, err)
		switch err {
		case application.ErrSelfDelete:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "self delete")
			return

		case entity.ErrPermissionDenied:
			ginAbortUnauthorized(c, responsevalue.ValueInvalidAccess, "invalid access")
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid user ID")
			return

		default:
			ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
			return
		}
	}

	ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
}
