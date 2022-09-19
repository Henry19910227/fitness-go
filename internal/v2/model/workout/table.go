package workout

import "github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"

type Table struct {
	optional.IDField
	optional.PlanIDField
	optional.NameField
	optional.EquipmentField
	optional.StartAudioField
	optional.EndAudioField
	optional.WorkoutSetCountField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "workouts"
}
