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

type GetPostResponse struct {
	DefaultResp
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

type FormMediaURLFunc func(post entity.Post, media entity.Media) string

func (resp *GetPostResponse) SetData(getPostResult *application.GetPostResult, formMediaURLFn FormMediaURLFunc) {
	if resp == nil || getPostResult == nil {
		return
	}

	// post section
	if getPostResult.Post != nil {
		resp.ID = getPostResult.Post.ID()
		resp.Described = getPostResult.Post.Content()
		resp.CreatedAt = getPostResult.Post.CreatedAt()
		resp.ModifiedAt = getPostResult.Post.UpdatedAt()

		for _, media := range getPostResult.Post.Media() {
			switch media.Type() {
			case entity.MediaTypeImage:
				resp.Images = append(resp.Images, ImageResponse{
					ID:  media.ID(),
					URL: formMediaURLFn(*getPostResult.Post, media),
				})

			case entity.MediaTypeVideo:
				resp.Video = VideoResponse{
					URL: formMediaURLFn(*getPostResult.Post, media),
					// FIXME: thumb
				}
			}
		}
	}

	resp.LikeCount = getPostResult.LikeCount
	resp.CommentCount = getPostResult.CommentCount

	// author section
	if getPostResult.Author != nil {
		resp.Author = AuthorResponse{
			ID:   getPostResult.Author.ID().String(),
			Name: getPostResult.Author.Name(),
			// FIXME: avatar
		}
		resp.State = getPostResult.Author.State().String()
	}

	// issuer section
	resp.IsLiked = boolTranslate(getPostResult.IsLiked)
	resp.IsBlocked = boolTranslate(false)
	resp.CanEdit = boolTranslate(getPostResult.CanEdit)
	resp.CanComment = boolTranslate(getPostResult.CanComment)
}

func (resp *GetPostResponse) SetBlockedResponse() {
	if resp == nil {
		return
	}

	resp.IsBlocked = boolTranslate(true)
}
