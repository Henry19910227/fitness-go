package plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_plan_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout"
)

type Output struct {
	IDField
	CourseIDField
	NameField
	WorkoutCountField
	CreateAtField
	UpdateAtField
	Workout           []*workout.Output           `json:"workouts,omitempty" gorm:"foreignKey:plan_id;references:id"`            // 訓練
	UserPlanStatistic *user_plan_statistic.Output `json:"user_plan_statistic,omitempty" gorm:"foreignKey:plan_id;references:id"` // 計畫統計
}

func (Output) TableName() string {
	return "plans"
}

// APIGetCMSPlansOutput v2/cms/course/{course_id}/plans [GET] 獲取計畫列表 API
type APIGetCMSPlansOutput struct {
	base.Output
	Data   APIGetCMSPlansData `json:"data"`
	Paging *paging.Output     `json:"paging,omitempty"`
}
type APIGetCMSPlansData []*struct {
	IDField
	NameField
	WorkoutCountField
	CreateAtField
	UpdateAtField
	Workouts []*struct {
		workout.IDField
		workout.NameField
		workout.EquipmentField
		workout.StartAudioField
		workout.EndAudioField
		workout.WorkoutSetCountField
		workout.CreateAtField
		workout.UpdateAtField
	} `json:"workouts,omitempty"`
}

// APICreateUserPlanOutput /v2/user/course/{course_id}/plan [POST]
type APICreateUserPlanOutput struct {
	base.Output
	Data *APICreateUserPlanData `json:"data,omitempty"`
}
type APICreateUserPlanData struct {
	IDField
}

// APIDeleteUserPlanOutput /v2/user/plan/{plan_id} [DELETE]
type APIDeleteUserPlanOutput struct {
	base.Output
}

// APIGetUserPlansOutput /v2/user/course/{course_id}/plans [GET]
type APIGetUserPlansOutput struct {
	base.Output
	Data APIGetUserPlansData `json:"data"`
}
type APIGetUserPlansData []*struct {
	IDField
	NameField
	WorkoutCountField
	CreateAtField
	UpdateAtField
	UserPlanStatistic *struct {
		user_plan_statistic.DurationField
		user_plan_statistic.FinishWorkoutCountField
	} `json:"user_plan_statistic,omitempty"`
}
