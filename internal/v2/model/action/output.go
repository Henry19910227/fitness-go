package action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
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
	MaxDistanceRecord *max_distance_record.Output `json:"max_distance_record" gorm:"foreignKey:action_id;references:id"` // 最長距離紀錄
	MaxRepsRecord     *max_reps_record.Output     `json:"max_reps_record" gorm:"foreignKey:action_id;references:id"`     // 最多次數紀錄
	MaxRMRecord       *max_rm_record.Output       `json:"max_rm_record" gorm:"foreignKey:action_id;references:id"`       // 最大RM紀錄
	MaxSpeedRecord    *max_speed_record.Output    `json:"max_speed_record" gorm:"foreignKey:action_id;references:id"`    // 最高速度紀錄
	MaxWeightRecord   *max_weight_record.Output   `json:"max_weight_record" gorm:"foreignKey:action_id;references:id"`   // 最大重量紀錄
	MinDurationRecord *min_duration_record.Output `json:"min_duration_record" gorm:"foreignKey:action_id;references:id"` // 最短時長紀錄
}

func (Output) TableName() string {
	return "actions"
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

// APIGetCMSActionsOutput /cms/actions [GET] 獲取動作列表 API
type APIGetCMSActionsOutput struct {
	base.Output
	Data   APIGetCMSActionsData `json:"data"`
	Paging *paging.Output       `json:"paging,omitempty"`
}
type APIGetCMSActionsData []*struct {
	IDField
	NameField
	SourceField
	TypeField
	CategoryField
	BodyField
	EquipmentField
	IntroField
	CoverField
	VideoField
	StatusField
	CreateAtField
	UpdateAtField
}

// APICreateCMSActionOutput /cms/action [POST] 創建動作 API
type APICreateCMSActionOutput struct {
	base.Output
	Data *APICreateCMSActionData `json:"data,omitempty"`
}
type APICreateCMSActionData struct {
	IDField
	NameField
	TypeField
	CategoryField
	BodyField
	EquipmentField
	IntroField
	CoverField
	VideoField
	StatusField
	CreateAtField
	UpdateAtField
}

// APICreateUserActionOutput /v2/user/action [POST] 新增個人動作 API
type APICreateUserActionOutput struct {
	base.Output
	Data *APICreateUserActionData `json:"data,omitempty"`
}
type APICreateUserActionData struct {
	IDField
}

// APIUpdateUserActionOutput /v2/user/action/{action_id} [PATCH] 修改個人動作 API
type APIUpdateUserActionOutput struct {
	base.Output
}

// APIGetUserActionsOutput /v2/user/actions [GET] 獲取個人動作庫 API
type APIGetUserActionsOutput struct {
	base.Output
	Data   *APIGetUserActionsData `json:"data,omitempty"`
	Paging *paging.Output         `json:"paging,omitempty"`
}
type APIGetUserActionsData []*struct {
	IDField
	NameField
	SourceField
	TypeField
	CategoryField
	BodyField
	EquipmentField
	IntroField
	CoverField
	VideoField
	StatusField
	CreateAtField
	UpdateAtField
}

// APIDeleteUserActionOutput /v2/user/action/{action_id} [DELETE]
type APIDeleteUserActionOutput struct {
	base.Output
}

// APIDeleteUserActionVideoOutput /v2/user/action/{action_id}/video [DELETE]
type APIDeleteUserActionVideoOutput struct {
	base.Output
}

// APIGetTrainerActionsOutput /v2/trainer/actions [GET] 獲取個人動作庫 API
type APIGetTrainerActionsOutput struct {
	base.Output
	Data   *APIGetTrainerActionsData `json:"data,omitempty"`
	Paging *paging.Output            `json:"paging,omitempty"`
}
type APIGetTrainerActionsData []*struct {
	IDField
	NameField
	SourceField
	TypeField
	CategoryField
	BodyField
	EquipmentField
	IntroField
	CoverField
	VideoField
	StatusField
	CreateAtField
	UpdateAtField
}

// APICreateTrainerActionOutput /v2/trainer/action [POST] 新增教練動作 API
type APICreateTrainerActionOutput struct {
	base.Output
	Data *APICreateTrainerActionData `json:"data,omitempty"`
}
type APICreateTrainerActionData struct {
	IDField
}

// APIUpdateTrainerActionOutput /v2/trainer/action/{action_id} [PATCH] 修改教練動作 API
type APIUpdateTrainerActionOutput struct {
	base.Output
}

// APIDeleteTrainerActionOutput /v2/trainer/action/{action_id} [DELETE]
type APIDeleteTrainerActionOutput struct {
	base.Output
}

// APIDeleteTrainerActionVideoOutput /v2/trainer/action/{action_id}/video [DELETE]
type APIDeleteTrainerActionVideoOutput struct {
	base.Output
}