package dto

import (
	"strconv"

	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type UpdateFriendRequestReq struct {
	IsAccept string `form:"is_accept"`
}

func (req UpdateFriendRequestReq) IsAcceptCode() bool {
	return req.IsAccept == "1"
}

func (req UpdateFriendRequestReq) IsRejectCode() bool {
	return req.IsAccept == "0"
}

type GetRequestedFriendsResp struct {
	DefaultRespWithoutData
	Data struct {
		Friends []GetRequestedFriendsRespData `json:"friends"`
	} `json:"data"`
}

type GetRequestedFriendsRespData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Created  int64  `json:"created"`
}

func (resp *GetRequestedFriendsResp) SetData(results []*application.GetRequestedFriendsRes, formUserMediaURLFn FormUserMediaFn) {
	if resp == nil || results == nil || formUserMediaURLFn == nil {
		return
	}

	for i := range results {
		avatarURL, _ := formUserMediaURLFn(results[i].Friend)

		resp.Data.Friends = append(resp.Data.Friends, GetRequestedFriendsRespData{
			ID:       results[i].Friend.ID().String(),
			Username: results[i].Friend.GetUsername(),
			Avatar:   avatarURL,
			Created:  results[i].Relation.CreatedAtUnix(),
		})
	}
}

type GetUserFriendsResp struct {
	DefaultRespWithoutData
	Data struct {
		Friends []GetUserFriendsRespData `json:"friends"`
		Total   string                   `json:"total"`
	} `json:"data"`
}

type GetUserFriendsRespData struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Status   string `json:"status"`
}

func (resp *GetUserFriendsResp) SetData(users []*entity.User, total int, formUserMediaURLFn FormUserMediaFn) {
	if resp == nil || formUserMediaURLFn == nil {
		return
	}

	for i := range users {
		avatarURL, _ := formUserMediaURLFn(users[i])

		status := "offline"
		if users[i].IsOnline() {
			status = "online"
		}

		resp.Data.Friends = append(resp.Data.Friends, GetUserFriendsRespData{
			UserID:   users[i].ID().String(),
			UserName: users[i].GetUsername(),
			Name:     users[i].GetName(),
			Avatar:   avatarURL,
			Status:   status,
		})
	}

	resp.Data.Total = strconv.Itoa(total)
}
