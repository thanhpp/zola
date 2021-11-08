package application

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrUnauthorizedCreator = errors.New("unauthorized creator")
)

// UpdatePost
func (p PostHandler) UpdatePost(ctx context.Context, creator, postID uuid.UUID, content string, deleteMediaID []string, opts ...MultipartOption) error {
	var cfg = new(multipartConfig)
	for i := range opts {
		opts[i](cfg)
	}
	if err := cfg.validate(); err != nil {
		return err
	}

	var (
		deletedMedia []*entity.Media
		addedMedia   []*entity.Media
		err          error
	)

	if err := p.repo.Update(ctx, postID.String(), func(ctx context.Context, post *entity.Post) (*entity.Post, error) {
		if !post.IsCreator(creator) {
			return nil, ErrUnauthorizedCreator
		}

		if err := post.UpdateContent(content); err != nil {
			return nil, err
		}

		deletedMedia, err = post.RemoveMedia(deleteMediaID...)
		if err != nil {
			return nil, err
		}

		switch {
		case cfg.haveImages():
			for i := range cfg.images {
				imagePath := p.generateFilePath(post.CreatorUUID(), cfg.images[i].Filename)
				newImage, err := p.fac.NewMediaImage(imagePath, cfg.images[i].Size, post.CreatorUUID())
				if err != nil {
					return nil, err
				}
				// persist the image
				if err := p.filehdl.SaveFileMultipart(imagePath, cfg.images[i]); err != nil {
					return nil, err
				}
				addedMedia = append(addedMedia, newImage)

				if err := post.AddMedia(*newImage); err != nil {
					return nil, err
				}
			}
		case cfg.haveVideo():
			videoMedia, err := p.CreateAndSaveVideoFromMultipart(post.CreatorUUID(), cfg.video)
			if err != nil {
				return nil, err
			}
			addedMedia = append(addedMedia, videoMedia)
			if err := post.AddMedia(*videoMedia); err != nil {
				return nil, err
			}
		}

		return post, nil
	}); err != nil {
		// cleanup the added media
		for i := range addedMedia {
			p.filehdl.Cleanup(addedMedia[i].Path())
		}

		return err
	}

	for i := range deletedMedia {
		p.filehdl.Cleanup(deletedMedia[i].Path())
	}

	return nil
}
