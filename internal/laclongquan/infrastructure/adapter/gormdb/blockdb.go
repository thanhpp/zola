package gormdb

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"gorm.io/gorm"
)

type BlockDB struct {
	Blocker   string `gorm:"Column:blocker; Type:text; primaryKey"`
	Blocked   string `gorm:"Column:blocked; Type:text; primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type blockGorm struct {
	db    *gorm.DB
	model *BlockDB
}

func (b blockGorm) marshalBlock(block *entity.Block) *BlockDB {
	return &BlockDB{
		Blocker: block.Blocker.String(),
		Blocked: block.Blocked.String(),
	}
}

func (b blockGorm) unmarshalBlock(blockDB *BlockDB) (*entity.Block, error) {
	blocker, err := uuid.Parse(blockDB.Blocker)
	if err != nil {
		return nil, err
	}

	blocked, err := uuid.Parse(blockDB.Blocked)
	if err != nil {
		return nil, err
	}
	return &entity.Block{
		Blocker: blocker,
		Blocked: blocked,
	}, nil
}

func (b blockGorm) Create(ctx context.Context, block *entity.Block) error {
	err := b.db.WithContext(ctx).Model(b.model).Create(b.marshalBlock(block)).Error
	if err != nil {
		if isDuplicate(err) {
			return repository.ErrUserAlreadyBlocked
		}

		return err
	}

	return nil
}

func (b blockGorm) Get(ctx context.Context, blocker, blocked string) (*entity.Block, error) {
	var blockDB = new(BlockDB)

	if err := b.db.WithContext(ctx).Model(b.model).Where("blocker = ? AND blocked = ?", blocker, blockDB).Take(blockDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrBlockNotFound
		}

		return nil, err
	}

	return b.unmarshalBlock(blockDB)
}

func (b blockGorm) Delete(ctx context.Context, blocker, blocked string) error {
	return b.db.WithContext(ctx).Model(b.model).
		Where("blocker = ? AND blocked = ?", blocker, blocked).
		Delete(BlockDB{}).Error
}
