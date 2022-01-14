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
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid request", req)
		return
	}

	if req.Comment == "" {
		logger.Error("Nil comment")
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid comment", req)
		return
	}

	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("invalid post id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid post id", nil)
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("invalid user id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	err = ctrl.handler.CreateComment(c, postID.String(), userID.String(), req.Comment)
	if err != nil {
		logger.Errorf("create comment error: %v", err)
		switch err {
		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodePostNotExist, "post not exist", nil)
			return

		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "not a friend", nil)
			return

		case entity.ErrNotFriend:
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

	res, err := ctrl.handler.GetPostComments(c, userID.String(), postID.String(), 0, 20)
	if err != nil {
		logger.Errorf("get post comments error: %v", err)
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	resp := new(dto.CreateCommentResp)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res, ctrl.formUserMediaURLFunc)
	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, resp)
}

func (ctrl PostController) UpdateComment(c *gin.Context) {
	var (
		req = new(dto.UpdateCommentReq)
	)
	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("invalid request: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid request", req)
		return
	}

	// logger.Debugf("lenght new content: %d", len(req.NewContent))
	if len(req.NewContent) == 0 {
		logger.Errorf("nil content")
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid content", req)
		return
	}

	updaterID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from claims: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("get post id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid post id", nil)
		return
	}

	commentID, err := getCommentID(c)
	if err != nil {
		logger.Errorf("get comment id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid comment id", nil)
		return
	}

	err = ctrl.handler.UpdateComment(c, updaterID.String(), postID.String(), commentID.String(), req.NewContent)
	if err != nil {
		logger.Errorf("update comment error: %v", err)
		switch err {
		case repository.ErrCommentNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "comment not exist", nil)
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "user not found", nil)
			return

		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodePostNotExist, "post not found", nil)
			return

		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "not a friend", nil)
			return

		case entity.ErrNotFriend:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "not a friend", nil)
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "permission denied", nil)
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.CodeActionHasBeenDone, "locked post", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "locked user", nil)
			return

		case entity.ErrInvalidCommentContent:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid comment content", nil)
			return
		}

		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
}

func (ctrl PostController) DeleteComment(c *gin.Context) {
	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from claims: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("get post id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid post id", nil)
		return
	}

	commentID, err := getCommentID(c)
	if err != nil {
		logger.Errorf("get comment id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid comment id", nil)
		return
	}

	if err := ctrl.handler.DeleteComment(c, userID.String(), postID.String(), commentID.String()); err != nil {
		logger.Errorf("delete comment error: %v", err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodePostNotExist, "invalid post id", nil)
			return

		case repository.ErrCommentNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid comment id", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "locked user", nil)
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.CodePostNotExist, "locked post", nil)
			return

		case entity.ErrNotCreator:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid permission", nil)
			return
		}

		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
}
