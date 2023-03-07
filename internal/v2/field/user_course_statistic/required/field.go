package required

type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶id
}
type CourseIDField struct {
	CourseID int64 `json:"course_id" gorm:"column:course_id" binding:"required" example:"10"` //課表id
}
type FinishWorkoutCountField struct {
	FinishWorkoutCount int `json:"finish_workout_count" gorm:"column:finish_workout_count" binding:"required" example:"10"` //完成訓練總數量(去除重複並累加)
}
type TotalFinishWorkoutCountField struct {
	TotalFinishWorkoutCount int `json:"total_finish_workout_count" gorm:"column:total_finish_workout_count" binding:"required" example:"10"` //訓練總量(可重複並累加)
}
type DurationField struct {
	Duration int `json:"duration" gorm:"column:duration" binding:"required" example:"10"` //總花費時間(秒)
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-12 00:00:00"` // 更新時間
}
