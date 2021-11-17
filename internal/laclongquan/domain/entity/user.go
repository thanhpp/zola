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

type User struct {
	id      uuid.UUID
	name    string
	avatar  string
	state   UserState
	account Account
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Avatar() string {
	return u.avatar
}

func (u User) Account() Account {
	return u.account
}

func (u User) PassEqual(pass string) error {
	return u.account.Equal(u.account.Phone, pass)
}

func (u User) State() UserState {
	return u.state
}
