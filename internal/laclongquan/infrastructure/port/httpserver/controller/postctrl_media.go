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
	// userID, err := getUserUUIDFromClaims(c)
	// if err != nil {
	// 	logger.Errorf("get user id from claims error: %v", err)
	// 	ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
	// 	return
	// }

	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("get post id error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid post id")
		return
	}

	mediaID, err := getMediaID(c)
	if err != nil {
		logger.Errorf("get media id error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid media id")
		return
	}

	var (
		videoThumbFlags bool
	)
	if strings.Contains(mediaID, ThumbPostfix) {
		videoThumbFlags = true
		mediaID = mediaID[:len(mediaID)-len(ThumbPostfix)]
	}
	media, err := ctrl.handler.GetMedia(c, "userID.String()", postID.String(), mediaID)
	if err != nil {
		logger.Errorf("get media %s error: %v", mediaID, err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "user not found")
			return

		case repository.ErrPostNotFound:
			ginAbortInternalError(c, responsevalue.ValuePostNotExist, "post not found")
			return

		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "relation not found")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "locked user")
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.ValueActionHasBeenDone, "locked post")
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidAccess, "permission denied")
			return

		case entity.ErrPostNotContainsMedia:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "post not contains media")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	if videoThumbFlags {
		// logger.Debugf("media %s, thumb %s", mediaID, media.ThumbPath())
		c.File(media.ThumbPath())
		return
	}

	c.File(media.Path())
}
