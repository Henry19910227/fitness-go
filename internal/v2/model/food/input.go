package food

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type GenerateInput struct {
	DataAmount     int
	UserID         []*base.GenerateSetting
	FoodCategoryID []*base.GenerateSetting
	Source         []*base.GenerateSetting
	IsDeleted      []*base.GenerateSetting
}

type ListInput struct {
	foodCategory.TagField
	UserIDOptional
	NameOptional
	SourceOptional
	StatusOptional
	IsDeletedOptional
	PagingInput
	PreloadInput
	OrderByInput
}

type APIGetFoodsInput struct {
	foodCategory.TagField
	UserIDField
	NameField
}
