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

func (ctrl PostController) CreateComment(c *gin.Context) {
	var (
		req = new(dto.CreateCommentReq)
	)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("invalid request: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid request")
		return
	}

	if req.Comment == "" {
		logger.Error("Nil comment")
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid comment")
		return
	}

	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("invalid post id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid post id")
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("invalid user id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	err = ctrl.handler.CreateComment(c, postID.String(), userID.String(), req.Comment)
	if err != nil {
		logger.Errorf("create comment error: %v", err)
		switch err {
		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValuePostNotExist, "post not exist")
			return

		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "not a friend")
			return

		case entity.ErrNotFriend:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "not a friend")
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.ValueActionHasBeenDone, "locked post")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
			return

		case application.ErrAlreadyBlocked:
			resp := new(dto.CreateCommentResp)
			resp.SetCode(responsevalue.ValueOK.Code)
			resp.SetMsg(responsevalue.MsgOK)
			resp.SetIsBlocked(true)
			c.JSON(http.StatusOK, resp)
			return

		case entity.ErrContentTooLong:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "content too long")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
		return
	}

	res, err := ctrl.handler.GetPostComments(c, userID.String(), postID.String(), 0, 20)
	if err != nil {
		logger.Errorf("get post comments error: %v", err)
		ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
		return
	}

	resp := new(dto.CreateCommentResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res, ctrl.formUserMediaURLFunc)
	ginRespOK(c, responsevalue.ValueOK, resp)
}

func (ctrl PostController) UpdateComment(c *gin.Context) {
	var (
		req = new(dto.UpdateCommentReq)
	)
	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("invalid request: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid request")
		return
	}

	// logger.Debugf("lenght new content: %d", len(req.NewContent))
	if len(req.NewContent) == 0 {
		logger.Errorf("nil content")
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid content")
		return
	}

	updaterID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from claims: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("get post id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid post id")
		return
	}

	commentID, err := getCommentID(c)
	if err != nil {
		logger.Errorf("get comment id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid comment id")
		return
	}

	err = ctrl.handler.UpdateComment(c, updaterID.String(), postID.String(), commentID.String(), req.NewContent)
	if err != nil {
		logger.Errorf("update comment error: %v", err)
		switch err {
		case repository.ErrCommentNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "comment not exist")
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "user not found")
			return

		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValuePostNotExist, "post not found")
			return

		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "not a friend")
			return

		case entity.ErrNotFriend:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "not a friend")
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "permission denied")
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.ValueActionHasBeenDone, "locked post")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "locked user")
			return

		case entity.ErrInvalidCommentContent:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid comment content")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
		return
	}

	ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
}

func (ctrl PostController) DeleteComment(c *gin.Context) {
	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from claims: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("get post id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid post id")
		return
	}

	commentID, err := getCommentID(c)
	if err != nil {
		logger.Errorf("get comment id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid comment id")
		return
	}

	if err := ctrl.handler.DeleteComment(c, userID.String(), postID.String(), commentID.String()); err != nil {
		logger.Errorf("delete comment error: %v", err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
			return

		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValuePostNotExist, "invalid post id")
			return

		case repository.ErrCommentNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid comment id")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "locked user")
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.ValuePostNotExist, "locked post")
			return

		case entity.ErrNotCreator:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid permission")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
		return
	}

	ginRespOK(c, responsevalue.ValueOK, nil)
}
