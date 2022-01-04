package dto

import "github.com/thanhpp/zola/internal/laclongquan/application"

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

func (resp *GetRequestedFriendsResp) SetData(results []*application.GetRequestedFriendsRes, formUSerMediaURLFn FormUserMediaFn) {
	if resp == nil || results == nil || formUSerMediaURLFn == nil {
		return
	}

	for i := range results {
		avatarURL, _ := formUSerMediaURLFn(results[i].Friend)

		resp.Data.Friends = append(resp.Data.Friends, GetRequestedFriendsRespData{
			ID:       results[i].Friend.ID().String(),
			Username: results[i].Friend.GetUsername(),
			Avatar:   avatarURL,
			Created:  results[i].Relation.CreatedAtUnix(),
		})
	}
}
