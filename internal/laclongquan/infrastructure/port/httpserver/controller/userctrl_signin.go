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
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
		return
	}

	user, err := ctrl.handler.GetUser(c, req.PhoneNumber, req.Password)
	if err != nil {
		logger.Errorf("handle %v", err)

		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "user is not validated", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "user is locked", nil)
			return
		}

		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	token, err := ctrl.authsrv.NewTokenFromUser(c, user)
	if err != nil {
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	resp := new(dto.SignInResp)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetData(token)

	c.JSON(http.StatusOK, resp)
}
