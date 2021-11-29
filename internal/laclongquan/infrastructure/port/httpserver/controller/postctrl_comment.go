package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl PostController) CreateComment(c *gin.Context) {
	var (
		req = new(dto.CreateCommentReq)
	)

	if err := c.ShouldBind(req); err != nil {
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid request", req)
		return
	}

	if req.Comment == "" {
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid comment", req)
		return
	}

	postID, err := getPostID(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid post id", nil)
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	err = ctrl.handler.CreateComment(c, postID.String(), userID.String(), req.Comment)
	if err != nil {
		switch err {
		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodePostNotExist, "post not exist", nil)
			return

		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "not a friend", nil)
			return

		case application.ErrNotFriend:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "not a friend", nil)
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.CodeActionHasBeenDone, "locked post", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case application.ErrAlreadyBlocked:
			resp := new(dto.CreateCommentResp)
			resp.SetCode(responsevalue.CodeOK)
			resp.SetMsg(responsevalue.MsgOK)
			resp.SetIsBlocked(true)
			c.JSON(http.StatusOK, resp)
			return

		case entity.ErrContentTooLong:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "content too long", nil)
			return
		}

		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	// FIXME: get comments by index and count

	// FIXME: temporary respone
	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
}
