package entity

import (
	"strings"

	"github.com/google/uuid"
)

type Post struct {
	id      uuid.UUID
	creator uuid.UUID
	content string
	status  PostStatus
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

func (p Post) Status() PostStatus {
	return p.status
}

func (p Post) IsLocked() bool {
	return p.status == PostStatusLocked
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
	// logger.Debugf("post media after added %v", p.media)
	p.media = append(p.media, m)

	return nil
}

func (p *Post) RemoveMedia(ids ...string) ([]*Media, error) {
	var (
		deleted    = make([]*Media, 0, len(ids))
		deleteflag bool
		newMedia   = make([]Media, 0, len(p.media))
	)

	for i := range p.media {
		deleteflag = false
		for j := range ids {
			if p.media[i].ID() == ids[j] {
				deleteflag = true
				deleted = append(deleted, &p.media[i])
				break
			}
			return nil, ErrPostNotContainsMedia
		}
		if !deleteflag {
			newMedia = append(newMedia, p.media[i])
		}
	}

	p.media = newMedia

	return deleted, nil
}

func contentLengthCheck(content string) bool {
	if strings.Count(content, " ") > 500 {
		return false
	}

	return true
}

func (p *Post) CanBeDeletedBy(userID string) bool {

	return p.Creator() == userID
}
