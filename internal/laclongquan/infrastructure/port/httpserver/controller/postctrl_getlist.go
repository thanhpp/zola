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
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid request", req)
		return
	}

	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		ginAbortUnauthorized(c, responsevalue.CodeInvalidateUser, "invalid token", err)
		return
	}

	offset, limit := pagination(c)

	res, err := ctrl.handler.GetListPost(c, requestorID.String(), req.LastID, offset, limit)
	if err != nil {
		logger.Errorf("GetListPost by user %s error: %v", requestorID, err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid post id", nil)
			return

		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidAccess, "invalid access", nil)
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidAccess, "invalid access", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "locked user", nil)
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "locked post", nil)
			return
		}
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	resp := new(dto.GetListPostResp)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res, req.LastID, ctrl.formMediaURLFunc, ctrl.formVideoThumbURLFunc, ctrl.formUserMediaURLFunc)

	c.JSON(http.StatusOK, resp)
}
