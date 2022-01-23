package gormdb

import (
	"context"
	"errors"
	"time"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostDB struct {
	PostUUID   string     `gorm:"Column:post_uuid; Type:text; primaryKey"`
	Status     string     `gorm:"Column:status; Type:text"`
	Content    string     `gorm:"Column:context; Type:text"`
	Creator    string     `gorm:"Column:creator; Type:text; index"`
	CanComment bool       `gorm:"Column:can_comment; Type:boolean; default:true"`
	CreatedAt  time.Time  `gorm:"Column:created_at"`
	UpdatedAt  time.Time  `gorm:"Column:updated_at"`
	MediaDB    []*MediaDB `gorm:"foreignKey:PostUUID"`
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
		PostUUID:   post.ID(),
		Content:    post.Content(),
		Status:     post.Status().String(),
		Creator:    post.Creator(),
		MediaDB:    mediaDB,
		CanComment: post.GetCanComment(),
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

	return entity.NewPostFromDB(
		postDB.PostUUID,
		postDB.Creator,
		postDB.Status,
		postDB.Content,
		mediaList,
		postDB.CanComment,
		postDB.CreatedAt,
		postDB.UpdatedAt,
	)
}

func (p postGorm) GetByID(ctx context.Context, id string) (*entity.Post, error) {
	var (
		postDB = new(PostDB)
	)

	err := p.getByIDTx(ctx, p.db, id, postDB)
	if err != nil {
		return nil, err
	}

	post, err := p.unmarshalPost(postDB)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p postGorm) getByIDTx(ctx context.Context, tx *gorm.DB, id string, expect *PostDB) error {
	err := tx.WithContext(ctx).Model(p.postModel).Preload(clause.Associations).Where("post_uuid = ?", id).Take(expect).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repository.ErrPostNotFound
		}
		return err
	}
	return nil
}

func (p postGorm) GetMediaByID(ctx context.Context, id string) (*entity.Media, error) {
	var (
		mediaDB = new(MediaDB)
	)

	err := p.db.WithContext(ctx).Model(p.mediaModel).Where("media_uuid = ?", id).Take(mediaDB).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrMediaNotFound
		}
		return nil, err
	}

	return p.unmarshalMedia(mediaDB)
}

func (p postGorm) GetListPost(ctx context.Context, requestorID string, timeMileStone time.Time, offset, limit int) ([]*entity.Post, int, error) {
	var (
		list     []*PostDB
		newItems int64
	)

	stmt := p.db.WithContext(ctx).Model(p.postModel).
		Preload(clause.Associations)

	p.joinGetListPostFromActiveFriends(stmt, requestorID)

	if !timeMileStone.IsZero() {
		stmt.Where("post_db.created_at < ?", timeMileStone)
	}

	if err := stmt.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, -1, err
	}

	var listPost = make([]*entity.Post, 0, len(list))
	for i := range list {
		post, err := p.unmarshalPost(list[i])
		if err != nil {
			return nil, -1, err
		}
		listPost = append(listPost, post)
	}

	if !timeMileStone.IsZero() {
		stmt2 := p.db.WithContext(ctx).Model(p.postModel)
		p.joinGetListPostFromActiveFriends(stmt2, requestorID)
		if err := stmt2.Where("post_db.created_at > ?", timeMileStone).Count(&newItems).Error; err != nil {
			return nil, 0, err
		}
	}

	return listPost, int(newItems), nil
}

func (p postGorm) GetListPostForAdmin(ctx context.Context, requestorID string, timeMileStone time.Time, offset, limit int) ([]*entity.Post, int, error) {
	var (
		list     []*PostDB
		newItems int64
	)

	stmt := p.db.WithContext(ctx).Model(p.postModel).
		Distinct("post_db.*").
		Preload(clause.Associations).Order("created_at desc")

	if !timeMileStone.IsZero() {
		stmt.Where("post_db.created_at < ?", timeMileStone)
	}

	if err := stmt.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, -1, err
	}

	var listPost = make([]*entity.Post, 0, len(list))
	for i := range list {
		post, err := p.unmarshalPost(list[i])
		if err != nil {
			return nil, -1, err
		}
		listPost = append(listPost, post)
	}

	if !timeMileStone.IsZero() {
		stmt2 := p.db.WithContext(ctx).Model(p.postModel).Distinct("post_db.*")
		if err := stmt2.Where("post_db.created_at > ?", timeMileStone).Order("created_at desc").Count(&newItems).Error; err != nil {
			return nil, 0, err
		}
	}

	return listPost, int(newItems), nil
}

func (p postGorm) joinGetListPostFromActiveFriends(db *gorm.DB, requestorID string) {
	db.Order("post_db.created_at desc").
		Joins(`
	LEFT JOIN user_db ON (post_db.creator = user_db.user_uuid AND user_db.state = 'active')
	LEFT JOIN relation_db 
	ON (
		(
			(relation_db.user_a = post_db.creator AND relation_db.user_b = ?)
			OR
			(relation_db.user_a = ? AND relation_db.user_b = post_db.creator)
			AND 
			(relation_db.status = 'friend')
		)
		OR 
		(
			post_db.creator = ?
		)
	) 
`, requestorID, requestorID, requestorID)
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
	return gDB.Transaction(func(tx *gorm.DB) error {
		var postDB = new(PostDB)
		if err := p.getByIDTx(ctx, tx, id, postDB); err != nil {
			return err
		}

		post, err := p.unmarshalPost(postDB)
		if err != nil {
			return err
		}

		post, err = fn(ctx, post)
		if err != nil {
			return err
		}

		postDB = p.marshalPost(post)

		// add media records
		var existMediaIDs = make([]string, 0, len(postDB.MediaDB))
		for i := range postDB.MediaDB {
			// if the media is already exist, ignore the error
			if err := tx.WithContext(ctx).Model(p.mediaModel).FirstOrCreate(postDB.MediaDB[i]).Error; err != nil {
				return err
			}
			existMediaIDs = append(existMediaIDs, postDB.MediaDB[i].MediaUUID)
		}

		// remove media
		if err := tx.WithContext(ctx).Model(p.mediaModel).
			Where("media_uuid NOT IN ? AND post_uuid = ?", existMediaIDs, postDB.PostUUID).
			Delete(p.mediaModel).Error; err != nil {
			return err
		}

		// save post data
		if err := tx.WithContext(ctx).Model(p.postModel).
			Preload(clause.Associations).
			Where("post_uuid = ?", id).Updates(postDB).Error; err != nil {
			return err
		}

		return nil
	})
}
func (p postGorm) Delete(ctx context.Context, id string) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		// delete likes
		if err := tx.WithContext(ctx).Model(&LikeDB{}).Where("post_uuid = ?", id).Delete(&LikeDB{}).Error; err != nil {
			return err
		}

		// delete comments
		if err := tx.WithContext(ctx).Model(&CommentDB{}).Where("post_uuid = ?", id).Delete(&CommentDB{}).Error; err != nil {
			return err
		}

		// delete media
		if err := tx.WithContext(ctx).Model(p.mediaModel).Where("post_uuid = ?", id).Delete(p.mediaModel).Error; err != nil {
			return err
		}

		// delete reports
		if err := tx.WithContext(ctx).Model(&ReportDB{}).Where("post_uuid = ?", id).Delete(&ReportDB{}).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(p.postModel).Where("post_uuid = ?", id).Delete(p.postModel).Error; err != nil {
			return err
		}

		return nil
	})
}
