package body_image

type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" binding:"omitempty" example:"10001"` //用戶id
}
