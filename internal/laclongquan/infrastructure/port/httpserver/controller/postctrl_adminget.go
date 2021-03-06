package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl PostController) AdminGetListPosts(c *gin.Context) {
	var req = new(dto.GetListPostReq)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("PostCtrl: AdminGetListPosts: invalid request: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid request")
		return
	}

	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("PostCtrl: AdminGetListPosts: invalid token: %v", err)
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser, "invalid token")
		return
	}

	offset, limit := pagination(c)

	res, err := ctrl.handler.AdminGetList(c, requestorID.String(), req.LastID, offset, limit)
	if err != nil {
		logger.Errorf("PostCtrl: admin get list error :%v", err)
		switch err {

		default:
			ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		}
	}

	resp := new(dto.GetListPostResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res, req.LastID, ctrl.formMediaURLFunc, ctrl.formVideoThumbURLFunc, ctrl.formUserMediaURLFunc)

	c.JSON(http.StatusOK, resp)
}
