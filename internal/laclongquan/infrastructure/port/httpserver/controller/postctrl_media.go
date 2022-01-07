package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

const (
	ThumbPostfix = "-thumb"
)

func (ctrl PostController) GetMedia(c *gin.Context) {
	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from claims error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("get post id error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid post id", nil)
		return
	}

	mediaID, err := getMediaID(c)
	if err != nil {
		logger.Errorf("get media id error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, "invalid media id", nil)
		return
	}

	var (
		videoThumbFlags bool
	)
	if strings.Contains(mediaID, ThumbPostfix) {
		videoThumbFlags = true
		mediaID = mediaID[:len(mediaID)-len(ThumbPostfix)]
	}
	media, err := ctrl.handler.GetMedia(c, userID.String(), postID.String(), mediaID)
	if err != nil {
		logger.Errorf("get media %s error: %v", mediaID, err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "user not found", nil)
			return

		case repository.ErrPostNotFound:
			ginAbortInternalError(c, responsevalue.CodePostNotExist, "post not found", nil)
			return

		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "relation not found", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "locked user", nil)
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.CodeActionHasBeenDone, "locked post", nil)
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidAccess, "permission denied", nil)
			return

		case entity.ErrPostNotContainsMedia:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "post not contains media", nil)
			return
		}

		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	if videoThumbFlags {
		// logger.Debugf("media %s, thumb %s", mediaID, media.ThumbPath())
		c.File(media.ThumbPath())
		return
	}

	c.File(media.Path())
}
