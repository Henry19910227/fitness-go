package workout_log

import (
	workoutOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_log/optional"
)

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

type WorkoutTable struct {
	workoutOptional.IDField
	workoutOptional.PlanIDField
	workoutOptional.NameField
	workoutOptional.EquipmentField
	workoutOptional.StartAudioField
	workoutOptional.EndAudioField
	workoutOptional.WorkoutSetCountField
	workoutOptional.CreateAtField
	workoutOptional.UpdateAtField
}

func (WorkoutTable) TableName() string {
	return "workouts"
}