package workout_set

import "github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/optional"

type Table struct {
	optional.IDField
	optional.WorkoutIDField
	optional.ActionIDField
	optional.TypeField
	optional.AutoNextField
	optional.StartAudioField
	optional.ProgressAudioField
	optional.RemarkField
	optional.WeightField
	optional.RepsField
	optional.DistanceField
	optional.DurationField
	optional.InclineField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "workout_sets"
}
