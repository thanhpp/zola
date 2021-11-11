package dto

type CreateReportReq struct {
	ID      string `form:"id"`
	Subject string `form:"subject"`
	Details string `form:"details"`
}
