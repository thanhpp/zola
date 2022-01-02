package dto

import "github.com/thanhpp/zola/internal/laclongquan/domain/entity"

type GetUserResp struct {
	DefaultRespWithoutData
	Data struct {
		ID          string `json:"id"`
		Username    string `json:"username"`
		Description string `json:"description"`
		Avatar      string `json:"avatar"`
		CoverImage  string `json:"cover_image"`
		Link        string `json:"link"`
		Address     string `json:"address"`
		City        string `json:"city"`
		Country     string `json:"country"`
		IsFriend    string `json:"is_friend"`
		Listing     int64  `json:"listing"`
		Created     int64  `json:"created"`
	} `json:"data"`
}

type FormUserMediaFn func(user *entity.User) (avatar, coverImg string)

func (resp *GetUserResp) SetData(user *entity.User, friendCount int, isFriend bool, formURL FormUserMediaFn) {
	if resp == nil || user == nil || formURL == nil {
		return
	}

	resp.Data.ID = user.ID().String()
	resp.Data.Username = user.GetUsername()
	resp.Data.Description = user.GetDescription()
	avatarURL, coverImgURL := formURL(user)
	resp.Data.Avatar = avatarURL
	resp.Data.CoverImage = coverImgURL
	resp.Data.Link = user.GetLink()
	resp.Data.Address = user.GetAddress()
	resp.Data.City = user.GetCity()
	resp.Data.Country = user.GetCountry()
	resp.Data.IsFriend = boolTranslate(isFriend)
	resp.Data.Listing = int64(friendCount)
	resp.Data.Created = user.CreatedAtUnix()
}
