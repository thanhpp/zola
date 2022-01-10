package esclient

type UserDataReq struct {
	UserID   string `json:"user_id"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	State    string `json:"state"`
}

type SearchReq struct {
	Keyword string `json:"keyword"`
	Index   string `json:"index,omitempty"`
	Count   string `json:"count,omitempty"`
}

type SearchResp []SearchRespData

type SearchRespData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Phonenumber string `json:"phonenumber"`
}
