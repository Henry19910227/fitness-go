package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Output struct {
	Table
	Action *action.Output `json:"action,omitempty" gorm:"foreignKey:id;references:action_id"` // 動作
}

func (Output) TableName() string {
	return "workout_sets"
}

// APIGetCMSWorkoutSetsOutput v2/cms/workout/{workout_id}/workout_sets [GET] 獲取訓練組列表 API
type APIGetCMSWorkoutSetsOutput struct {
	base.Output
	Data   APIGetCMSWorkoutSetsData `json:"data"`
	Paging *paging.Output           `json:"paging,omitempty"`
}
type APIGetCMSWorkoutSetsData []*struct {
	IDField
	TypeField
	AutoNextField
	StartAudioField
	ProgressAudioField
	RemarkField
	WeightField
	RepsField
	DistanceField
	DurationField
	InclineField
	CreateAtField
	UpdateAtField
	Action *struct {
		action.IDField
		action.NameField
		action.SourceField
		action.TypeField
		action.CategoryField
		action.BodyField
		action.EquipmentField
		action.IntroField
		action.CoverField
		action.VideoField
		action.CreateAtField
		action.UpdateAtField
	} `json:"action,omitempty"`
}
