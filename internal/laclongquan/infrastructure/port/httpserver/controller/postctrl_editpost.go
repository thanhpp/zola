package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl PostController) EditPost(c *gin.Context) {
	var req = new(dto.EditPostReq)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind req %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
		return
	}

	creator, err := getUserUUIDFromCtx(c)
	if err != nil {
		logger.Errorf("get user uuid error: %v", err)
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	postUUID, err := getPostID(c)
	if err != nil {
		logger.Errorf("get post uuid error: %v", err)
		ginAbortInternalError(c, responsevalue.CodeInvalidParameterValue, "invalid post id", nil)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		logger.Errorf("multipart form %v", err)
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}
	images := form.File["image"]
	video, _ := c.FormFile("video")

	err = ctrl.handler.UpdatePost(
		c,
		creator,
		postUUID,
		req.Described,
		req.ImageDel,
		application.WithImagesMultipart(images),
		application.WithVideoMultipart(video),
	)
	if err != nil {
		switch err {
		case application.ErrUnauthorizedCreator:
			ginAbortUnauthorized(c, responsevalue.CodeInvalidateUser, responsevalue.MsgUnauthorized, nil)
			return

		case entity.ErrTooManyImages:
			ginAbortNotAcceptable(c, responsevalue.CodeMaxImagesReached, "too many images", nil)
			return

		case entity.ErrInvalidImageExtension:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid image extension", nil)
			return

		case entity.ErrMediaImageTooBig:
			ginAbortNotAcceptable(c, responsevalue.CodeFileTooBig, "invalid image size", nil)
			return

		case entity.ErrInvalidVideoExtension:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid video extension", nil)
			return

		case entity.ErrMediaVideoTooBig:
			ginAbortNotAcceptable(c, responsevalue.CodeFileTooBig, "invalid video size", nil)
			return

		case entity.ErrInvalidVideoDuration:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid video duration", nil)
			return
		}
		return
	}

	resp := new(dto.DefaultResp)
	resp.SetCode(responsevalue.CodeOK)
	c.JSON(http.StatusOK, resp)
}