package action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	IDOptional
}

type ListInput struct {
	IDs []int64 `json:"ids"` //動作id
	TypeOptional
	SourceOptional
	PagingInput
	OrderByInput
}

type APIGetCMSActionsInput struct {
	PagingInput
}

// APICreateCMSActionInput /cms/action [POST] 創建動作 API
type APICreateCMSActionInput struct {
	Form      APICreateCMSActionForm
	CoverFile *file.Input
	VideoFile *file.Input
}
type APICreateCMSActionForm struct {
	NameRequired
	TypeRequired
	CategoryRequired
	BodyRequired
	EquipmentRequired
	IntroRequired
}

// APIUpdateCMSActionInput /v2/cms/action/{action_id} [PATCH] 更新動作 API
type APIUpdateCMSActionInput struct {
	Uri       APIUpdateCMSActionUri
	Form      APIUpdateCMSActionForm
	CoverFile *file.Input
	VideoFile *file.Input
}
type APIUpdateCMSActionForm struct {
	NameOptional
	IntroOptional
	StatusOptional
}
type APIUpdateCMSActionUri struct {
	IDRequired
}

// APICreateUserActionInput /v2/user/action [POST] 新增個人動作 API
type APICreateUserActionInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Cover *file.Input
	Video *file.Input
	Form APICreateUserActionForm
}
type APICreateUserActionForm struct {
	NameRequired
	TypeRequired
	CategoryRequired
	BodyRequired
	EquipmentRequired
	IntroRequired
}