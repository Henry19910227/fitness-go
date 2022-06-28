package body_image

type IDOptional struct {
	ID *int64 `json:"id,omitempty" binding:"omitempty" example:"1"` //主鍵id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" binding:"omitempty" example:"10001"` //用戶id
}
