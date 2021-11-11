package entity

import (
	"github.com/google/uuid"
)

type Images struct {
	Path string
	Size int64
}

type PostFactory interface {
	NewPost(creator uuid.UUID, content string) (*Post, error)
	NewMediaImage(path string, size int64, owner uuid.UUID) (*Media, error)
	NewPostWithImages(creator uuid.UUID, content string, images ...Media) (*Post, error)
	NewMediaVideo(path string, size int64, owner uuid.UUID) (*Media, error)
	NewPostWithVideo(creator uuid.UUID, content string, video Media) (*Post, error)
}

func NewPostFactory() PostFactory {
	return new(postFactoryImpl)
}

type postFactoryImpl struct{}

func (fac postFactoryImpl) NewPost(creator uuid.UUID, content string) (*Post, error) {
	// content
	if !contentLengthCheck(content) {
		return nil, ErrContentTooLong
	}

	postID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &Post{
		id:      postID,
		creator: creator,
		content: content,
	}, nil
}

func (fac postFactoryImpl) NewPostWithImages(creator uuid.UUID, content string, images ...Media) (*Post, error) {
	// content
	if !contentLengthCheck(content) {
		return nil, ErrContentTooLong
	}

	// images check
	if len(images) > 4 {
		return nil, ErrTooManyImages
	}

	postID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &Post{
		id:      postID,
		creator: creator,
		content: content,
		media:   images,
	}, nil
}

func (fac postFactoryImpl) NewPostWithVideo(creator uuid.UUID, content string, video Media) (*Post, error) {
	if !contentLengthCheck(content) {
		return nil, ErrContentTooLong
	}

	postID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &Post{
		id:      postID,
		creator: creator,
		content: content,
		media:   []Media{video},
	}, nil
}
