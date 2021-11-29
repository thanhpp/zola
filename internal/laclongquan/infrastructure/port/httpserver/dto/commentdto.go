package dto

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
	if isBlocked {
		resp.IsBlocked = "1"
		return
	}
	resp.IsBlocked = "0"
	return
}
