package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

func (p PostHandler) CreatePostWithMultipart(ctx context.Context, creator uuid.UUID, content string, opts ...MultipartOption) (*entity.Post, error) {
	var cfg = new(multipartConfig)
	for _, apply := range opts {
		apply(cfg)
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	var (
		post *entity.Post
		err  error
	)
	switch {
	case cfg.haveImages():
		var (
			imageList  = make([]entity.Media, 0, len(cfg.images))
			imagePaths = make([]string, 0, len(cfg.images))
		)
		for _, image := range cfg.images {
			imagePath := p.generateFilePath(creator, image.Filename)
			imagePaths = append(imagePaths, imagePath)

			imageEntity, err := p.fac.NewMediaImage(imagePath, image.Size, creator)
			if err != nil {
				return nil, err
			}
			imageList = append(imageList, *imageEntity)

			imageFile, err := image.Open()
			if err != nil {
				return nil, err
			}
			defer imageFile.Close()

			if err := p.filehdl.SaveFile(imagePath, imageFile); err != nil {
				return nil, err
			}
		}

		post, err = p.fac.NewPostWithImages(creator, content, imageList...)
		if err != nil {
			p.filehdl.Cleanup(imagePaths...)
			return nil, err
		}

	case cfg.haveVideo():
		videoPath := p.generateFilePath(creator, cfg.video.Filename)
		videoMedia, err := p.fac.NewMediaVideo(videoPath, cfg.video.Size, creator)
		if err != nil {
			return nil, err
		}

		// the video file must be created to see it's duration
		videoFile, err := cfg.video.Open()
		if err != nil {
			return nil, err
		}
		defer videoFile.Close()
		if err := p.filehdl.SaveFile(videoPath, videoFile); err != nil {
			return nil, err
		}

		duration, err := p.filehdl.GetVideoDuration(videoPath)
		if err != nil {
			p.filehdl.Cleanup(videoPath)
			return nil, err
		}

		if err := videoMedia.DurationCheck(duration); err != nil {
			p.filehdl.Cleanup(videoPath)
			return nil, err
		}

		post, err = p.fac.NewPostWithVideo(creator, content, *videoMedia)
		if err != nil {
			p.filehdl.Cleanup(videoPath)
			return nil, err
		}

	default:
		post, err = p.fac.NewPost(creator, content)
		if err != nil {
			return nil, err
		}
	}

	if err := p.repo.Create(ctx, post); err != nil {
		mediaList := post.Media()
		for i := range mediaList {
			p.filehdl.Cleanup(mediaList[i].Path())
		}
		return nil, err
	}

	return post, nil
}
