package application

import (
	"errors"
	"mime/multipart"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrHasVideoAndImages = errors.New("has both video and images")
)

type createPostMultipartConfig struct {
	images []*multipart.FileHeader
	video  *multipart.FileHeader
}

func (cfg createPostMultipartConfig) haveImages() bool {
	return len(cfg.images) > 0
}

func (cfg createPostMultipartConfig) haveVideo() bool {
	return cfg.video != nil
}

func (cfg createPostMultipartConfig) validate() error {
	if len(cfg.images) > 4 {
		return entity.ErrTooManyImages
	}

	if len(cfg.images) > 0 && cfg.video != nil {
		return ErrHasVideoAndImages
	}

	return nil
}

type CreatePostMultipartOption func(*createPostMultipartConfig)

func WithImagesMultipart(images []*multipart.FileHeader) CreatePostMultipartOption {
	return func(c *createPostMultipartConfig) {
		c.images = images
	}
}

func WithVideoMultipart(video *multipart.FileHeader) CreatePostMultipartOption {
	return func(c *createPostMultipartConfig) {
		c.video = video
	}
}
