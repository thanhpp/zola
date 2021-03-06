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

func (ctrl UserController) AdminCreateUser(c *gin.Context) {
	var req = new(dto.AdminCreateUserReq)
	if err := c.ShouldBind(req); err != nil {
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid request")
		return
	}

	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser, "invalidate user")
		return
	}

	newUser, err := ctrl.handler.AdminCreateUser(c,
		requestorID.String(),
		req.Phone, req.Pass,
		req.Name, req.Username, req.Description,
		req.Address, req.City, req.Country)
	if err != nil {
		logger.Errorf("admin %s create user error: %v", requestorID, err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalidate user")
			return

		case entity.ErrPermissionDenied:
			ginAbortUnauthorized(c, responsevalue.ValueInvalidAccess, "invalid access")
			return

		case repository.ErrDuplicateUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "duplicate user")
			return

		case entity.ErrInvalidCountry:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid country")
			return

		case entity.ErrInvalidPhone:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid phone number")
			return

		case entity.ErrInvalidPassword:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid password")
			return
		}
		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	resp := new(dto.AdminCreateUserResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(newUser.ID().String())

	c.JSON(http.StatusOK, resp)
}
