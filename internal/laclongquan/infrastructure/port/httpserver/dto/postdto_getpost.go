package dto

import (
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type ImageResponse struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type VideoResponse struct {
	URL   string `json:"url"`
	Thumb string `json:"thumb"`
}

type AuthorResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type GetPostResponseData struct {
	// post data-section
	ID           string          `json:"id"`
	Described    string          `json:"described"`
	CreatedAt    int64           `json:"created"`
	ModifiedAt   int64           `json:"modified"`
	LikeCount    int             `json:"like"`
	CommentCount int             `json:"comment"`
	Images       []ImageResponse `json:"images"`
	Video        VideoResponse   `json:"video"`

	// author section
	Author AuthorResponse `json:"author"`
	State  string         `json:"state"`

	// issuer section
	IsLiked    string `json:"is_liked"`
	IsBlocked  string `json:"is_blocked"`
	CanEdit    string `json:"can_edit"`
	CanComment string `json:"can_comment"`
}

type GetPostResponse struct {
	DefaultRespWithoutData
	Data GetPostResponseData `json:"data"`
}

type FormVideoThumbURLFunc func(post entity.Post, video entity.Media) string

type FormMediaURLFunc func(post entity.Post, media entity.Media) string

func (resp *GetPostResponse) SetData(
	getPostResult *application.GetPostResult,
	formMediaURLFn FormMediaURLFunc,
	formVideoThumbFn FormVideoThumbURLFunc,
	formUserMediaFN FormUserMediaFn) {
	if resp == nil || getPostResult == nil {
		return
	}

	// post section
	if getPostResult.Post != nil {
		resp.Data.ID = getPostResult.Post.ID()
		resp.Data.Described = getPostResult.Post.Content()
		resp.Data.CreatedAt = getPostResult.Post.CreatedAt()
		resp.Data.ModifiedAt = getPostResult.Post.UpdatedAt()

		for _, media := range getPostResult.Post.Media() {
			switch media.Type() {
			case entity.MediaTypeImage:
				resp.Data.Images = append(resp.Data.Images, ImageResponse{
					ID:  media.ID(),
					URL: formMediaURLFn(*getPostResult.Post, media),
				})

			case entity.MediaTypeVideo:
				resp.Data.Video = VideoResponse{
					URL:   formMediaURLFn(*getPostResult.Post, media),
					Thumb: formVideoThumbFn(*getPostResult.Post, media),
				}
			}
		}
	}

	resp.Data.LikeCount = getPostResult.LikeCount
	resp.Data.CommentCount = getPostResult.CommentCount

	// author section
	if getPostResult.Author != nil {
		avatarURL, _ := formUserMediaFN(getPostResult.Author)
		resp.Data.Author = AuthorResponse{
			ID:     getPostResult.Author.ID().String(),
			Name:   getPostResult.Author.Name(),
			Avatar: avatarURL,
		}
		resp.Data.State = getPostResult.Author.State().String()
	}

	// issuer section
	resp.Data.IsLiked = boolTranslate(getPostResult.IsLiked)
	resp.Data.IsBlocked = boolTranslate(false)
	resp.Data.CanEdit = boolTranslate(getPostResult.CanEdit)
	// logger.Debugf("post %s can comment: %v", getPostResult.Post.ID(), getPostResult.Post.GetCanComment())
	resp.Data.CanComment = boolTranslate(getPostResult.Post.GetCanComment())
}

func (resp *GetPostResponse) SetBlockedResponse() {
	if resp == nil {
		return
	}

	resp.Data.IsBlocked = boolTranslate(true)
}
