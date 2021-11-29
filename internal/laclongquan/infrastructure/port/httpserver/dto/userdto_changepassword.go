package dto

type ChangePasswordReq struct {
	Password    string `form:"password"`
	NewPassword string `form:"new_password"`
}
