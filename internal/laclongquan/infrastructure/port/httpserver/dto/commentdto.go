package dto

import (
	"github.com/thanhpp/zola/internal/laclongquan/application"
)

type CreateCommentReq struct {
	Comment string `form:"comment"`
	Index   string `form:"index"`
	Count   string `form:"count"`
}

type CreateCommentResp struct {
	DefaultResp
	Data      []*GetCommentRespData `json:"data"`
	IsBlocked string                `json:"is_blocked"`
}

func (resp *CreateCommentResp) SetIsBlocked(isBlocked bool) {
	if resp == nil {
		return
	}

	resp.IsBlocked = boolTranslate(isBlocked)
}

func (resp *CreateCommentResp) SetData(res []*application.GetPostCommentRes, formUserMediaURL FormUserMediaFn) {
	if resp == nil || res == nil || formUserMediaURL == nil {
		return
	}

	resp.IsBlocked = boolTranslate(false)
	resp.Data = make([]*GetCommentRespData, 0, len(res))
	for i := range res {
		respData := new(GetCommentRespData)
		respData.setData(res[i], formUserMediaURL)
		resp.Data = append(resp.Data, respData)
	}
}

type UpdateCommentReq struct {
	NewContent string `form:"comment"`
}

type GetCommentResp struct {
	DefaultRespWithoutData
	Data      []GetCommentRespData `json:"data"`
	IsBlocked string               `json:"is_blocked"`
}

type GetCommentRespData struct {
	ID      string                   `json:"id"`
	Comment string                   `json:"comment"`
	Created string                   `json:"created"`
	Poster  GetCommentRespPosterData `json:"poster"`
}

type GetCommentRespPosterData struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (resp *GetCommentResp) SetIsBlocked() {
	if resp == nil {
		return
	}

	resp.IsBlocked = boolTranslate(true)
}

func (data *GetCommentRespData) setData(res *application.GetPostCommentRes, formUserMediaURL FormUserMediaFn) {
	if data == nil || res == nil || formUserMediaURL == nil {
		return
	}

	avatarURL, _ := formUserMediaURL(res.Comment.GetCreator())
	data.ID = res.Comment.IDString()
	data.Comment = res.Comment.GetContent()
	data.Created = res.Comment.CreatedAt.String()
	data.Poster = GetCommentRespPosterData{
		ID:     res.Comment.GetCreator().ID().String(),
		Name:   res.Comment.GetCreator().GetName(),
		Avatar: avatarURL,
	}
}

func (resp *GetCommentResp) SetData(res []*application.GetPostCommentRes, formUserMediaURL FormUserMediaFn) {
	if resp == nil || res == nil || formUserMediaURL == nil {
		return
	}

	resp.IsBlocked = boolTranslate(false)
	resp.Data = make([]GetCommentRespData, len(res))
	for i := range res {
		resp.Data[i].setData(res[i], formUserMediaURL)
	}
}
