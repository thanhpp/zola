package dto

type BlockUserReq struct {
	BlockedUserID string `form:"user_id"`
	Type          string `form:"type"` // 0 = block or 1 = unblock
}
