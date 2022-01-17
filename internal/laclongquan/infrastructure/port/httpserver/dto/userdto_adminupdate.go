package dto

type AdminUpdateStateReq struct {
	State string `form:"state"`
}

type AdminUpdatePasswordReq struct {
	Password string `form:"password"`
}
