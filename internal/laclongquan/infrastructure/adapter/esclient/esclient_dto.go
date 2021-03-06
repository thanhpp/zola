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

type PostDataReq struct {
	ID        string         `json:"id"`
	Described string         `json:"described"`
	Author    PostAuthorData `json:"author"`
	Created   string         `json:"created"`
	Modified  string         `json:"modified"`
}

type PostAuthorData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SearchPostResp []SearchPostRespData

type SearchPostRespData struct {
	ID        string `json:"id"`
	Described string `json:"described"`
	Author    struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"author"`
	Created string `json:"created"`
}
