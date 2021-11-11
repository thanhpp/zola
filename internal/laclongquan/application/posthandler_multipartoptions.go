package application

import (
	"errors"
	"mime/multipart"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrHasVideoAndImages = errors.New("has both video and images")
)

type multipartConfig struct {
	images []*multipart.FileHeader
	video  *multipart.FileHeader
}

func (cfg multipartConfig) haveImages() bool {
	return len(cfg.images) > 0
}

func (cfg multipartConfig) haveVideo() bool {
	return cfg.video != nil
}

func (cfg multipartConfig) validate() error {
	if len(cfg.images) > 4 {
		return entity.ErrTooManyImages
	}

	if len(cfg.images) > 0 && cfg.video != nil {
		return ErrHasVideoAndImages
	}

	return nil
}

type MultipartOption func(*multipartConfig)

func WithImagesMultipart(images []*multipart.FileHeader) MultipartOption {
	return func(c *multipartConfig) {
		c.images = images
	}
}

func WithVideoMultipart(video *multipart.FileHeader) MultipartOption {
	return func(c *multipartConfig) {
		c.video = video
	}
}
