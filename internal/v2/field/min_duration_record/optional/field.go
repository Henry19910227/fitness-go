package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` // 訓練 id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` // 用戶id
}
type ActionIDField struct {
	ActionID *int64 `json:"action_id,omitempty" gorm:"column:action_id" binding:"omitempty" example:"1"` //動作id
}
type DurationField struct {
	Duration *int `json:"duration,omitempty" gorm:"column:duration" binding:"omitempty" example:"10"` //時長
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
