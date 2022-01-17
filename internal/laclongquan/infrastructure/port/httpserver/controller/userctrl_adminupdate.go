package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) AdminSetState(c *gin.Context) {
	var req = new(dto.AdminUpdateStateReq)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("Failed to bind request: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid request", nil)
		return
	}

	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("Failed to get user ID from claims: %v", err)
		ginAbortUnauthorized(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	requestedID, err := getUserUUIDFromParam(c)
	if err != nil {
		logger.Errorf("Failed to get user ID from param: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid user ID", nil)
		return
	}

	err = ctrl.handler.AdminUpdateState(c.Request.Context(), requestorID.String(), requestedID.String(), req.State)
	if err != nil {
		logger.Errorf("admin %s update user %s state to %s failed: %v", requestorID, requestedID, req.State, err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid user ID", nil)
			return

		case entity.ErrPermissionDenied:
			ginAbortUnauthorized(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case entity.ErrInvalidState:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid state", nil)
			return

		default:
			ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, err.Error())
			return
		}
	}

	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
}

func (ctrl UserController) AdminUpdatePassword(c *gin.Context) {
	var req = new(dto.AdminUpdatePasswordReq)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("Failed to bind request: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid request", nil)
		return
	}

	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("Failed to get user ID from claims: %v", err)
		ginAbortUnauthorized(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	requestedID, err := getUserUUIDFromParam(c)
	if err != nil {
		logger.Errorf("Failed to get user ID from param: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid user ID", nil)
		return
	}

	err = ctrl.handler.AdminUpdatePass(c, requestorID.String(), requestedID.String(), req.Password)
	if err != nil {
		logger.Errorf("admin %s update user %s password failed: %v", requestorID, requestedID, err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid user ID", nil)
			return

		case entity.ErrPermissionDenied:
			ginAbortUnauthorized(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case entity.ErrInvalidPassword:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid password", nil)
			return

		default:
			ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, err.Error())
			return
		}
	}

	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
}
