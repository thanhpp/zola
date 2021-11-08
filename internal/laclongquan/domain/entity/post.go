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

func (p Post) IsCreator(id uuid.UUID) bool {
	return p.creator.String() == id.String()
}

func (p Post) Creator() string {
	return p.creator.String()
}

func (p Post) CreatorUUID() uuid.UUID {
	return p.creator
}

func (p Post) Content() string {
	return p.content
}

func (p *Post) UpdateContent(content string) error {
	if !contentLengthCheck(content) {
		return ErrContentTooLong
	}

	p.content = content

	return nil
}

func (p Post) Media() []Media {
	return p.media
}

func (p *Post) AddMedia(m Media) error {
	// already have 4 images, can not add more
	if len(p.media) == 4 {
		return ErrTooManyImages
	}

	// post has 1 media
	if len(p.media) == 1 {
		// if it is a video, can not add more
		if p.media[0].mediaType == MediaTypeVideo {
			return ErrTooManyVideos
		}
	}

	p.media = append(p.media, m)

	return nil
}

func (p *Post) RemoveMedia(ids ...string) ([]*Media, error) {
	var deleted = make([]*Media, 0, len(ids))
	for i := range ids {
		for j := range p.media {
			if p.media[j].ID() == ids[i] {
				p.media = append(p.media[:j], p.media[j+1:]...)
				deleted = append(deleted, &p.media[j])
				break
			}
			return nil, ErrPostNotContainsMedia
		}
	}

	return deleted, nil
}

func contentLengthCheck(content string) bool {
	if strings.Count(content, " ") > 500 {
		return false
	}

	return true
}
