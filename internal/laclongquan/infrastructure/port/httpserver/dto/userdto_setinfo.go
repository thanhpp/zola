package dto

type SetUserInfoReq struct {
	Username    string `form:"username"`
	Description string `form:"description"`
	Address     string `form:"address"`
	City        string `form:"city"`
	Country     string `form:"country"`
	Avatar      string `form:"avatar"`
	CoverImage  string `form:"cover_image"`
	Link        string `form:"link"`
}

type SetUserInfoResp struct {
	DefaultRespWithoutData
	Data struct {
		Avatar  string `json:"avatar"`
		Cover   string `json:"cover_image"`
		Link    string `json:"link"`
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"data"`
}
