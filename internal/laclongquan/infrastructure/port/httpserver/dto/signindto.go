package dto

type SignInReq struct {
	PhoneNumber string `form:"phonenumber"`
	Password    string `form:"password"`
}

type SignInResp struct {
	DefaultResp
	Data string `json:"data"`
}

func (resp *SignInResp) SetData(data string) {
	resp.Data = data
}
