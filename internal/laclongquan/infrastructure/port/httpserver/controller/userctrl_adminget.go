package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) AdminGetByID(c *gin.Context) {
	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		ginAbortUnauthorized(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	requestedID, err := getUserUUIDFromParam(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid user id", nil)
		return
	}

	res, err := ctrl.handler.AdminGetUser(c, requestorID.String(), requestedID.String())
	if err != nil {
		logger.Errorf("admin %s failed to get user %s err: %v", requestorID, requestedID, err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid user id", nil)
			return

		case entity.ErrPermissionDenied:
			ginAbortUnauthorized(c, responsevalue.CodeInvalidAccess, "invalid access", nil)
			return
		}
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, err.Error())
		return
	}

	resp := new(dto.GetUserResp)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res.User, res.FriendCount, res.IsFriend, ctrl.formUserMediaUrlFn)

	c.JSON(http.StatusOK, resp)
}
