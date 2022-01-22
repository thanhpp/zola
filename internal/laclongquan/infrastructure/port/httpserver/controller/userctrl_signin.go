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

func (ctrl UserController) SignIn(c *gin.Context) {
	var (
		req = new(dto.SignInReq)
	)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind req %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, responsevalue.MsgInvalidRequest)
		return
	}

	user, err := ctrl.handler.GetUser(c, req.PhoneNumber, req.Password)
	if err != nil {
		logger.Errorf("handle %v", err)

		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "user is not validated")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "user is locked")
			return

		case entity.ErrPassNotEqual:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "password mismatch")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	token, err := ctrl.authsrv.NewTokenFromUser(c, user)
	if err != nil {
		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	resp := new(dto.SignInResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetData(token)

	c.JSON(http.StatusOK, resp)
}
