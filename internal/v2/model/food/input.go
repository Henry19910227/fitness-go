package food

import (
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	foodCategory.TagField
	UserIDField
	NameField
	IsDeletedField
	PagingInput
	PreloadInput
	OrderByInput
}

type APIGetFoodsInput struct {
	foodCategory.TagField
	UserIDField
	NameField
}
