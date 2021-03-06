package dto

import (
	"strconv"

	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type GetUserResp struct {
	DefaultRespWithoutData
	Data struct {
		ID          string `json:"id"`
		Phone       string `json:"phone"`
		Username    string `json:"username"`
		Description string `json:"description"`
		Name        string `json:"name"`
		Avatar      string `json:"avatar"`
		CoverImage  string `json:"cover_image"`
		Link        string `json:"link"`
		Address     string `json:"address"`
		City        string `json:"city"`
		Country     string `json:"country"`
		IsFriend    string `json:"is_friend"`
		State       string `json:"state"`
		IsActive    string `json:"is_active"`
		IsOnline    string `json:"is_online"`
		Listing     string `json:"listing"`
		Created     string `json:"created"`
	} `json:"data"`
}

type FormUserMediaFn func(user *entity.User) (avatar, coverImg string)

func (resp *GetUserResp) SetData(user *entity.User, friendCount int, isFriend bool, formURL FormUserMediaFn) {
	if resp == nil || user == nil || formURL == nil {
		return
	}

	resp.Data.ID = user.ID().String()
	resp.Data.Phone = user.Account().Phone
	resp.Data.Username = user.GetUsername()
	resp.Data.Description = user.GetDescription()
	resp.Data.Name = user.GetName()
	avatarURL, coverImgURL := formURL(user)
	resp.Data.Avatar = avatarURL
	resp.Data.CoverImage = coverImgURL
	resp.Data.Link = user.GetLink()
	resp.Data.Address = user.GetAddress()
	resp.Data.City = user.GetCity()
	resp.Data.Country = user.GetCountry()
	resp.Data.IsFriend = boolTranslate(isFriend)
	resp.Data.State = user.State().String()
	resp.Data.IsActive = boolTranslate(user.IsActive())
	resp.Data.IsOnline = boolTranslate(user.IsOnline())
	resp.Data.Listing = strconv.Itoa(friendCount)
	resp.Data.Created = strconv.Itoa(int(user.CreatedAtUnix()))
}

type UserData struct {
	UserID    string `json:"user_id"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	State     string `json:"state"`
	IsActive  string `json:"is_active"`
	LastLogin string `json:"last_login"`
}

type GetUserListResp struct {
	DefaultRespWithoutData
	Data struct {
		Users []UserData `json:"users"`
		Total string     `json:"total"`
	} `json:"data"`
}

func (resp *GetUserListResp) SetData(res *application.GetUserRes, formUserMediaURLFn FormUserMediaFn) {
	if resp == nil || res == nil || formUserMediaURLFn == nil {
		return
	}

	resp.Data.Total = strconv.Itoa(res.Total)
	resp.Data.Users = make([]UserData, 0, len(res.UserList))
	for _, user := range res.UserList {
		avatarURL, _ := formUserMediaURLFn(user)
		resp.Data.Users = append(resp.Data.Users, UserData{
			UserID:    user.ID().String(),
			Phone:     user.Account().Phone,
			Username:  user.GetUsername(),
			Name:      user.Name(),
			Avatar:    avatarURL,
			IsActive:  boolTranslate(user.IsActive()),
			State:     user.State().String(),
			LastLogin: "",
		})
	}
}
