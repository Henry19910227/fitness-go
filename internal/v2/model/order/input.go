package order

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	IDOptional
	UserIDOptional
	OrderTypeOptional
	OrderStatusOptional
	OrderByInput
	PagingInput
	PreloadInput
}

// APIGetCMSOrdersInput /v2/cms/orders [GET]
type APIGetCMSOrdersInput struct {
	Form APIGetCMSOrdersForm
}
type APIGetCMSOrdersForm struct {
	IDOptional
	UserIDOptional
	OrderTypeOptional
	OrderStatusOptional
	PagingInput
	OrderByInput
}
