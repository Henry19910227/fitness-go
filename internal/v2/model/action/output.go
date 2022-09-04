package action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Output struct {
	Table
}

func (Output) TableName() string {
	return "actions"
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
