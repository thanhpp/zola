package repository

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type ReportRepository interface {
	Create(ctx context.Context, report *entity.Report) error
}
