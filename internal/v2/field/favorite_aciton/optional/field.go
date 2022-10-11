package optional

type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` // 用戶 id
}
type ActionIDField struct {
	ActionID *int64 `json:"action_id,omitempty" uri:"action_id" gorm:"column:action_id" binding:"omitempty" example:"10"` // 動作id
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 創建時間
}
