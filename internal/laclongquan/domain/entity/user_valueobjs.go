package entity

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type UserState string

func (s UserState) String() string {
	return string(s)
}

const (
	UserStateActive UserState = "active"
	UserStateLocked UserState = "locked"
)

var (
	ErrInvalidState = errors.New("invalid state")
)

type UserRole string

func (r UserRole) String() string {
	return string(r)
}

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

// ---------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- ADDRESS ----------------------------------------------------------
type UserAddress struct {
	Address string
	City    string
	Country string
}

var (
	ErrInvalidCountry = errors.New("invalid country")
	notSupportCountry = map[string]struct{}{
		"forbiden country": {},
	}
)

func countryCheck(country string) error {
	country = strings.ToLower(strings.TrimSpace(country))
	if _, ok := notSupportCountry[country]; ok {
		return ErrInvalidCountry
	}

	return nil
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
