package workout_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log"
)

type Output struct {
	Table
}

func (Output) TableName() string {
	return "workout_logs"
}

// APICreateUserWorkoutLogOutput /v2/user/workout/{workout_id}/workout_log [POST]
type APICreateUserWorkoutLogOutput struct {
	base.Output
	Data *APICreateUserWorkoutLogData `json:"data,omitempty"`
}
type APICreateUserWorkoutLogData []*struct {
	APICreateUserWorkoutLogItem
}
type APICreateUserWorkoutLogItem struct {
	workout_set_log.WorkoutSetIDField
	workout_set_log.WeightField
	workout_set_log.DistanceField
	workout_set_log.InclineField
	workout_set_log.RepsField
	workout_set_log.DurationField
	NewRecord  *int `json:"new_record" example:"1"` //是否是新紀錄(0:否/1:是)
	WorkoutSet *struct {
		workout_set.IDField
		workout_set.TypeField
		workout_set.WeightField
		workout_set.RepsField
		workout_set.DistanceField
		workout_set.DurationField
		workout_set.InclineField
		Action *struct {
			action.IDField
			action.NameField
			action.SourceField
			action.TypeField
		} `json:"action,omitempty"`
	} `json:"workout_set,omitempty"`
}
