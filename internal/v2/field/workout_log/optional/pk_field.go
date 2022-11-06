package optional

type WorkoutLogIDField struct {
	WorkoutLogID *int64 `json:"workout_log_id,omitempty" uri:"workout_log_id" gorm:"column:workout_log_id" binding:"omitempty" example:"1"` // 訓練記錄id
}
