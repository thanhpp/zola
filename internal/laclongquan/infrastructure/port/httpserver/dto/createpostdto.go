package dto

type CreatePostReq struct {
	Described string `form:"described"`
}

type CreatePostResp struct {
	DefaultResp
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (resp *CreatePostResp) SetData(id string) {
	resp.Data.ID = id
}
