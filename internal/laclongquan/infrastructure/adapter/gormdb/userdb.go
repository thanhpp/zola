package gormdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"gorm.io/gorm"
)

type UserDB struct {
	UserUUID  string       `gorm:"Column:user_uuid; Type:text; primaryKey; not null"`
	Name      string       `gorm:"Column:name; Type:text"`
	Avatar    string       `gorm:"Column:avatar; Type:text"`
	Phone     string       `gorm:"Column:phone; Type:text; unique; index"`
	HashPass  string       `gorm:"Column:hash_pass; Type:text"`
	CreatedAt time.Time    `gorm:"Column:created_at"`
	UpdatedAt time.Time    `gorm:"Column:updated_at"`
	DeletedAt sql.NullTime `gorm:"Column:deleted_at"`
}

type userGorm struct {
	db    *gorm.DB
	model *UserDB
}

func (u userGorm) marshalUser(user *entity.User) (*UserDB, error) {
	if user == nil {
		return nil, errors.New("nil input")
	}

	return &UserDB{
		UserUUID: user.ID().String(),
		Name:     user.Name(),
		Avatar:   user.Avatar(),
		Phone:    user.Account().Phone,
		HashPass: user.Account().HashPass,
	}, nil
}

func (u userGorm) unmarshalUser(userDB *UserDB) (*entity.User, error) {
	if userDB == nil {
		return nil, errors.New("nil input")
	}

	return entity.NewUserFromDB(
		userDB.UserUUID,
		userDB.Name,
		userDB.Avatar,
		userDB.Phone,
		userDB.HashPass,
	)
}

func (u userGorm) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var (
		userDB = new(UserDB)
	)

	err := u.db.WithContext(ctx).Model(u.model).Where("user_uuid = ?", id).Find(userDB).Error
	if err != nil {
		return nil, err
	}

	return u.unmarshalUser(userDB)
}

func (u userGorm) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	var (
		userDB = new(UserDB)
	)

	err := u.db.WithContext(ctx).Model(u.model).Where("phone = ?", phone).Find(userDB).Error
	if err != nil {
		return nil, err
	}

	return u.unmarshalUser(userDB)
}

func (u userGorm) Create(ctx context.Context, user *entity.User) error {
	userDB, err := u.marshalUser(user)
	if err != nil {
		return err
	}

	err = u.db.WithContext(ctx).Model(u.model).Create(userDB).Error
	if err != nil {
		if isDuplicate(err) {
			return repository.ErrDuplicateUser
		}
		return err
	}

	return nil
}

func (u userGorm) Update(ctx context.Context, id string, fn repository.UserUpdateFunc) error {
	user, err := u.GetByID(ctx, id)
	if err != nil {
		return err
	}

	user, err = fn(ctx, user)
	if err != nil {
		return err
	}

	return u.db.WithContext(ctx).Model(u.db).Updates(user).Error
}

func (u userGorm) DeleteByID(ctx context.Context, id string) error {
	return u.db.WithContext(ctx).Model(u.model).Where("id = ?", id).Delete(u.model).Error
}
