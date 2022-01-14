package dto

import "github.com/thanhpp/zola/internal/laclongquan/application"

type CreateCommentReq struct {
	Comment string `form:"comment"`
	Index   string `form:"index"`
	Count   string `form:"count"`
}

type CreateCommentResp struct {
	DefaultResp
	IsBlocked string `json:"is_blocked"`
}

func (resp *CreateCommentResp) SetIsBlocked(isBlocked bool) {
	if resp == nil {
		return
	}

	resp.IsBlocked = boolTranslate(isBlocked)
}

type UpdateCommentReq struct {
	NewContent string `form:"comment"`
}

type GetCommentResp struct {
	DefaultRespWithoutData
	Data []GetCommentRespData `json:"data"`
}

type GetCommentRespData struct {
	ID        string                   `json:"id"`
	Comment   string                   `json:"comment"`
	Created   string                   `json:"created"`
	Poster    GetCommentRespPosterData `json:"poster"`
	IsBlocked string                   `json:"is_blocked"`
}

type GetCommentRespPosterData struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (resp *GetCommentResp) SetData(res []*application.GetPostCommentRes, formUserMediaURL FormUserMediaFn) {
	if resp == nil || res == nil || formUserMediaURL == nil {
		return
	}

	resp.Data = make([]GetCommentRespData, 0, len(res))
	for i := range res {
		if res[i].IsBlocked {
			resp.Data = append(resp.Data, GetCommentRespData{
				IsBlocked: boolTranslate(res[i].IsBlocked),
			})
			continue
		}
		avatarURL, _ := formUserMediaURL(res[i].Comment.GetCreator())
		resp.Data = append(resp.Data, GetCommentRespData{
			ID:      res[i].Comment.IDString(),
			Comment: res[i].Comment.GetContent(),
			Created: res[i].Comment.CreatedAt.String(),
			Poster: GetCommentRespPosterData{
				ID:     res[i].Comment.GetCreator().ID().String(),
				Name:   res[i].Comment.GetCreator().GetName(),
				Avatar: avatarURL,
			},
			IsBlocked: boolTranslate(res[i].IsBlocked),
		})
	}
}
