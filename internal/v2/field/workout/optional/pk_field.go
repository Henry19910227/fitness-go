package optional

type WorkoutIDField struct {
	WorkoutID *int64 `json:"workout_id,omitempty" gorm:"column:workout_id" binding:"omitempty" example:"1"` // 訓練 id
}
