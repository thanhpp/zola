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

func (ctrl PostController) EditPost(c *gin.Context) {
	var req = new(dto.EditPostReq)
	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind req %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, responsevalue.MsgInvalidRequest)
		return
	}

	// logger.Debugf("delete id %v", req.MediaDel)

	creator, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user uuid error: %v", err)
		ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
		return
	}

	postUUID, err := getPostID(c)
	if err != nil {
		logger.Errorf("get post uuid error: %v", err)
		ginAbortInternalError(c, responsevalue.ValueInvalidParameterValue, "invalid post id")
		return
	}

	err = ctrl.handler.UpdatePost(
		c,
		creator,
		postUUID,
		req.Described,
		req.MediaDel,
		genMultipartOpts(c)...,
	)
	if err != nil {
		logger.Errorf("update post error: %v", err)
		switch err {
		case application.ErrUnauthorizedCreator:
			ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser, responsevalue.MsgUnauthorized)
			return

		case entity.ErrTooManyImages:
			ginAbortNotAcceptable(c, responsevalue.ValueMaxImagesReached, "too many images")
			return

		case entity.ErrInvalidImageExtension:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid image extension")
			return

		case entity.ErrMediaImageTooBig:
			ginAbortNotAcceptable(c, responsevalue.ValueFileTooBig, "invalid image size")
			return

		case entity.ErrInvalidVideoExtension:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid video extension")
			return

		case entity.ErrMediaVideoTooBig:
			ginAbortNotAcceptable(c, responsevalue.ValueFileTooBig, "invalid video size")
			return

		case entity.ErrInvalidVideoDuration:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid video duration")
			return

		default:
			ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
			return
		}
	}

	resp := new(dto.DefaultResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	c.JSON(http.StatusOK, resp)
}

func (ctrl PostController) EditPerm(c *gin.Context) {
	var req = new(dto.EditPostPermReq)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind req %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, responsevalue.MsgInvalidRequest)
		return
	}

	canCommentPerm, err := dto.BoolTranslateStr(req.CanComment)
	if err != nil {
		logger.Errorf("translate can comment perm error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, responsevalue.MsgInvalidRequest)
		return
	}

	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("get post id error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid post id")
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user uuid error: %v", err)
		ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser, responsevalue.MsgUnauthorized)
		return
	}

	if err := ctrl.handler.EditPerm(c, userID.String(), postID.String(),
		application.PermReq{
			CanComment: canCommentPerm,
		}); err != nil {
		logger.Errorf("edit perm error: %v", err)
		switch err {
		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid post id")
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid user id")
			return

		case entity.ErrNotCreator:
			ginAbortUnauthorized(c, responsevalue.ValueInvalidateUser, responsevalue.MsgUnauthorized)
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "post is locked")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user is locked")
			return
		}
		ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
		return
	}

	ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
}
