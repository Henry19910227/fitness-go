package required

type WorkoutLogIDField struct {
	WorkoutLogID int64 `json:"workout_log_id" uri:"workout_log_id" gorm:"column:workout_log_id" binding:"required" example:"1"` // 訓練記錄id
}

