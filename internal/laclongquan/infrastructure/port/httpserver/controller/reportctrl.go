package controller

import "github.com/thanhpp/zola/internal/laclongquan/application"

type ReportController struct {
	app application.ReportHandler
}

func NewReportCtrl(app application.ReportHandler) *ReportController {
	return &ReportController{
		app: app,
	}
}
