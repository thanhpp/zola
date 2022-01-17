package dto

type AdminCreateUserReq struct {
	Phone       string `form:"phone"`
	Pass        string `form:"pass"`
	Name        string `form:"name"`
	Username    string `form:"username"`
	Description string `form:"description"`
	Address     string `form:"address"`
	City        string `form:"city"`
	Country     string `form:"country"`
}

type AdminCreateUserResp struct {
	DefaultRespWithoutData
	Data struct {
		UserID string `json:"user_id"`
	} `json:"data"`
}

func (resp *AdminCreateUserResp) SetData(userID string) {
	if resp == nil {
		return
	}
	resp.Data.UserID = userID
}
