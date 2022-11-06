package workout_set_log

import "github.com/Henry19910227/fitness-go/internal/v2/field/workout_set_log/optional"

type Table struct {
	optional.IDField
	optional.WorkoutLogIDField
	optional.WorkoutSetIDField
	optional.WeightField
	optional.RepsField
	optional.DistanceField
	optional.DurationField
	optional.InclineField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "workout_set_logs"
}
