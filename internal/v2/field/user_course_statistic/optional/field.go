package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` //id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` // 用戶id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" binding:"omitempty" example:"10"` //課表id
}
type FinishWorkoutCountField struct {
	FinishWorkoutCount *int `json:"finish_workout_count,omitempty" gorm:"column:finish_workout_count" binding:"omitempty" example:"10"` //完成訓練總數量(去除重複並累加)
}
type TotalFinishWorkoutCountField struct {
	TotalFinishWorkoutCount *int `json:"total_finish_workout_count,omitempty" gorm:"column:total_finish_workout_count" binding:"omitempty" example:"10"` //訓練總量(可重複並累加)
}
type DurationField struct {
	Duration *int `json:"duration,omitempty" gorm:"column:duration" binding:"omitempty" example:"10"` //總花費時間(秒)
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 更新時間
}
