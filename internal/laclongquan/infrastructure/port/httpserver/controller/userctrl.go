package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

type UserController struct {
	handler application.UserHandler
}

func (ctrl UserController) SignUp(c *gin.Context) {
	var (
		req = new(dto.SignUpReq)
	)

	if err := c.ShouldBindJSON(req); err != nil {
		logger.Errorf("bind req %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid request", nil)
		return
	}

	err := ctrl.handler.SignUp(c, req.PhoneNumber, req.Password, "", "")
	if err != nil {
		logger.Errorf("handle %v", err)
		switch err {
		case entity.ErrInvalidPhone:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid phone number", nil)
			return

		case entity.ErrInvalidPassword:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid password", nil)
			return

			// case ErrUserExisted:
			// ginAbortNotAcceptable(c, responsevalue.CodeUserExisted, "user existed", nil)
			// return
		}

		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
}
