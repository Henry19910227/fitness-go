package user_subscribe_info

type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" binding:"omitempty" example:"10001"` //用戶id
}
