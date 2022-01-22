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
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid post id")
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("Error while getting userID: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	result, err := ctrl.handler.GetPost(c, userID.String(), postID.String())
	if err != nil {
		logger.Errorf("Error while getting post: %v", err)
		switch err {
		case application.ErrAlreadyBlocked:
			resp := new(dto.GetPostResponse)
			resp.SetBlockedResponse()
			resp.SetCode(responsevalue.ValueOK.Code)
			resp.SetMsg(responsevalue.MsgOK)
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
			return

		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValuePostNotExist, "invalid post id")
			return

		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidAccess, "invalid access")
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.ValuePostNotExist, "invalid access")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "locked user")
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.ValuePostNotExist, "locked post")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	resp := new(dto.GetPostResponse)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(result, ctrl.formMediaURLFunc, ctrl.formVideoThumbURLFunc, ctrl.formUserMediaURLFunc)

	c.JSON(http.StatusOK, resp)
}
