package meal

import (
	mealOptional "github.com/Henry19910227/fitness-go/internal/v2/field/meal/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

// GenerateInput Test Input
type GenerateInput struct {
	DataAmount int
	DietID     []*base.GenerateSetting
	FoodID     []*base.GenerateSetting
	Type       []*base.GenerateSetting
}

// UserIDInput Public Input
type UserIDInput struct {
	UserID *int64 `json:"user_id,omitempty"` //用戶id
}

// ListInput Service Input
type ListInput struct {
	mealOptional.DietIDField
	UserIDInput
	PagingInput
	PreloadInput
	OrderByInput
}

// DeleteInput Service Input
type DeleteInput struct {
	mealOptional.IDField
	mealOptional.DietIDField
}

//// APIPutMealsInput /v2/diet/{diet_id}/meals [PUT] 修改並覆蓋餐食 API
//type APIPutMealsInput struct {
//	mealOptional.DietIDField
//	UserIDInput
//	Meals []*APIPutMealsInputItem
//}
//type APIPutMealsInputItem struct {
//	mealOptional.FoodIDField
//	mealOptional.TypeField
//	mealOptional.AmountField
//}

// APIGetMealsInput /v2/meals [GET] 獲取餐食列表 API
type APIGetMealsInput struct {
	UserIDInput
}
