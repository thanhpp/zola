package entity

import "errors"

var (
	ErrTooManyImages        = errors.New("too many images")
	ErrTooManyVideos        = errors.New("too many videos")
	ErrContentTooLong       = errors.New("content too long")
	ErrPostNotContainsMedia = errors.New("post does not contains media with given id")
)

var (
	ErrMediaImageTooBig      = errors.New("image is too big")
	ErrInvalidImageExtension = errors.New("invalid image extension")
	ErrMediaVideoTooBig      = errors.New("video is too big")
	ErrInvalidVideoExtension = errors.New("invalid video extension")
	ErrInvalidVideoDuration  = errors.New("invalid video duration")
)
