package entity

type PostStatus string

var (
	PostStatusActive PostStatus = "active"
	PostStatusLocked PostStatus = "locked"
)

func (s PostStatus) String() string {
	return string(s)
}
