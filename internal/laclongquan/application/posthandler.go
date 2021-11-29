package application

import (
	"fmt"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

type PostHandler struct {
	fac          entity.PostFactory
	repo         repository.PostRepository
	commentRepo  repository.CommentRepository
	relationRepo repository.RelationRepository
	userRepo     repository.UserRepository
	filehdl      *FileHandler
	saveDir      string
}

func NewPostHandler(repo repository.PostRepository, saveDir string,
	commentRepo repository.CommentRepository,
	relationRepo repository.RelationRepository,
	userRepo repository.UserRepository,
) PostHandler {
	return PostHandler{
		fac:          entity.NewPostFactory(),
		repo:         repo,
		commentRepo:  commentRepo,
		relationRepo: relationRepo,
		userRepo:     userRepo,
		filehdl:      new(FileHandler),
		saveDir:      saveDir,
	}
}

func (p PostHandler) generateFilePath(creator uuid.UUID, filename string) string {
	return fmt.Sprintf("%s/%s/%s-%s", p.saveDir, creator.String(), randomString(15), filename)
}

func (p PostHandler) CreateAndSaveVideoFromMultipart(creator uuid.UUID, fileHeader *multipart.FileHeader) (*entity.Media, error) {
	if fileHeader == nil {
		return nil, nil
	}

	// generate path
	videoPath := p.generateFilePath(creator, fileHeader.Filename)

	// video media
	videoMedia, err := p.fac.NewMediaVideo(videoPath, fileHeader.Size, creator)
	if err != nil {
		return nil, err
	}

	// save the video
	err = p.filehdl.SaveFileMultipart(videoPath, fileHeader)
	if err != nil {
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

	return videoMedia, nil
}
