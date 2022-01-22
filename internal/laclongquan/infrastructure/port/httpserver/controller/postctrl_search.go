package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl PostController) Search(c *gin.Context) {
	var req = new(dto.SearchReq)
	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("Error binding search request: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, err.Error())
		return
	}

	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("Error getting user id from claims: %v", err)
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser, err.Error())
		return
	}

	offset, limit := pagination(c)

	res, err := ctrl.handler.Search(c, requestorID.String(), req.Keyword, offset, limit)
	if err != nil {
		logger.Errorf("Error searching posts: %v", err)
		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	resp := new(dto.SearchResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.ValueOK.Message)
	resp.SetData(res, ctrl.formMediaURLFunc, ctrl.formVideoThumbURLFunc, ctrl.formUserMediaURLFunc)

	c.JSON(http.StatusOK, resp)
}
