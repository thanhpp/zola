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
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid post id", nil)
		return
	}

	userID := getUserUUID(c)

	if err := ctrl.handler.DeletePost(c, userID, postID.String()); err != nil {
		logger.Errorf("delete post: %v", err)
		switch err {
		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodePostNotExist, "post not exist", nil)
			return

		case application.ErrPostCannotBeDeleted:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "post cannot be deleted", nil)
			return
		}

		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
}
