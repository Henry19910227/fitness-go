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
