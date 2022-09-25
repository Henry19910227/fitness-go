package required

type UserIDField struct {
	UserID int64 `json:"user_id" uri:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 帳戶id
}
