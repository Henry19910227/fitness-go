package action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/favorite_action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/max_distance_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/max_reps_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/max_rm_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/max_speed_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/max_weight_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/min_duration_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Output struct {
	Table
	Course            *CourseOutput               `json:"course" gorm:"foreignKey:id;references:course_id"`              // 課表
	MaxDistanceRecord *max_distance_record.Output `json:"max_distance_record" gorm:"foreignKey:action_id;references:id"` // 最長距離紀錄
	MaxRepsRecord     *max_reps_record.Output     `json:"max_reps_record" gorm:"foreignKey:action_id;references:id"`     // 最多次數紀錄
	MaxRMRecord       *max_rm_record.Output       `json:"max_rm_record" gorm:"foreignKey:action_id;references:id"`       // 最大RM紀錄
	MaxSpeedRecord    *max_speed_record.Output    `json:"max_speed_record" gorm:"foreignKey:action_id;references:id"`    // 最高速度紀錄
	MaxWeightRecord   *max_weight_record.Output   `json:"max_weight_record" gorm:"foreignKey:action_id;references:id"`   // 最大重量紀錄
	MinDurationRecord *min_duration_record.Output `json:"min_duration_record" gorm:"foreignKey:action_id;references:id"` // 最短時長紀錄
	FavoriteAction    *favorite_action.Output     `json:"favorite_action" gorm:"foreignKey:action_id;references:id"`     // 動作收藏
}

func (Output) TableName() string {
	return "actions"
}

type CourseOutput struct {
	CourseTable
}

func (CourseOutput) TableName() string {
	return "courses"
}

func (o Output) CourseOnSafe() CourseOutput {
	if o.Course != nil {
		return *o.Course
	}
	return CourseOutput{}
}

func (o Output) MaxDistanceRecordOnSafe() max_distance_record.Output {
	if o.MaxDistanceRecord != nil {
		return *o.MaxDistanceRecord
	}
	return max_distance_record.Output{}
}

func (o Output) MaxRepsRecordOnSafe() max_reps_record.Output {
	if o.MaxRepsRecord != nil {
		return *o.MaxRepsRecord
	}
	return max_reps_record.Output{}
}

func (o Output) MaxRMRecordOnSafe() max_rm_record.Output {
	if o.MaxRMRecord != nil {
		return *o.MaxRMRecord
	}
	return max_rm_record.Output{}
}

func (o Output) MaxSpeedRecordOnSafe() max_speed_record.Output {
	if o.MaxSpeedRecord != nil {
		return *o.MaxSpeedRecord
	}
	return max_speed_record.Output{}
}

func (o Output) MaxWeightRecordOnSafe() max_weight_record.Output {
	if o.MaxWeightRecord != nil {
		return *o.MaxWeightRecord
	}
	return max_weight_record.Output{}
}

func (o Output) MinDurationRecordOnSafe() min_duration_record.Output {
	if o.MinDurationRecord != nil {
		return *o.MinDurationRecord
	}
	return min_duration_record.Output{}
}

func (o Output) FavoriteActionOnSafe() favorite_action.Output {
	if o.FavoriteAction != nil {
		return *o.FavoriteAction
	}
	return favorite_action.Output{}
}

// APIGetCMSActionsOutput /cms/actions [GET] 獲取動作列表 API
type APIGetCMSActionsOutput struct {
	base.Output
	Data   APIGetCMSActionsData `json:"data"`
	Paging *paging.Output       `json:"paging,omitempty"`
}
type APIGetCMSActionsData []*struct {
	optional.IDField
	optional.NameField
	optional.SourceField
	optional.TypeField
	optional.CategoryField
	optional.BodyField
	optional.EquipmentField
	optional.IntroField
	optional.CoverField
	optional.VideoField
	optional.StatusField
	optional.CreateAtField
	optional.UpdateAtField
}

// APICreateCMSActionOutput /cms/action [POST] 創建動作 API
type APICreateCMSActionOutput struct {
	base.Output
	Data *APICreateCMSActionData `json:"data,omitempty"`
}
type APICreateCMSActionData struct {
	optional.IDField
	optional.NameField
	optional.TypeField
	optional.CategoryField
	optional.BodyField
	optional.EquipmentField
	optional.IntroField
	optional.CoverField
	optional.VideoField
	optional.StatusField
	optional.CreateAtField
	optional.UpdateAtField
}

// APICreateUserActionOutput /v2/user/action [POST] 新增個人動作 API
type APICreateUserActionOutput struct {
	base.Output
	Data *APICreateUserActionData `json:"data,omitempty"`
}
type APICreateUserActionData struct {
	optional.IDField
	optional.NameField
	optional.TypeField
	optional.CategoryField
	optional.BodyField
	optional.EquipmentField
	optional.IntroField
	optional.CoverField
	optional.VideoField
	optional.StatusField
	optional.CreateAtField
	optional.UpdateAtField
}

// APIUpdateUserActionOutput /v2/user/action/{action_id} [PATCH] 修改個人動作 API
type APIUpdateUserActionOutput struct {
	base.Output
	Data *APIUpdateUserActionData `json:"data,omitempty"`
}
type APIUpdateUserActionData struct {
	optional.IDField
	optional.NameField
	optional.TypeField
	optional.CategoryField
	optional.BodyField
	optional.EquipmentField
	optional.IntroField
	optional.CoverField
	optional.VideoField
	optional.StatusField
	optional.CreateAtField
	optional.UpdateAtField
}

// APIGetUserActionsOutput /v2/user/actions [GET] 獲取個人動作庫 API
type APIGetUserActionsOutput struct {
	base.Output
	Data   *APIGetUserActionsData `json:"data,omitempty"`
	Paging *paging.Output         `json:"paging,omitempty"`
}
type APIGetUserActionsData []*struct {
	APIGetUserActionsItem
}
type APIGetUserActionsItem struct {
	optional.IDField
	optional.NameField
	optional.SourceField
	optional.TypeField
	optional.CategoryField
	optional.BodyField
	optional.EquipmentField
	optional.IntroField
	optional.CoverField
	optional.VideoField
	optional.StatusField
	Favorite *int `json:"favorite,omitempty" example:"1"` //是否收藏(0:否/1:是)
	optional.CreateAtField
	optional.UpdateAtField
}

// APIDeleteUserActionOutput /v2/user/action/{action_id} [DELETE]
type APIDeleteUserActionOutput struct {
	base.Output
}

// APIDeleteUserActionVideoOutput /v2/user/action/{action_id}/video [DELETE]
type APIDeleteUserActionVideoOutput struct {
	base.Output
}

// APIGetUserActionSystemImagesOutput /v2/user/action/system_images [GET]
type APIGetUserActionSystemImagesOutput struct {
	base.Output
	Data *[]string `json:"data,omitempty"`
}

// APIGetTrainerActionsOutput /v2/trainer/actions [GET] 獲取個人動作庫 API
type APIGetTrainerActionsOutput struct {
	base.Output
	Data   *APIGetTrainerActionsData `json:"data,omitempty"`
	Paging *paging.Output            `json:"paging,omitempty"`
}
type APIGetTrainerActionsData []*struct {
	optional.IDField
	optional.NameField
	optional.SourceField
	optional.TypeField
	optional.CategoryField
	optional.BodyField
	optional.EquipmentField
	optional.IntroField
	optional.CoverField
	optional.VideoField
	optional.StatusField
	optional.CreateAtField
	optional.UpdateAtField
}

// APIDeleteTrainerActionVideoOutput /v2/trainer/action/{action_id}/video [DELETE]
type APIDeleteTrainerActionVideoOutput struct {
	base.Output
}
