package entity

import (
	"os"
	"time"

	"github.com/google/uuid"
)

func NewMediaFromDB(id, owner, mediaType, path string) (*Media, error) {
	mediaID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	ownerID, err := uuid.Parse(owner)
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &Media{
		id:        mediaID,
		owner:     ownerID,
		mediaType: MediaType(mediaType),
		path:      path,
	}, nil
}

func NewPostFromDB(id, creator, status, content string, media []Media, canComment bool, createdAt, updatedAt time.Time) (*Post, error) {
	postID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	creatorID, err := uuid.Parse(creator)
	if err != nil {
		return nil, err
	}

	return &Post{
		id:         postID,
		creator:    creatorID,
		status:     PostStatus(status),
		content:    content,
		media:      media,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
		CanComment: canComment,
	}, nil
}
