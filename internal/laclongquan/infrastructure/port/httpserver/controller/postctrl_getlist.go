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

func (ctrl PostController) GetListPost(c *gin.Context) {
	var (
		req = new(dto.GetListPostReq)
	)
	if err := c.ShouldBind(req); err != nil {
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid request")
		return
	}

	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser, "invalid token")
		return
	}

	offset, limit := pagination(c)

	res, err := ctrl.handler.GetListPost(c, requestorID.String(), req.LastID, offset, limit)
	if err != nil {
		logger.Errorf("GetListPost by user %s error: %v", requestorID, err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
			return

		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid post id")
			return

		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidAccess, "invalid access")
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidAccess, "invalid access")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "locked user")
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "locked post")
			return
		}
		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	resp := new(dto.GetListPostResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res, req.LastID, ctrl.formMediaURLFunc, ctrl.formVideoThumbURLFunc, ctrl.formUserMediaURLFunc)

	c.JSON(http.StatusOK, resp)
}
