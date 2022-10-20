package plan

import (
	planOptional "github.com/Henry19910227/fitness-go/internal/v2/field/plan/optional"
	workoutOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_plan_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout"
)

type Output struct {
	Table
	Workouts          []*workout.Output           `json:"workouts,omitempty" gorm:"foreignKey:plan_id;references:id"`            // 訓練
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
	planOptional.IDField
	planOptional.NameField
	planOptional.WorkoutCountField
	planOptional.CreateAtField
	planOptional.UpdateAtField
	Workouts []*struct {
		workoutOptional.IDField
		workoutOptional.NameField
		workoutOptional.EquipmentField
		workoutOptional.StartAudioField
		workoutOptional.EndAudioField
		workoutOptional.WorkoutSetCountField
		workoutOptional.CreateAtField
		workoutOptional.UpdateAtField
	} `json:"workouts,omitempty"`
}

// APICreateUserPlanOutput /v2/user/course/{course_id}/plan [POST]
type APICreateUserPlanOutput struct {
	base.Output
	Data *APICreateUserPlanData `json:"data,omitempty"`
}
type APICreateUserPlanData struct {
	planOptional.IDField
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
	planOptional.IDField
	planOptional.NameField
	planOptional.WorkoutCountField
	planOptional.CreateAtField
	planOptional.UpdateAtField
	UserPlanStatistic *struct {
		user_plan_statistic.DurationField
		user_plan_statistic.FinishWorkoutCountField
	} `json:"user_plan_statistic,omitempty"`
}

// APIUpdateUserPlanOutput /v2/user/plan/{plan_id} [PATCH]
type APIUpdateUserPlanOutput struct {
	base.Output
}

// APICreateTrainerPlanOutput /v2/trainer/course/{course_id}/plan [POST]
type APICreateTrainerPlanOutput struct {
	base.Output
	Data *APICreateTrainerPlanData `json:"data,omitempty"`
}
type APICreateTrainerPlanData struct {
	planOptional.IDField
}

// APIGetTrainerPlansOutput /v2/trainer/course/{course_id}/plans [GET]
type APIGetTrainerPlansOutput struct {
	base.Output
	Data *APIGetTrainerPlansData `json:"data,omitempty"`
}
type APIGetTrainerPlansData []*struct {
	planOptional.IDField
	planOptional.NameField
	planOptional.WorkoutCountField
	planOptional.CreateAtField
	planOptional.UpdateAtField
	UserPlanStatistic *struct {
		user_plan_statistic.DurationField
		user_plan_statistic.FinishWorkoutCountField
	} `json:"user_plan_statistic,omitempty"`
}

// APIDeleteTrainerPlanOutput /v2/trainer/plan/{plan_id} [DELETE]
type APIDeleteTrainerPlanOutput struct {
	base.Output
}

// APIGetProductPlansOutput /v2/product/course/{course_id}/plans [GET]
type APIGetProductPlansOutput struct {
	base.Output
	Data *APIGetProductPlansData `json:"data,omitempty"`
}
type APIGetProductPlansData []*struct {
	planOptional.IDField
	planOptional.NameField
	planOptional.WorkoutCountField
	planOptional.CreateAtField
	planOptional.UpdateAtField
}
