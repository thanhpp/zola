package entity

import "errors"

type PostStatus string

var (
	PostStatusActive PostStatus = "active"
	PostStatusLocked PostStatus = "locked"
)

func (s PostStatus) String() string {
	return string(s)
}

var (
	ErrLockedPost = errors.New("locked post")
)
