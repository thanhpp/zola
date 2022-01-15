package dto

import (
	"strconv"

	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type GetListPostReq struct {
	LastID string `form:"last_id"`
}

type GetListPostResp struct {
	DefaultRespWithoutData
	Data struct {
		Posts    []GetListPostRespPostData `json:"posts"`
		NewItems string                    `json:"new_items"`
		LastID   string                    `json:"last_id"`
	} `json:"data"`
}

func (resp *GetListPostResp) SetData(res *application.GetListPostRes, lastID string, formMediaURL FormMediaURLFunc, formVideoThumbURL FormVideoThumbURLFunc, formUserMediaURL FormUserMediaFn) {
	if resp == nil || res == nil || formMediaURL == nil || formVideoThumbURL == nil || formUserMediaURL == nil {
		return
	}
	resp.Data.LastID = lastID
	resp.Data.NewItems = strconv.Itoa(res.NewItems)
	// logger.Debugf("res.Posts %d", len(res.Posts))
	for i := range res.Posts {
		var (
			avatarURL, _ = formUserMediaURL(res.Posts[i].Creator)
			imageList    []string
			videoData    GetListPostRespVideoData
		)
		for _, media := range res.Posts[i].Post.Media() {
			switch media.Type() {
			case entity.MediaTypeImage:
				imageList = append(imageList, formMediaURL(*res.Posts[i].Post, media))

			case entity.MediaTypeVideo:
				videoData = GetListPostRespVideoData{
					URL:   formMediaURL(*res.Posts[i].Post, media),
					Thumb: formVideoThumbURL(*res.Posts[i].Post, media),
				}
			}
		}
		resp.Data.Posts = append(resp.Data.Posts, GetListPostRespPostData{
			ID:         res.Posts[i].Post.ID(),
			Described:  res.Posts[i].Post.Content(),
			Created:    strconv.FormatInt(res.Posts[i].Post.CreatedAt(), 10),
			Like:       strconv.Itoa(res.Posts[i].LikeCount),
			Comment:    strconv.Itoa(res.Posts[i].CommentCount),
			IsLiked:    boolTranslate(res.Posts[i].IsLiked),
			IsBlocked:  boolTranslate(false),
			CanComment: boolTranslate(res.Posts[i].Post.CanComment),
			CanEdit:    boolTranslate(res.Posts[i].CanEdit),
			Image:      imageList,
			Video:      videoData,
			Author: GetListPostRespAuthorData{
				ID:       res.Posts[i].Creator.ID().String(),
				Username: res.Posts[i].Creator.GetUsername(),
				Avatar:   avatarURL,
				Online:   boolTranslate(res.Posts[i].Creator.IsOnline()),
			},
		})
	}
}

type GetListPostRespPostData struct {
	ID string `json:"id"`
	// Name       string                    `json:"name"`
	Described  string                    `json:"described"`
	Created    string                    `json:"created"`
	Like       string                    `json:"like"`
	Comment    string                    `json:"comment"`
	IsLiked    string                    `json:"is_liked"`
	IsBlocked  string                    `json:"is_blocked"`
	CanComment string                    `json:"can_comment"`
	CanEdit    string                    `json:"can_edit"`
	Image      []string                  `json:"image"`
	Video      GetListPostRespVideoData  `json:"video"`
	Author     GetListPostRespAuthorData `json:"author"`
}

type GetListPostRespVideoData struct {
	URL   string `json:"url"`
	Thumb string `json:"thumb"`
}

type GetListPostRespAuthorData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Online   string `json:"online"`
}
