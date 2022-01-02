package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) GetUserInfo(c *gin.Context) {
	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		ginAbortUnauthorized(c, responsevalue.CodeInvalidToken, "invalid token", nil)
		return
	}

	requestedID, err := getUserUUIDFromParam(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid user ID", nil)
	}

	res, err := ctrl.handler.GetUserByID(c, requestorID.String(), requestedID.String())
	if err != nil {
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user not found", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user not found", nil)
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user not found", nil)
			return
		}

		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	var resp = new(dto.GetUserResp)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res.User, res.FriendCount, res.IsFriend, ctrl.formUserMediaUrlFn)

	c.JSON(http.StatusOK, resp)
}
