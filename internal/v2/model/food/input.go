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

type FindInput struct {
	IDOptional
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

// APICreateCMSFoodInput /v2/cms/food [POST]
type APICreateCMSFoodInput struct {
	Body APICreateCMSFoodBody
}
type APICreateCMSFoodBody struct {
	FoodCategoryIDRequired
	NameRequired
	AmountDescRequired
}

// APIUpdateCMSFoodInput /v2/cms/food/{food_id} [PATCH]
type APIUpdateCMSFoodInput struct {
	Uri  APIUpdateCMSFoodURI
	Body APIUpdateCMSFoodBody
}
type APIUpdateCMSFoodURI struct {
	IDRequired
}
type APIUpdateCMSFoodBody struct {
	NameOptional
	AmountDescOptional
	StatusOptional
}

// APIGetCMSFoodsInput /v2/cms/foods [GET]
type APIGetCMSFoodsInput struct {
	Form APIGetCMSFoodsForm
}
type APIGetCMSFoodsForm struct {
	PagingInput
}
