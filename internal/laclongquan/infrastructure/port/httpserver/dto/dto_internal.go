package dto

import "github.com/thanhpp/zola/internal/laclongquan/domain/entity"

type InternalGetUserResp struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	State       string `json:"state"`
	Phone       string `json:"phone"`
	Role        string `json:"role"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Avatar      string `json:"avatar"`
	CoverImg    string `json:"cover_img"`
	LastOnline  int64  `json:"last_online"`
	CreatedAt   int64  `json:"created_at"`
}

func (resp *InternalGetUserResp) SetData(user *entity.User, formUserMedia FormUserMediaFn) {
	if resp == nil || user == nil || formUserMedia == nil {
		return
	}
	avatarURL, coverURL := formUserMedia(user)
	resp.ID = user.ID().String()
	resp.Username = user.GetUsername()
	resp.Description = user.GetDescription()
	resp.Name = user.GetName()
	resp.Link = user.GetLink()
	resp.State = user.State().String()
	resp.Phone = user.Account().Phone
	resp.Role = user.Role()
	resp.Address = user.GetAddress()
	resp.City = user.GetCity()
	resp.Country = user.GetCountry()
	resp.Avatar = avatarURL
	resp.CoverImg = coverURL
	resp.LastOnline = user.GetLastOnline().Unix()
	resp.CreatedAt = user.CreatedAtUnix()
}

type InternalIsBlockResp struct {
	IsBlock bool `json:"is_block"`
}

func (resp *InternalIsBlockResp) SetData(isBlock bool) {
	if resp == nil {
		return
	}
	resp.IsBlock = isBlock
}
