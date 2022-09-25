package required

type WorkoutSetIDField struct {
	WorkoutSetID int64 `json:"workout_set_id" uri:"workout_set_id" gorm:"column:workout_set_id" binding:"omitempty" example:"2"` // 訓練組 id
}
