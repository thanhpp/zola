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

func (ctrl PostController) CreatePost(c *gin.Context) {
	var (
		req = new(dto.CreatePostReq)
	)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind req %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
		return
	}

	creator, err := getUserUUIDFromCtx(c)
	if err != nil {
		logger.Errorf("get user uuid %v", err)
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
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

	post, err := ctrl.handler.CreatePostWithMultipart(
		c,
		creator,
		req.Described,
		application.WithImagesMultipart(images),
		application.WithVideoMultipart(video))
	if err != nil {
		logger.Errorf("create post %v", err)
		switch err {
		case application.ErrHasVideoAndImages:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "having images and video", nil)
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

		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
		return
	}

	// FIXME: missing URL
	resp := new(dto.CreatePostResp)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetData(post.ID())

	c.JSON(http.StatusOK, resp)
}
