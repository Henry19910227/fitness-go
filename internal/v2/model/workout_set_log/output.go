package workout_set_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
)

type Output struct {
	Table
	WorkoutSet *workout_set.Output `json:"workout_set,omitempty" gorm:"foreignKey:id;references:workout_set_id"` // 訓練log列表
}

func (Output) TableName() string {
	return "workout_set_logs"
}

func (o Output) WorkoutSetOnSafe() workout_set.Output {
	if o.WorkoutSet != nil {
		return *o.WorkoutSet
	}
	return workout_set.Output{}
}