package receipt

import (
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
	OrderIDOptional
	TransactionIDOptional
	OriginalTransactionIDOptional
	PreloadInput
	OrderByInput
	PagingInput
}

// APIGetCMSOrderReceiptsInput /v2/cms/order/{order_id}/receipts [GET]
type APIGetCMSOrderReceiptsInput struct {
	Uri  APIGetCMSOrderReceiptsUri
	Form APIGetCMSOrderReceiptsForm
}
type APIGetCMSOrderReceiptsUri struct {
	OrderIDRequired
}
type APIGetCMSOrderReceiptsForm struct {
	OrderByInput
	PagingInput
}
