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
	UserUUID    string       `gorm:"Column:user_uuid; Type:text; primaryKey; not null"`
	Name        string       `gorm:"Column:name; Type:text"`
	State       string       `gorm:"Column:state; Type:text"`
	Phone       string       `gorm:"Column:phone; Type:text; unique; index"`
	HashPass    string       `gorm:"Column:hash_pass; Type:text"`
	Role        string       `gorm:"Column:role; Type:text"`
	Link        string       `gorm:"Column:link; Type:text"`
	Avatar      string       `gorm:"Column:avatar; Type:text"`
	CoverImage  string       `gorm:"Column:cover_image; Type:text"`
	Username    string       `gorm:"Column:username; Type:text"`
	Description string       `gorm:"Column:description; Type:text"`
	Address     string       `gorm:"Column:address; Type:text"`
	City        string       `gorm:"Column:city; Type:text"`
	Country     string       `gorm:"Column:country; Type:text"`
	CreatedAt   time.Time    `gorm:"Column:created_at"`
	UpdatedAt   time.Time    `gorm:"Column:updated_at"`
	DeletedAt   sql.NullTime `gorm:"Column:deleted_at"`
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
		UserUUID:    user.ID().String(),
		Name:        user.Name(),
		Phone:       user.Account().Phone,
		HashPass:    user.Account().HashPass,
		State:       user.State().String(),
		Role:        user.Role(),
		Link:        user.GetLink(),
		Avatar:      user.GetAvatar(),
		CoverImage:  user.GetCoverImage(),
		Username:    user.GetUsername(),
		Description: user.GetDescription(),
		Address:     user.GetAddress(),
		City:        user.GetCity(),
		Country:     user.GetCountry(),
	}, nil
}

func (u userGorm) unmarshalUser(userDB *UserDB) (*entity.User, error) {
	if userDB == nil {
		return nil, errors.New("nil input")
	}

	return entity.NewUserFromDB(
		userDB.UserUUID,
		userDB.Name,
		userDB.Phone,
		userDB.HashPass,
		userDB.State,
		userDB.Role,
		userDB.Link,
		userDB.Avatar,
		userDB.CoverImage,
		userDB.Username,
		userDB.Description,
		userDB.Address,
		userDB.City,
		userDB.Country,
		userDB.CreatedAt,
	)
}

func (u userGorm) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var (
		userDB = new(UserDB)
	)

	err := u.db.WithContext(ctx).Model(u.model).Where("user_uuid = ?", id).Take(userDB).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}

	return u.unmarshalUser(userDB)
}

func (u userGorm) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	var (
		userDB = new(UserDB)
	)

	err := u.db.WithContext(ctx).Model(u.model).Where("phone = ?", phone).Take(userDB).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrUserNotFound
		}

		return nil, err
	}

	return u.unmarshalUser(userDB)
}

func (u userGorm) GetAllUsers(ctx context.Context, offset, limit int, sortBy, order, usernameLike, phoneLike string) ([]*entity.User, int, error) {
	var (
		list  []*UserDB
		total = new(int64)
	)

	stmt := u.db.WithContext(ctx).Model(u.model)

	if len(sortBy) != 0 {
		if len(order) == 0 {
			order = "asc"
		}
		stmt.Order(sortBy + " " + order)
	}

	if len(usernameLike) != 0 {
		stmt.Where("username LIKE ?", "%"+usernameLike+"%")
	}

	if len(phoneLike) != 0 {
		stmt.Where("phone LIKE ?", "%"+phoneLike+"%")
	}

	if err := stmt.Count(total).Error; err != nil {
		return nil, 0, err
	}

	if offset >= 0 {
		stmt.Offset(offset)
	}

	if limit > 0 {
		stmt.Limit(limit)
	}

	if err := stmt.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	users := make([]*entity.User, 0, len(list))
	for i := range list {
		user, err := u.unmarshalUser(list[i])
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	return users, int(*total), nil
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

	userDB, err := u.marshalUser(user)
	if err != nil {
		return err
	}
	// logger.Debugf("userDB - avatar: %s", userDB.Avatar)

	return u.db.WithContext(ctx).Model(u.model).
		Where("user_uuid = ?", id).
		Updates(userDB).Error
}

func (u userGorm) DeleteByID(ctx context.Context, id string) error {
	return u.db.WithContext(ctx).Model(u.model).Where("user_uuid = ?", id).Delete(u.model).Error
}
