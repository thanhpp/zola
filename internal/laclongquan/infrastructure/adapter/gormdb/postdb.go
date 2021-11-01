package gormdb

import (
	"context"
	"time"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostDB struct {
	PostUUID  string    `gorm:"Column:post_uuid; Type:text; primaryKey"`
	Content   string    `gorm:"Column:context; Type:text"`
	Creator   string    `gorm:"Column:creator; Type:text; index"`
	CreatedAt time.Time `gorm:"Column:created_at"`
	UpdatedAt time.Time `gorm:"Column:updated_at"`

	MediaDB []*MediaDB `gorm:"foreignKey:PostUUID"`
}

type MediaDB struct {
	MediaUUID string    `gorm:"Column:media_uuid; Type:text; primaryKey"`
	PostUUID  string    `gorm:"Column:post_uuid; Type:text"`
	Owner     string    `gorm:"Column:owner; Type:text"`
	Type      string    `gorm:"Column:type; Type:text"`
	Path      string    `gorm:"Column:path; Type:text"`
	CreatedAt time.Time `gorm:"Column:created_at"`
	UpdatedAt time.Time `gorm:"Column:updated_at"`
}

type postGorm struct {
	db         *gorm.DB
	postModel  *PostDB
	mediaModel *MediaDB
}

func (p postGorm) marshalMedia(postUUID string, media *entity.Media) *MediaDB {
	return &MediaDB{
		MediaUUID: media.ID(),
		PostUUID:  postUUID,
		Type:      string(media.Type()),
		Owner:     media.Owner(),
		Path:      media.Path(),
	}
}

func (p postGorm) marshalPost(post *entity.Post) *PostDB {
	var (
		postMedia = post.Media()
		mediaDB   = make([]*MediaDB, 0, len(postMedia))
	)

	for i := range postMedia {
		mediaDB = append(mediaDB, p.marshalMedia(post.ID(), &postMedia[i]))
	}

	return &PostDB{
		PostUUID: post.ID(),
		Content:  post.Content(),
		Creator:  post.Creator(),
		MediaDB:  mediaDB,
	}
}

func (p postGorm) unmarshalMedia(mediaDB *MediaDB) (*entity.Media, error) {
	return entity.NewMediaFromDB(mediaDB.MediaUUID, mediaDB.Owner, mediaDB.Type, mediaDB.Path)
}

func (p postGorm) unmarshalPost(postDB *PostDB) (*entity.Post, error) {
	var (
		mediaList = make([]entity.Media, 0, len(postDB.MediaDB))
	)

	for i := range postDB.MediaDB {
		media, err := p.unmarshalMedia(postDB.MediaDB[i])
		if err != nil {
			return nil, err
		}
		mediaList = append(mediaList, *media)
	}

	return entity.NewPostFromDB(postDB.PostUUID, postDB.Creator, postDB.Content, mediaList)
}

func (p postGorm) GetByID(ctx context.Context, id string) (*entity.Post, error) {
	var (
		postDB = new(PostDB)
	)

	err := p.db.WithContext(ctx).Model(p.postModel).
		Preload(clause.Associations).
		Take(postDB).Error
	if err != nil {
		return nil, err
	}

	post, err := p.unmarshalPost(postDB)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p postGorm) Create(ctx context.Context, post *entity.Post) error {
	postDB := p.marshalPost(post)

	return p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Model(p.postModel).Create(postDB).Error; err != nil {
			return err
		}

		return nil
	})
}

func (p postGorm) Update(ctx context.Context, id string, fn repository.PostUpdateFn) error {

	return nil
}
func (p postGorm) Delete(ctx context.Context, id string) error {

	return nil
}
