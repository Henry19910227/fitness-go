package optional

type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" form:"user_id" uri:"user_id" gorm:"column:user_id" binding:"omitempty" example:"10001"` // 帳戶id
}
