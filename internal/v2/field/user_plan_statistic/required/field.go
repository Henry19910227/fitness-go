package required

type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶id
}
type PlanIDField struct {
	PlanID int64 `json:"plan_id" gorm:"column:plan_id" binding:"required" example:"1"` // 計畫 id
}
type FinishWorkoutCountField struct {
	FinishWorkoutCount int `json:"finish_workout_count" gorm:"column:finish_workout_count" binding:"required" example:"10"` // 完成訓練數量(去除重複)
}
type DurationField struct {
	Duration int `json:"duration" gorm:"column:duration" binding:"required" example:"60"` // 總花費時間(秒)
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
