package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrPermissionDenied      = errors.New("permission denied")
	ErrInvalidCommentContent = errors.New("invalid comment content")
	ErrNotCreator            = errors.New("not creator")
)

type Comment struct {
	ID        uuid.UUID
	Content   string
	Creator   *User
	Post      *Post
	CreatedAt time.Time
}

func (c Comment) IDString() string {
	return c.ID.String()
}

func (c Comment) GetContent() string {
	return c.Content
}

func (c Comment) GetCreator() *User {
	return c.Creator
}

func (c Comment) CreatorUUID() uuid.UUID {
	return c.Creator.ID()
}

func (c Comment) GetPost() *Post {
	return c.Post
}

func (c Comment) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func commentContentCheck(content string) bool {
	return len(content) <= 500
}

func (c *Comment) UpdateContent(updater *User, content string) error {
	// permission check
	if c.Creator.ID() != updater.id {
		return ErrPermissionDenied
	}

	if updater.IsLocked() || c.Creator.IsLocked() {
		return ErrLockedUser
	}

	if c.Post.IsLocked() {
		return ErrLockedPost
	}

	if !commentContentCheck(content) {
		return ErrInvalidCommentContent
	}

	c.Content = content

	return nil
}

func (c *Comment) IsDeletable(deleter *User) error {
	if deleter.IsLocked() {
		return ErrLockedUser
	}

	if c.Post.IsLocked() {
		return ErrLockedPost
	}

	if c.Creator.ID() != deleter.ID() {
		return ErrNotCreator
	}

	return nil
}
