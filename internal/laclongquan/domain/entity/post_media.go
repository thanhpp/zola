package entity

import (
	"path/filepath"

	"github.com/google/uuid"
)

type MediaType string

const (
	MediaTypeImage MediaType = "image"
	MediaTypeVideo MediaType = "video"
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

func (m Media) IsOwner(user *User) bool {
	return m.owner.String() == user.ID().String()
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
