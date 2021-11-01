package entity

import (
	"errors"
	"path/filepath"

	"github.com/google/uuid"
)

type MediaType string

const (
	MediaTypeImage MediaType = "image"
	MediaTypeVideo MediaType = "video"
)

var (
	ErrMediaImageTooBig      = errors.New("image is too big")
	ErrInvalidImageExtension = errors.New("invalid image extension")
	ErrMediaVideoTooBig      = errors.New("video is too big")
	ErrInvalidVideoExtension = errors.New("invalid video extension")
	ErrInvalidVideoDuration  = errors.New("invalid video duration")
)

type Media struct {
	id        uuid.UUID
	owner     uuid.UUID
	mediaType MediaType
	size      int64 // bytes
	path      string
}

func (m Media) ID() string {
	return m.id.String()
}

func (m Media) Type() MediaType {
	return m.mediaType
}

func (m Media) Path() string {
	return m.path
}

func (m Media) Owner() string {
	return m.owner.String()
}

func (m Media) Size() int64 {
	return m.size
}

func extensionCheck(path string, exts ...string) bool {
	ext := filepath.Ext(path)
	
	for i := range exts {
		if ext == exts[i] {
			return true
		}
	}

	return false
}
