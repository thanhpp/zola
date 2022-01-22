package application

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/adapter/esclient"
	"github.com/thanhpp/zola/pkg/logger"
)

type PostHandler struct {
	fac          entity.PostFactory
	repo         repository.PostRepository
	likeRepo     repository.LikeRepository
	commentRepo  repository.CommentRepository
	relationRepo repository.RelationRepository
	userRepo     repository.UserRepository
	filehdl      *FileHandler
	saveDir      string
	esClient     *esclient.EsClient
}

func NewPostHandler(repo repository.PostRepository, saveDir string,
	likeRepo repository.LikeRepository,
	commentRepo repository.CommentRepository,
	relationRepo repository.RelationRepository,
	userRepo repository.UserRepository,
	esClient *esclient.EsClient,
) PostHandler {
	return PostHandler{
		fac:          entity.NewPostFactory(),
		repo:         repo,
		likeRepo:     likeRepo,
		commentRepo:  commentRepo,
		relationRepo: relationRepo,
		userRepo:     userRepo,
		filehdl:      new(FileHandler),
		saveDir:      saveDir,
		esClient:     esClient,
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

func (p PostHandler) SyncAllPostToEs() {
	// get all post from es
	postIDs, err := p.esClient.GetAllPost()
	if err != nil {
		logger.Errorf("error getting all post from elasticsearch %v", err)
	}
	logger.Debugf("got %d post from elasticsearch", len(postIDs))

	// get all post from database
	var nullTime time.Time
	dbPosts, _, err := p.repo.GetListPostForAdmin(context.Background(), "", nullTime, 0, 1000)
	if err != nil {
		logger.Errorf("error getting all post from database %v", err)
		return
	}

	postIDsMap := make(map[string]struct{})
	for _, dbPost := range dbPosts {
		postIDsMap[dbPost.ID()] = struct{}{}
		author, err := p.userRepo.GetByID(context.Background(), dbPost.Creator())
		if err != nil {
			logger.Errorf("error getting user %s from database %v", dbPost.Creator(), err)
			continue
		}
		if err := p.esClient.CreateUpdatePost(dbPost, author.Name()); err != nil {
			logger.Errorf("error syncing post %s to elasticsearch %v", dbPost.ID(), err)
			continue
		}
		logger.Infof("synced post %s to elasticsearch", dbPost.ID())
	}

	for i := range postIDs {
		if _, ok := postIDsMap[postIDs[i]]; !ok {
			if err := p.esClient.DeletePost(postIDs[i]); err != nil {
				logger.Errorf("error deleting post %s from elasticsearch %v", postIDs[i], err)
				continue
			}
			logger.Infof("deleted post %s from elasticsearch", postIDs[i])
		}
		// delete(postIDsMap, postIDs[i])
	}
}
