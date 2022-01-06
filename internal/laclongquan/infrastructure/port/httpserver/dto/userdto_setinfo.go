package dto

import "github.com/thanhpp/zola/internal/laclongquan/domain/entity"

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

func (resp *SetUserInfoResp) SetData(user *entity.User, formUserMediaURLFn FormUserMediaFn) {
	if resp == nil || user == nil || formUserMediaURLFn == nil {
		return
	}

	avatarURL, coverImgURL := formUserMediaURLFn(user)
	resp.Data.Avatar = avatarURL
	resp.Data.Cover = coverImgURL
	resp.Data.Link = user.GetLink()
	resp.Data.City = user.GetCity()
	resp.Data.Country = user.GetCountry()
}
