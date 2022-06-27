package body_image

type UserIDRequired struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` //用戶id
}
