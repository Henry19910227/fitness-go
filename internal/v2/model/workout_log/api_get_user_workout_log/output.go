package api_get_user_workout_log

import (
	actionOptional "github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	workoutOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_log/optional"
	workoutSetOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/optional"
	workoutSetLogOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout_set_log/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/user/workout_log/{workout_log_id} [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	optional.IDField
	optional.DurationField
	optional.IntensityField
	optional.PlaceField
	optional.CreateAtField
	Workout        *Workout         `json:"workout,omitempty"`
	Course         *Course          `json:"course,omitempty"`
	WorkoutSetLogs []*WorkoutSetLog `json:"workout_set_logs,omitempty"`
}

type Course struct {
	courseOptional.IDField
	courseOptional.SaleTypeField
	courseOptional.ScheduleTypeField
	courseOptional.NameField
	courseOptional.CoverField
}

type Workout struct {
	workoutOptional.IDField
	workoutOptional.NameField
	workoutOptional.EquipmentField
	workoutOptional.StartAudioField
	workoutOptional.EndAudioField
	workoutOptional.WorkoutSetCountField
	workoutOptional.CreateAtField
	workoutOptional.UpdateAtField
}

type WorkoutSetLog struct {
	workoutSetLogOptional.IDField
	workoutSetLogOptional.WeightField
	workoutSetLogOptional.RepsField
	workoutSetLogOptional.DistanceField
	workoutSetLogOptional.DurationField
	workoutSetLogOptional.InclineField
	workoutSetLogOptional.CreateAtField
	WorkoutSet *struct {
		workoutSetOptional.IDField
		workoutSetOptional.TypeField
		workoutSetOptional.WeightField
		workoutSetOptional.RepsField
		workoutSetOptional.DistanceField
		workoutSetOptional.DurationField
		workoutSetOptional.InclineField
		Action *struct {
			actionOptional.IDField
			actionOptional.NameField
			actionOptional.SourceField
			actionOptional.TypeField
		} `json:"action,omitempty"`
	} `json:"workout_set,omitempty"`
}
