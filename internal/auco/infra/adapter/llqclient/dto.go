package llqclient

type ValidateTokenResp struct {
	Exp  int    `json:"exp"`
	Jti  string `json:"jti"`
	Iat  int    `json:"iat"`
	Iss  string `json:"iss"`
	User struct {
		ID   string `json:"id"`
		Role string `json:"role"`
	} `json:"user"`
}

type GetUserInfoResp struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	State       string `json:"state"`
	Phone       string `json:"phone"`
	Role        string `json:"role"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Avatar      string `json:"avatar"`
	CoverImg    string `json:"cover_img"`
	LastOnline  int64  `json:"last_online"`
	CreatedAt   int64  `json:"created_at"`
}

type IsBlockResp struct {
	IsBlock bool `json:"is_block"`
}
