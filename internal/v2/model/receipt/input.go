package receipt

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	PreloadInput
	OrderByInput
	PagingInput
}

// APIGetCMSReceiptsInput /v2/cms/receipts [GET]
type APIGetCMSReceiptsInput struct {
	Form APIGetCMSReceiptsForm
}
type APIGetCMSReceiptsForm struct {
	OrderByInput
	PagingInput
}
