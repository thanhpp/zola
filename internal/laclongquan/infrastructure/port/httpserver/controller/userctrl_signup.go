package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) SignUp(c *gin.Context) {
	var (
		req = new(dto.SignUpReq)
	)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind req %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid request")
		return
	}
	// logger.Debugf("signup req: %v", req)

	err := ctrl.handler.CreateUser(c, req.PhoneNumber, req.Password, "", "")
	if err != nil {
		logger.Errorf("handle %v", err)
		switch err {
		case entity.ErrInvalidPhone:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid phone number")
			return

		case entity.ErrInvalidPassword:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid password")
			return

		case repository.ErrDuplicateUser:
			ginAbortNotAcceptable(c, responsevalue.ValueUserExisted, "user existed")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
		return
	}

	ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
}
