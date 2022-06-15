package meal

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	DietIDField
	PagingInput
	PreloadInput
	OrderByInput
}

type DeleteInput struct {
	IDField
	DietIDField
}

func (ListInput) TableName() string {
	return "meals"
}
