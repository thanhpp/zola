package entity

import (
	"github.com/google/uuid"
)

const (
	maximumVideoSize     int64   = 10 * 1024 * 1024 // 10MB
	minimumVideoDuration float64 = 1                // second
	maximumVideoDuration float64 = 10               // second
)

func (f postFactoryImpl) NewMediaVideo(path string, size int64, owner uuid.UUID) (*Media, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	if !extensionCheck(path, ".mp4") {
		return nil, ErrInvalidVideoExtension
	}

	if size > maximumVideoSize {
		return nil, ErrMediaVideoTooBig
	}

	return &Media{
		id:        id,
		mediaType: MediaTypeVideo,
		owner:     owner,
		path:      path,
	}, nil
}

func (m Media) DurationCheck(dur float64) error {
	if dur < minimumVideoDuration || dur > maximumVideoDuration {
		return ErrInvalidVideoDuration
	}

	return nil
}
