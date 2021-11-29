package entity

import (
	"time"

	"github.com/google/uuid"
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
