package workout

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_log"
)

type Output struct {
	Table
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
	Finish *int `json:"finish" example:"1"` // 是否完成(0:未完成/2:已完成)
	optional.CreateAtField
	optional.UpdateAtField
}

// APIUpdateUserWorkoutOutput /v2/user/workout/{workout_id} [PATCH]
type APIUpdateUserWorkoutOutput struct {
	base.Output
}

// APIDeleteUserWorkoutStartAudioOutput /v2/user/workout/{workout_id}/start_audio [DELETE]
type APIDeleteUserWorkoutStartAudioOutput struct {
	base.Output
}

// APIDeleteUserWorkoutEndAudioOutput /v2/user/workout/{workout_id}/end_audio [DELETE]
type APIDeleteUserWorkoutEndAudioOutput struct {
	base.Output
}
