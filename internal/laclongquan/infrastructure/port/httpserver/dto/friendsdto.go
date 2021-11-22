package dto

type UpdateFriendRequestReq struct {
	IsAccept string `form:"is_accept"`
}

func (req UpdateFriendRequestReq) IsAcceptCode() bool {
	return req.IsAccept == "1"
}

func (req UpdateFriendRequestReq) IsRejectCode() bool {
	return req.IsAccept == "0"
}
