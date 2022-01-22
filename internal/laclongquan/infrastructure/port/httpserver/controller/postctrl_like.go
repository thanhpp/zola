package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl PostController) LikePost(c *gin.Context) {
	postID, err := getPostID(c)
	if err != nil {
		logger.Errorf("get post id error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid postID")
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid userID")
		return
	}

	count, err := ctrl.likeHandler.LikePost(c, postID.String(), userID.String())
	if err != nil {
		switch err {
		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValuePostNotExist, "post not found")
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.ValueActionHasBeenDone, "locked post")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	resp := new(dto.LikePostResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetData(count)

	c.JSON(http.StatusOK, resp)
}
