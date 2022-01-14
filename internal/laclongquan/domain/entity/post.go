package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNilUser   = errors.New("nil user")
	ErrNotFriend = errors.New("not friend")
)

type Post struct {
	id         uuid.UUID
	creator    uuid.UUID
	content    string
	status     PostStatus
	media      []Media
	CanComment bool
	createdAt  time.Time
	updatedAt  time.Time
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

func (p Post) GetCanComment() bool {
	return p.CanComment
}

func (p Post) CanCreateComment(postCreator, commentCreator *User, relation *Relation) error {
	if postCreator == nil || commentCreator == nil {
		return ErrEmptyInput
	}

	if postCreator.Equal(commentCreator) {
		return nil
	}

	if relation == nil {
		return ErrEmptyInput
	}

	if commentCreator.IsLocked() {
		return ErrLockedUser
	}

	if p.IsLocked() {
		return ErrLockedPost
	}

	if !p.CanComment {
		return ErrPermissionDenied
	}

	if relation.IsBlock() {
		return ErrAlreadyBlocked
	}

	if !relation.IsFriend() {
		return ErrNotFriend
	}

	return nil
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

func (p Post) CreatedAt() int64 {
	return p.createdAt.Unix()
}

func (p Post) UpdatedAt() int64 {
	return p.updatedAt.Unix()
}

func (p Post) CanUserGetMedia(user *User, relation *Relation, mediaID string) (*Media, error) {
	if !user.IsAdmin() {
		if user.IsLocked() {
			return nil, ErrLockedUser
		}

		if p.IsLocked() {
			return nil, ErrLockedPost
		}

		if (user.ID().String() != p.Creator() && relation == nil) || (relation != nil && relation.IsFriend()) {
			return nil, ErrPermissionDenied
		}
	}

	for i := range p.media {
		if mediaID == p.media[i].ID() {
			return &p.media[i], nil
		}
	}

	return nil, ErrPostNotContainsMedia
}

func (p Post) CanUserGetPost(user *User, relation *Relation) error {
	if user.IsAdmin() {
		return nil
	}

	if user.IsLocked() {
		return ErrLockedUser
	}

	if p.IsLocked() {
		return ErrLockedPost
	}

	if (user.ID().String() != p.Creator() && relation == nil) || (relation != nil && relation.IsFriend()) {
		return ErrPermissionDenied
	}

	return nil
}

func (p Post) CanUserEditPost(user *User) error {
	if user == nil {
		return ErrNilUser
	}

	if user.IsLocked() {
		return ErrLockedUser
	}

	if p.IsLocked() {
		return ErrLockedPost
	}

	if p.Creator() != user.ID().String() {
		return ErrNotCreator
	}

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

func (p *Post) CanBeDeletedBy(user *User) bool {
	if user.IsAdmin() {
		return true
	}

	return p.Creator() == user.ID().String()
}

func (p *Post) UpdateCanComment(user *User, value bool) error {
	if p == nil || user == nil {
		return ErrEmptyInput
	}

	if user.IsAdmin() {
		p.CanComment = value
		return nil
	}

	if !p.IsCreator(user.ID()) {
		return ErrNotCreator
	}

	if p.IsLocked() {
		return ErrLockedPost
	}

	if user.IsLocked() {
		return ErrLockedUser
	}

	p.CanComment = value
	return nil
}
