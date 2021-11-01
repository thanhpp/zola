package entity

import (
	"github.com/google/uuid"
	"github.com/thanhpp/zola/pkg/logger"
)

const (
	maximumImageSize int64 = 4 * 1024 * 1024 // 4MB
)

func (f postFactoryImpl) NewMediaImage(path string, size int64, owner uuid.UUID) (*Media, error) {
	id := uuid.New()
	logger.Debug(id.String())

	if !extensionCheck(path, ".jpg", ".jpeg", ".png") {
		return nil, ErrInvalidImageExtension
	}

	if size > maximumImageSize {
		return nil, ErrMediaImageTooBig
	}

	return &Media{
		id:        id,
		mediaType: MediaTypeImage,
		owner:     owner,
		path:      path,
	}, nil
}
