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
	CategoryField
	BodyField
	StatusField
	CreateAtField
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
