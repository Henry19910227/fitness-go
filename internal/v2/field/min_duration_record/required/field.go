package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` // 訓練 id
}
type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶id
}
type ActionIDField struct {
	ActionID int64 `json:"action_id" gorm:"column:action_id" binding:"required" example:"1"` //動作id
}
type DurationField struct {
	Duration int `json:"duration" gorm:"column:duration" binding:"required" example:"10"` //時長
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
