package entity

import (
	"strings"

	"github.com/google/uuid"
)

type Post struct {
	id      uuid.UUID
	creator uuid.UUID
	content string
	media   []Media
}

func (p Post) ID() string {
	return p.id.String()
}

func (p Post) Creator() string {
	return p.creator.String()
}

func (p Post) Content() string {
	return p.content
}

func (p Post) Media() []Media {
	return p.media
}

func contentLengthCheck(content string) bool {
	if strings.Count(content, " ") > 500 {
		return false
	}

	return true
}
