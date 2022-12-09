package food

import (
	foodOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food/optional"
	foodRequired "github.com/Henry19910227/fitness-go/internal/v2/field/food/required"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/where"
)

type PagingInput = struct {
	pagingOptional.PageField
	pagingOptional.SizeField
}
type PreloadInput = preload.Input
type WhereInput = where.Input
type JoinInput = join.Input
type OrderByInput = orderBy.Input
type CustomOrderByInput = orderBy.CustomInput

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
	foodOptional.UserIDField
	foodOptional.NameField
	foodOptional.SourceField
	foodOptional.StatusField
	foodOptional.IsDeletedField
	PagingInput
	JoinInput
	WhereInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
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
