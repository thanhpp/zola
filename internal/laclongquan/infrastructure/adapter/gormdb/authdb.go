package gormdb

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
	"gorm.io/gorm"
)

type AuthDB struct {
	ID        string `gorm:"Column:id; Type:text; PRIMARY KEY"`
	UserID    string `gorm:"Column:user_id; Type:text"`
	ExpiredAt int64  `gorm:"Column:expired_at; Type:bigint"`
	CreatedAt int64  `gorm:"Column:created_at; Type:bigint"`
}

type authGorm struct {
	db    *gorm.DB
	model *AuthDB
}

func (a authGorm) marshalClaims(claims *auth.Claims) *AuthDB {
	return &AuthDB{
		ID:        claims.Id,
		UserID:    claims.User.ID,
		ExpiredAt: claims.ExpiresAt,
		CreatedAt: claims.IssuedAt,
	}
}

func (a authGorm) CheckByID(ctx context.Context, id string) error {
	err := a.db.WithContext(ctx).Model(a.model).Where("id = ?", id).Take(&AuthDB{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return auth.ErrTokenNotFound
		}

		return err
	}

	return nil
}

func (a authGorm) Cache(ctx context.Context, claims *auth.Claims) error {
	authDB := a.marshalClaims(claims)
	return a.db.WithContext(ctx).Model(a.model).Create(authDB).Error
}

func (a authGorm) Delete(ctx context.Context, id string) error {
	return a.db.WithContext(ctx).Model(a.model).Where("id = ?", id).Delete(&AuthDB{}).Error
}

func (a authGorm) DeleteByUserID(ctx context.Context, userID string) error {
	return a.db.WithContext(ctx).Model(a.model).Where("user_id = ?", userID).Delete(&AuthDB{}).Error
}

func (a authGorm) DeleteExpired(ctx context.Context) error {
	return a.db.WithContext(ctx).Model(a.model).Where("expired_at <= ?", time.Now().Unix()).Delete(&AuthDB{}).Error
}
