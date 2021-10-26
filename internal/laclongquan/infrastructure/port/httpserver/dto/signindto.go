package dto

type SignInReq struct {
	PhoneNumber string `form:"phonenumber"`
	Password    string `form:"password"`
}
