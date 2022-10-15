package workout_set

import (
	actionOptional "github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/optional"
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

func (o Output) ActionOnSafe() action.Output {
	if o.Action != nil {
		return *o.Action
	}
	return action.Output{}
}

// APIGetCMSWorkoutSetsOutput v2/cms/workout/{workout_id}/workout_sets [GET] 獲取訓練組列表 API
type APIGetCMSWorkoutSetsOutput struct {
	base.Output
	Data   APIGetCMSWorkoutSetsData `json:"data"`
	Paging *paging.Output           `json:"paging,omitempty"`
}
type APIGetCMSWorkoutSetsData []*struct {
	optional.IDField
	optional.TypeField
	optional.AutoNextField
	optional.StartAudioField
	optional.ProgressAudioField
	optional.RemarkField
	optional.WeightField
	optional.RepsField
	optional.DistanceField
	optional.DurationField
	optional.InclineField
	optional.CreateAtField
	optional.UpdateAtField
	Action *struct {
		actionOptional.IDField
		actionOptional.NameField
		actionOptional.SourceField
		actionOptional.TypeField
		actionOptional.CategoryField
		actionOptional.BodyField
		actionOptional.EquipmentField
		actionOptional.IntroField
		actionOptional.CoverField
		actionOptional.VideoField
		actionOptional.CreateAtField
		actionOptional.UpdateAtField
	} `json:"action,omitempty"`
}

// APICreateUserWorkoutSetsOutput /v2/user/workout/{workout_id}/workout_sets [POST]
type APICreateUserWorkoutSetsOutput struct {
	base.Output
	Data *APICreateUserWorkoutSetsData `json:"data,omitempty"`
}
type APICreateUserWorkoutSetsData []int64

// APICreateUserWorkoutSetByDuplicateOutput /v2/user/workout_set/{workout_set_id}/duplicate [POST]
type APICreateUserWorkoutSetByDuplicateOutput struct {
	base.Output
}

// APICreateUserRestSetOutput /v2/user/workout/{workout_id}/rest_set [POST]
type APICreateUserRestSetOutput struct {
	base.Output
}

// APIDeleteUserWorkoutSetOutput /v2/user/workout_set_is/{workout_set_id} [Delete]
type APIDeleteUserWorkoutSetOutput struct {
	base.Output
}

// APIGetUserWorkoutSetsOutput /v2/user/workout/{workout_id}/workout_sets [GET]
type APIGetUserWorkoutSetsOutput struct {
	base.Output
	Data APIGetUserWorkoutSetsData `json:"data"`
}
type APIGetUserWorkoutSetsData []*struct {
	optional.IDField
	optional.TypeField
	optional.AutoNextField
	optional.StartAudioField
	optional.ProgressAudioField
	optional.RemarkField
	optional.WeightField
	optional.RepsField
	optional.DistanceField
	optional.DurationField
	optional.InclineField
	optional.CreateAtField
	optional.UpdateAtField
	Action *struct {
		actionOptional.IDField
		actionOptional.NameField
		actionOptional.SourceField
		actionOptional.TypeField
		actionOptional.CategoryField
		actionOptional.BodyField
		actionOptional.EquipmentField
		actionOptional.IntroField
		actionOptional.CoverField
		actionOptional.VideoField
		actionOptional.CreateAtField
		actionOptional.UpdateAtField
	} `json:"action,omitempty"`
}

// APIUpdateUserWorkoutSetOutput /v2/user/workout_set/{workout_set_id} [PATCH]
type APIUpdateUserWorkoutSetOutput struct {
	base.Output
	Data *APIUpdateUserWorkoutSetData `json:"data,omitempty"`
}
type APIUpdateUserWorkoutSetData struct {
	optional.IDField
	optional.TypeField
	optional.AutoNextField
	optional.StartAudioField
	optional.ProgressAudioField
	optional.RemarkField
	optional.WeightField
	optional.RepsField
	optional.DistanceField
	optional.DurationField
	optional.InclineField
	optional.CreateAtField
	optional.UpdateAtField
	Action *struct {
		actionOptional.IDField
		actionOptional.NameField
		actionOptional.SourceField
		actionOptional.TypeField
		actionOptional.CategoryField
		actionOptional.BodyField
		actionOptional.EquipmentField
		actionOptional.IntroField
		actionOptional.CoverField
		actionOptional.VideoField
		actionOptional.CreateAtField
		actionOptional.UpdateAtField
	} `json:"action,omitempty"`
}

// APIDeleteUserWorkoutSetStartAudioOutput /v2/user/workout_set/{workout_set_id}/start_audio [DELETE]
type APIDeleteUserWorkoutSetStartAudioOutput struct {
	base.Output
}

// APIDeleteUserWorkoutSetProgressAudioOutput /v2/user/workout_set/{workout_set_id}/progress_audio [DELETE]
type APIDeleteUserWorkoutSetProgressAudioOutput struct {
	base.Output
}

// APIGetTrainerWorkoutSetsOutput /v2/trainer/workout/{workout_id}/workout_sets [GET]
type APIGetTrainerWorkoutSetsOutput struct {
	base.Output
	Data *APIGetTrainerWorkoutSetsData `json:"data,omitempty"`
}
type APIGetTrainerWorkoutSetsData []*struct {
	optional.IDField
	optional.TypeField
	optional.AutoNextField
	optional.StartAudioField
	optional.ProgressAudioField
	optional.RemarkField
	optional.WeightField
	optional.RepsField
	optional.DistanceField
	optional.DurationField
	optional.InclineField
	optional.CreateAtField
	optional.UpdateAtField
	Action *struct {
		actionOptional.IDField
		actionOptional.NameField
		actionOptional.SourceField
		actionOptional.TypeField
		actionOptional.CategoryField
		actionOptional.BodyField
		actionOptional.EquipmentField
		actionOptional.IntroField
		actionOptional.CoverField
		actionOptional.VideoField
		actionOptional.CreateAtField
		actionOptional.UpdateAtField
	} `json:"action,omitempty"`
}

// APICreateTrainerWorkoutSetsOutput /v2/trainer/workout/{workout_id}/workout_sets [POST]
type APICreateTrainerWorkoutSetsOutput struct {
	base.Output
	Data *APICreateTrainerWorkoutSetsData `json:"data,omitempty"`
}
type APICreateTrainerWorkoutSetsData []int64

// APICreateTrainerRestSetOutput /v2/user/workout/{workout_id}/rest_set [POST]
type APICreateTrainerRestSetOutput struct {
	base.Output
}

// APIDeleteTrainerWorkoutSetOutput /v2/trainer/workout_set_is/{workout_set_id} [DELETE]
type APIDeleteTrainerWorkoutSetOutput struct {
	base.Output
}

// APIUpdateTrainerWorkoutSetOutput /v2/trainer/workout_set/{workout_set_id} [PATCH]
type APIUpdateTrainerWorkoutSetOutput struct {
	base.Output
	Data *APIUpdateTrainerWorkoutSetData `json:"data,omitempty"`
}
type APIUpdateTrainerWorkoutSetData struct {
	optional.IDField
	optional.TypeField
	optional.AutoNextField
	optional.StartAudioField
	optional.ProgressAudioField
	optional.RemarkField
	optional.WeightField
	optional.RepsField
	optional.DistanceField
	optional.DurationField
	optional.InclineField
	optional.CreateAtField
	optional.UpdateAtField
	Action *struct {
		actionOptional.IDField
		actionOptional.NameField
		actionOptional.SourceField
		actionOptional.TypeField
		actionOptional.CategoryField
		actionOptional.BodyField
		actionOptional.EquipmentField
		actionOptional.IntroField
		actionOptional.CoverField
		actionOptional.VideoField
		actionOptional.CreateAtField
		actionOptional.UpdateAtField
	} `json:"action,omitempty"`
}

// APIDeleteTrainerWorkoutSetStartAudioOutput /v2/trainer/workout_set/{workout_set_id}/start_audio [DELETE]
type APIDeleteTrainerWorkoutSetStartAudioOutput struct {
	base.Output
}

// APIDeleteTrainerWorkoutSetProgressAudioOutput /v2/trainer/workout_set/{workout_set_id}/progress_audio [DELETE]
type APIDeleteTrainerWorkoutSetProgressAudioOutput struct {
	base.Output
}

// APICreateTrainerWorkoutSetByDuplicateOutput /v2/trainer/workout_set/{workout_set_id}/duplicate [POST]
type APICreateTrainerWorkoutSetByDuplicateOutput struct {
	base.Output
}

// APIGetProductWorkoutSetsOutput /v2/product/workout/{workout_id}/workout_sets [GET]
type APIGetProductWorkoutSetsOutput struct {
	base.Output
	Data *APIGetProductWorkoutSetsData `json:"data,omitempty"`
}
type APIGetProductWorkoutSetsData []*struct {
	optional.IDField
	optional.TypeField
	optional.AutoNextField
	optional.StartAudioField
	optional.ProgressAudioField
	optional.RemarkField
	optional.WeightField
	optional.RepsField
	optional.DistanceField
	optional.DurationField
	optional.InclineField
	optional.CreateAtField
	optional.UpdateAtField
	Action *struct {
		actionOptional.IDField
		actionOptional.NameField
		actionOptional.SourceField
		actionOptional.TypeField
		actionOptional.CategoryField
		actionOptional.BodyField
		actionOptional.EquipmentField
		actionOptional.IntroField
		actionOptional.CoverField
		actionOptional.VideoField
		actionOptional.CreateAtField
		actionOptional.UpdateAtField
	} `json:"action,omitempty"`
}