package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) ChangePassword(c *gin.Context) {
	var (
		req = new(dto.ChangePasswordReq)
	)

	if err := c.ShouldBind(req); err != nil {
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, responsevalue.MsgInvalidRequest, req)
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	if err := ctrl.handler.ChangePassword(c, userID.String(), req.Password, req.NewPassword); err != nil {
		switch err {
		case repository.ErrUserNotFound:
			ginAbortInternalError(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "user is locked", nil)
			return

		case entity.ErrPassNotEqual:
			ginAbortInternalError(c, responsevalue.CodeInvalidParameterValue, "invalid old password", nil)
			return

		case entity.ErrSameOldPass:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "old password is same with new password", nil)
			return

		case entity.ErrCommonPass:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "new password is too common with old password", nil)
			return

		case entity.ErrInvalidPassword:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid new password", nil)
			return
		}

		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
}
