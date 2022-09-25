package required

type WorkoutIDField struct {
	WorkoutID int64 `json:"workout_id" gorm:"column:workout_id" binding:"required" example:"1"` // 訓練 id
}
