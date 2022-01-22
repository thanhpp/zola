package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl PostController) DeletePost(c *gin.Context) {
	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("Error get post id: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid post id")
		return
	}

	userID := getUserUUID(c)

	if err := ctrl.handler.DeletePost(c, userID, postID.String()); err != nil {
		logger.Errorf("delete post: %v", err)
		switch err {
		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValuePostNotExist, "post not exist")
			return

		case application.ErrPostCannotBeDeleted:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "post cannot be deleted")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	ginRespOK(c, responsevalue.ValueOK, nil)
}
