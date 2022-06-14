package plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout"
)

type Table struct {
	IDField
	CourseIDField
	NameField
	WorkoutCountField
	CreateAtField
	UpdateAtField
	Workout []*workout.Table `json:"workouts,omitempty" gorm:"foreignKey:plan_id;references:id"` // 訓練
}

func (Table) TableName() string {
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
