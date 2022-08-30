package workout

import (
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
	IDField
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
	IDField
	NameField
	EquipmentField
	StartAudioField
	EndAudioField
	WorkoutSetCountField
	Finish *int `json:"finish" example:"1"` // 是否完成(0:未完成/2:已完成)
	CreateAtField
	UpdateAtField
}
