package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

type ReportHandler struct {
	fac        entity.ReportFactory
	reportRepo repository.ReportRepository
	postRepo   repository.PostRepository
}

func NewReportHandler(fac entity.ReportFactory, reportRepo repository.ReportRepository, postRepo repository.PostRepository) ReportHandler {
	return ReportHandler{
		fac:        fac,
		reportRepo: reportRepo,
		postRepo:   postRepo,
	}
}

func (r ReportHandler) CreateReport(ctx context.Context, postID, creatorID string, subjectID int, detail string) (*entity.Report, error) {
	// check post
	_, err := r.postRepo.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	// create report
	report, err := r.fac.NewReport(postID, creatorID, subjectID, detail)
	if err != nil {
		return nil, err
	}

	if err := r.reportRepo.Create(ctx, &report); err != nil {
		return nil, err
	}

	return &report, nil
}
