package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotAFriendRequest = errors.New("not a friend request")
	ErrSelfRelation      = errors.New("self relation")
	ErrInvalidUser       = errors.New("invalid user")
	ErrAlreadyBlocked    = errors.New("already blocked")
	ErrInvalidRelation   = errors.New("invalid relation")
)

type RelationStatus string

func (rs RelationStatus) String() string {
	return string(rs)
}

const (
	RelationRequesting   RelationStatus = "requesting"
	RelationFriend       RelationStatus = "friend"
	RelationBlocked      RelationStatus = "blocked"
	RelationDiaryBlocked RelationStatus = "diary_blocked"
)

// Relation there are 3 relationshop within this application
// only the friend request need to specify
// the requestor (userA) and the requestee (userB)
type Relation struct {
	UserA     uuid.UUID
	UserB     uuid.UUID
	Status    RelationStatus
	CreatedAt time.Time
}

func (r Relation) UserAID() uuid.UUID {
	return r.UserA
}

func (r Relation) UserAIDStr() string {
	return r.UserA.String()
}

func (r Relation) UserBID() uuid.UUID {
	return r.UserB
}

func (r Relation) UserBIDStr() string {
	return r.UserB.String()
}

func (r Relation) IsBlock() bool {
	return r.Status == RelationBlocked
}

func (r Relation) IsFriend() bool {
	return r.Status == RelationFriend
}

func (r Relation) IsDiaryBlocked() bool {
	return r.Status == RelationDiaryBlocked
}

func (r *Relation) BlockDiary(blocker, blocked *User) error {
	if r == nil || blocker == nil || blocked == nil {
		return ErrNilUser
	}

	if r.UserA != blocker.ID() || r.UserB != blocked.ID() {
		return ErrInvalidUser
	}

	if blocker.IsLocked() || blocked.IsLocked() {
		return ErrLockedUser
	}

	if r.IsBlock() {
		return ErrAlreadyBlocked
	}

	if !r.IsFriend() {
		return ErrInvalidRelation
	}

	r.Status = RelationDiaryBlocked

	return nil
}

func (r *Relation) UnblockDiary(blocker, blocked *User) error {
	if r == nil || blocker == nil || blocked == nil {
		return ErrNilUser
	}

	if r.UserA != blocker.ID() || r.UserB != blocked.ID() {
		return ErrInvalidUser
	}

	if blocker.IsLocked() || blocked.IsLocked() {
		return ErrLockedUser
	}

	if !r.IsDiaryBlocked() {
		return ErrInvalidRelation
	}

	r.Status = RelationFriend

	return nil
}

func (r *Relation) AcceptFriendRequest() error {
	if r.Status != RelationRequesting {
		return ErrNotAFriendRequest
	}

	r.Status = RelationFriend

	return nil
}

func (r *Relation) RejectFriendRequest() error {
	if r.Status != RelationRequesting {
		return ErrNotAFriendRequest
	}

	return nil
}

func (r *Relation) Block() {
	r.Status = RelationBlocked
}

func (r Relation) CreatedAtUnix() int64 {
	return r.CreatedAt.Unix()
}

var (
	ErrLockedUser = errors.New("locked user")
)

func (fac userFactoryImpl) preRelationCheck(userA, userB *User) error {
	if userA.Equal(userB) {
		return ErrSelfRelation
	}

	if userA.IsLocked() || userB.IsLocked() {
		return ErrLockedUser
	}

	return nil
}

func (fac userFactoryImpl) NewFriendRequest(requestor, requestee *User) (*Relation, error) {
	if err := fac.preRelationCheck(requestor, requestee); err != nil {
		return nil, err
	}

	return &Relation{
		UserA:  requestor.ID(),
		UserB:  requestee.ID(),
		Status: RelationRequesting,
	}, nil
}

func (fac userFactoryImpl) NewBlockRelation(blocker, blocked *User) (*Relation, error) {
	if err := fac.preRelationCheck(blocker, blocked); err != nil {
		return nil, err
	}

	return &Relation{
		UserA:  blocker.ID(),
		UserB:  blocked.ID(),
		Status: RelationBlocked,
	}, nil
}

func (fac userFactoryImpl) NewDiaryBlockRelation(blocker, blocked *User) (*Relation, error) {
	if err := fac.preRelationCheck(blocker, blocked); err != nil {
		return nil, err
	}

	return &Relation{
		UserA:  blocker.ID(),
		UserB:  blocked.ID(),
		Status: RelationDiaryBlocked,
	}, nil
}

func NewRelationFromDB(userAIDStr, userBIDStr, status string, createdAt time.Time) (*Relation, error) {
	userAID, err := uuid.Parse(userAIDStr)
	if err != nil {
		return nil, err
	}

	userBID, err := uuid.Parse(userBIDStr)
	if err != nil {
		return nil, err
	}

	return &Relation{
		UserA:     userAID,
		UserB:     userBID,
		Status:    RelationStatus(status),
		CreatedAt: createdAt,
	}, nil
}
