package workout

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_log"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
)

type Output struct {
	Table
	WorkoutSets *[]workout_set.Output `json:"workout_sets,omitempty" gorm:"foreignKey:workout_id;references:id"` // 訓練組列表
	WorkoutLogs *[]workout_log.Output `json:"workout_logs,omitempty" gorm:"foreignKey:workout_id;references:id"` // 訓練log列表
}

func (Output) TableName() string {
	return "workouts"
}

// APICreateUserWorkoutOutput /v2/user/plan/{plan_id}/workout [POST]
type APICreateUserWorkoutOutput struct {
	base.Output
	Data *APICreateUserWorkoutData `json:"data,omitempty"`
}
type APICreateUserWorkoutData struct {
	optional.IDField
}

// APIDeleteUserWorkoutOutput /v2/user/workout/{workout_id} [DELETE]
type APIDeleteUserWorkoutOutput struct {
	base.Output
}

// APIGetUserWorkoutsOutput /v2/user/plan/{plan_is}/workouts [GET]
type APIGetUserWorkoutsOutput struct {
	base.Output
	Data APIGetUserWorkoutsData `json:"data"`
}
type APIGetUserWorkoutsData []*struct {
	APIGetUserWorkoutsDataItem
}
type APIGetUserWorkoutsDataItem struct {
	optional.IDField
	optional.NameField
	optional.EquipmentField
	optional.StartAudioField
	optional.EndAudioField
	optional.WorkoutSetCountField
	Finish *int `json:"finish" example:"1"` // 是否完成(0:未完成/1:已完成)
	optional.CreateAtField
	optional.UpdateAtField
}

// APIUpdateUserWorkoutOutput /v2/user/workout/{workout_id} [PATCH]
type APIUpdateUserWorkoutOutput struct {
	base.Output
	Data *APIUpdateUserWorkoutData `json:"data,omitempty"`
}
type APIUpdateUserWorkoutData struct {
	optional.IDField
	optional.NameField
	optional.EquipmentField
	optional.StartAudioField
	optional.EndAudioField
	optional.WorkoutSetCountField
	optional.CreateAtField
	optional.UpdateAtField
}

// APIDeleteUserWorkoutStartAudioOutput /v2/user/workout/{workout_id}/start_audio [DELETE]
type APIDeleteUserWorkoutStartAudioOutput struct {
	base.Output
}

// APIDeleteUserWorkoutEndAudioOutput /v2/user/workout/{workout_id}/end_audio [DELETE]
type APIDeleteUserWorkoutEndAudioOutput struct {
	base.Output
}

// APIGetTrainerWorkoutsOutput /v2/trainer/plan/{plan_is}/workouts [GET]
type APIGetTrainerWorkoutsOutput struct {
	base.Output
	Data *APIGetTrainerWorkoutsData `json:"data,omitempty"`
}
type APIGetTrainerWorkoutsData []*struct {
	optional.IDField
	optional.NameField
	optional.EquipmentField
	optional.StartAudioField
	optional.EndAudioField
	optional.WorkoutSetCountField
	optional.CreateAtField
	optional.UpdateAtField
}

// APICreateTrainerWorkoutOutput /v2/trainer/plan/{plan_id}/workout [POST]
type APICreateTrainerWorkoutOutput struct {
	base.Output
	Data *APICreateTrainerWorkoutData `json:"data,omitempty"`
}
type APICreateTrainerWorkoutData struct {
	optional.IDField
}

// APIUpdateTrainerWorkoutOutput /v2/trainer/workout/{workout_id} [PATCH]
type APIUpdateTrainerWorkoutOutput struct {
	base.Output
	Data *APIUpdateTrainerWorkoutData `json:"data,omitempty"`
}
type APIUpdateTrainerWorkoutData struct {
	optional.IDField
	optional.NameField
	optional.EquipmentField
	optional.StartAudioField
	optional.EndAudioField
	optional.WorkoutSetCountField
	optional.CreateAtField
	optional.UpdateAtField
}

// APIDeleteTrainerWorkoutOutput /v2/trainer/workout/{workout_id} [DELETE]
type APIDeleteTrainerWorkoutOutput struct {
	base.Output
}

// APIDeleteTrainerWorkoutStartAudioOutput /v2/trainer/workout/{workout_id}/start_audio [DELETE]
type APIDeleteTrainerWorkoutStartAudioOutput struct {
	base.Output
}

// APIDeleteTrainerWorkoutEndAudioOutput /v2/trainer/workout/{workout_id}/end_audio [DELETE]
type APIDeleteTrainerWorkoutEndAudioOutput struct {
	base.Output
}

// APIGetProductWorkoutsOutput /v2/product/plan/{plan_is}/workouts [GET]
type APIGetProductWorkoutsOutput struct {
	base.Output
	Data *APIGetProductWorkoutsData `json:"data,omitempty"`
}
type APIGetProductWorkoutsData []*struct {
	optional.IDField
	optional.NameField
	optional.EquipmentField
	optional.StartAudioField
	optional.EndAudioField
	optional.WorkoutSetCountField
	optional.CreateAtField
	optional.UpdateAtField
}
