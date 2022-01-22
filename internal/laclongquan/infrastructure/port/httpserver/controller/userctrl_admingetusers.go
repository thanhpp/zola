package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) AdminGetUsers(c *gin.Context) {
	adminID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get userUUID from claims: %v", err)
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser, "invalidate user")
		return
	}

	res, err := ctrl.handler.GetUsers(c, adminID.String())
	if err != nil {
		logger.Errorf("admin get users - %s - error: %v", adminID.String(), err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalidate user")
			return

		case application.ErrPermissionDenied:
			ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser, "invalidate user")
			return
		}
		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	var resp = new(dto.GetUserListResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res, ctrl.formUserMediaUrlFn)

	c.JSON(http.StatusOK, resp)
}
