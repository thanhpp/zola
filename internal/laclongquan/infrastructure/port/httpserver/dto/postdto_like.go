package dto

type LikePostResp struct {
	DefaultResp
	Data struct {
		Like int `json:"like"`
	} `json:"data"`
}

func (r *LikePostResp) SetData(like int) {
	r.Data.Like = like
}
