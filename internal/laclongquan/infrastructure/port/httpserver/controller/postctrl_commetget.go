package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl PostController) GetComments(c *gin.Context) {
	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("invalid user id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("invalid post id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid post id", nil)
		return
	}

	offset, limit := pagination(c)
	logger.Debugf("offset %d, limit %d", offset, limit)

	res, err := ctrl.handler.GetPostComments(c, requestorID.String(), postID.String(), offset, limit)
	if err != nil {
		logger.Errorf("get post %s comments error: %v", postID, err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodePostNotExist, "invalid post id", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.CodeActionHasBeenDone, "invalid post id", nil)
			return

		case application.ErrAlreadyBlocked:
			resp := new(dto.GetCommentResp)
			resp.SetCode(responsevalue.CodeOK)
			resp.SetMsg(responsevalue.MsgOK)
			resp.SetIsBlocked()
			c.JSON(http.StatusOK, resp)
			return
		}
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	resp := new(dto.GetCommentResp)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res, ctrl.formUserMediaURLFunc)

	c.JSON(http.StatusOK, resp)
}
