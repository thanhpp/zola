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

func (ctrl PostController) GetPost(c *gin.Context) {
	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("Error while getting postID: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid post id", nil)
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("Error while getting userID: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	result, err := ctrl.handler.GetPost(c, userID.String(), postID.String())
	if err != nil {
		logger.Errorf("Error while getting post: %v", err)
		switch err {
		case application.ErrAlreadyBlocked:
			resp := new(dto.GetPostResponse)
			resp.SetBlockedResponse()
			resp.SetCode(responsevalue.CodeOK)
			resp.SetMsg(responsevalue.MsgOK)
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

	resp := new(dto.GetPostResponse)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(result, ctrl.formMediaURLFunc)

	c.JSON(http.StatusOK, resp)
}
