package dto

import (
	"strconv"

	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type SearchReq struct {
	Keyword string `form:"keyword"`
}

type SearchResp struct {
	DefaultRespWithoutData
	Data []SearchRespPostData `json:"data"`
}

func (resp *SearchResp) SetData(res []*application.GetPostResult, formMediaURL FormMediaURLFunc, formVideoThumbURL FormVideoThumbURLFunc, formUserMediaURL FormUserMediaFn) {
	if resp == nil || res == nil || formMediaURL == nil || formVideoThumbURL == nil || formUserMediaURL == nil {
		return
	}
	resp.Data = make([]SearchRespPostData, 0, len(res))
	for i := range res {
		avatarURL, _ := formUserMediaURL(res[i].Author)
		dataElem := SearchRespPostData{
			ID:        res[i].Post.ID(),
			Like:      strconv.Itoa(res[i].LikeCount),
			Comment:   strconv.Itoa(res[i].CommentCount),
			Described: res[i].Post.Content(),
			Author: SearchRespAuthorData{
				ID:       res[i].Author.ID().String(),
				Name:     res[i].Author.Name(),
				Username: res[i].Author.GetUsername(),
				Avatar:   avatarURL,
			},
		}

		for _, media := range res[i].Post.Media() {
			switch media.Type() {
			case entity.MediaTypeImage:
				dataElem.Image = append(dataElem.Image, formMediaURL(*res[i].Post, media))

			case entity.MediaTypeVideo:
				dataElem.Video = SearchRespPostVideoData{
					URL:   formMediaURL(*res[i].Post, media),
					Thumb: formVideoThumbURL(*res[i].Post, media),
				}
			}
		}

		resp.Data = append(resp.Data, dataElem)
	}
}

type SearchRespPostData struct {
	ID        string                  `json:"id"`
	Image     []string                `json:"image"`
	Video     SearchRespPostVideoData `json:"video"`
	Like      string                  `json:"like"`
	Comment   string                  `json:"comment"`
	Author    SearchRespAuthorData    `json:"author"`
	Described string                  `json:"described"`
}

type SearchRespPostVideoData struct {
	URL   string `json:"url"`
	Thumb string `json:"thumb"`
}

type SearchRespAuthorData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}
