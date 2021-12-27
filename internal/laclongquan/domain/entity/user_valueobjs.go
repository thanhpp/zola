package entity

import "github.com/google/uuid"

type UserState string

func (s UserState) String() string {
	return string(s)
}

const (
	UserStateActive UserState = "active"
	UserStateLocked UserState = "locked"
)

type UserRole string

func (r UserRole) String() string {
	return string(r)
}

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type UserAddress struct {
	Address string
	City    string
	Country string
}

type UserProfileMediaType string

const (
	UserProfileAvatar UserProfileMediaType = "avatar"
	UserProfileCover  UserProfileMediaType = "cover"
)

type UserProfileMedia struct {
	ID        string
	Type      UserProfileMediaType
	LocalPath string
}

func (fac userFactoryImpl) NewUserMedia(user *User, mediaType UserProfileMediaType, path string) (*UserProfileMedia, error) {
	return &UserProfileMedia{
		ID:        uuid.NewString(),
		Type:      mediaType,
		LocalPath: path,
	}, nil
}
