package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Block struct {
	Blocker uuid.UUID
	Blocked uuid.UUID
}

func (b *Block) GetBlocker() string {
	return b.Blocker.String()
}

func (b *Block) GetBlocked() string {
	return b.Blocked.String()
}

var (
	ErrSelfBlock = errors.New("can't block yourself")
)

func (fac userFactoryImpl) NewBlock(blocker, blocked *User) (*Block, error) {
	if blocker.ID().String() == blocked.ID().String() {
		return nil, ErrSelfBlock
	}

	return &Block{
		Blocker: blocker.ID(),
		Blocked: blocked.ID(),
	}, nil
}
