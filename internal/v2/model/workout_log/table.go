package workout_log

import "github.com/Henry19910227/fitness-go/internal/v2/field/workout_log/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.WorkoutIDField
	optional.DurationField
	optional.IntensityField
	optional.PlaceField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "workout_logs"
}
