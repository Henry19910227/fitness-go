package workout_log

import (
	actionOptional "github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	workoutOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_log/optional"
	workoutSetOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/optional"
	workoutSetLogOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout_set_log/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log"
)

type Output struct {
	Table
	Workout        *WorkoutOutput            `json:"workout,omitempty" gorm:"foreignKey:workout_id;references:id"`
	WorkoutSetLogs []*workout_set_log.Output `json:"workout_set_logs,omitempty" gorm:"foreignKey:workout_log_id;references:id"`
}

func (Output) TableName() string {
	return "workout_logs"
}

type WorkoutOutput struct {
	WorkoutTable
}

func (WorkoutOutput) TableName() string {
	return "workouts"
}

func (o Output) WorkoutOnSafe() WorkoutOutput {
	if o.Workout != nil {
		return *o.Workout
	}
	return WorkoutOutput{}
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
	workoutSetLogOptional.WorkoutSetIDField
	workoutSetLogOptional.WeightField
	workoutSetLogOptional.DistanceField
	workoutSetLogOptional.InclineField
	workoutSetLogOptional.RepsField
	workoutSetLogOptional.DurationField
	NewRecord  *int `json:"new_record" example:"1"` //是否是新紀錄(0:否/1:是)
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

// APIGetUserWorkoutLogsOutput /v2/user/workout_logs [GET]
type APIGetUserWorkoutLogsOutput struct {
	base.Output
	Data   *APIGetUserWorkoutLogsData `json:"data,omitempty"`
	Paging *paging.Output             `json:"paging,omitempty"`
}
type APIGetUserWorkoutLogsData []*struct {
	optional.IDField
	optional.DurationField
	optional.IntensityField
	optional.PlaceField
	optional.CreateAtField
	Workout *struct {
		workoutOptional.IDField
		workoutOptional.NameField
		workoutOptional.EquipmentField
		workoutOptional.StartAudioField
		workoutOptional.EndAudioField
		workoutOptional.WorkoutSetCountField
		workoutOptional.CreateAtField
		workoutOptional.UpdateAtField
	} `json:"workout,omitempty"`
}
