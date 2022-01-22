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
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, responsevalue.MsgInvalidRequest)
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	if err := ctrl.handler.ChangePassword(c, userID.String(), req.Password, req.NewPassword); err != nil {
		switch err {
		case repository.ErrUserNotFound:
			ginAbortInternalError(c, responsevalue.ValueInvalidateUser, "invalid user")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "user is locked")
			return

		case entity.ErrPassNotEqual:
			ginAbortInternalError(c, responsevalue.ValueInvalidParameterValue, "invalid old password")
			return

		case entity.ErrSameOldPass:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "old password is same with new password")
			return

		case entity.ErrCommonPass:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "new password is too common with old password")
			return

		case entity.ErrInvalidPassword:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid new password")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
}
