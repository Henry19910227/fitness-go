package food

import (
	foodOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food/optional"
	foodRequired "github.com/Henry19910227/fitness-go/internal/v2/field/food/required"
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
	foodOptional.IDField
}

type ListInput struct {
	foodCategory.TagField
	foodOptional.UserIDField
	foodOptional.NameField
	foodOptional.SourceField
	foodOptional.StatusField
	foodOptional.IsDeletedField
	PagingInput
	PreloadInput
	OrderByInput
}

type APIGetFoodsInput struct {
	foodCategory.TagField
	foodOptional.UserIDField
	foodOptional.NameField
}

// APICreateCMSFoodInput /v2/cms/food [POST]
type APICreateCMSFoodInput struct {
	Body APICreateCMSFoodBody
}
type APICreateCMSFoodBody struct {
	foodRequired.FoodCategoryIDField
	foodRequired.NameField
	foodRequired.AmountDescField
}

// APIUpdateCMSFoodInput /v2/cms/food/{food_id} [PATCH]
type APIUpdateCMSFoodInput struct {
	Uri  APIUpdateCMSFoodURI
	Body APIUpdateCMSFoodBody
}
type APIUpdateCMSFoodURI struct {
	foodRequired.IDField
}
type APIUpdateCMSFoodBody struct {
	foodOptional.NameField
	foodOptional.AmountDescField
	foodOptional.StatusField
}

// APIGetCMSFoodsInput /v2/cms/foods [GET]
type APIGetCMSFoodsInput struct {
	Form APIGetCMSFoodsForm
}
type APIGetCMSFoodsForm struct {
	PagingInput
}
