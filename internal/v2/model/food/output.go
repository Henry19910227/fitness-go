package food

import (
	foodOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food/optional"
	foodCategoryOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food_category/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Output struct {
	Table
	FoodCategory *foodCategory.Output `json:"food_category,omitempty" gorm:"foreignkey:id;references:food_category_id"` // 餐食
}

func (Output) TableName() string {
	return "foods"
}



// APIGetCMSFoodsOutput /v2/cms/foods [GET] 獲取食物列表 API
type APIGetCMSFoodsOutput struct {
	base.Output
	Data   APIGetCMSFoodsData `json:"data"`
	Paging *paging.Output     `json:"paging,omitempty"`
}
type APIGetCMSFoodsData []*struct {
	foodOptional.IDField
	foodOptional.NameField
	foodOptional.SourceField
	foodOptional.StatusField
	foodOptional.AmountDescField
	foodOptional.CreateAtField
	foodOptional.UpdateAtField
	FoodCategory *struct {
		foodCategoryOptional.IDField
		foodCategoryOptional.TagField
		foodCategoryOptional.TitleField
	} `json:"food_category,omitempty"`
}

// APICreateCMSFoodOutput /v2/cms/food [POST] 創建食物 API
type APICreateCMSFoodOutput struct {
	base.Output
	Data *APICreateCMSFoodData `json:"data,omitempty"`
}
type APICreateCMSFoodData struct {
	foodOptional.IDField
	foodOptional.UserIDField
	foodOptional.SourceField
	foodOptional.NameField
	foodOptional.CalorieField
	foodOptional.AmountDescField
	foodOptional.CreateAtField
	foodOptional.UpdateAtField
	FoodCategory *struct {
		foodCategoryOptional.IDField
		foodCategoryOptional.TagField
		foodCategoryOptional.TitleField
	} `json:"food_category,omitempty"`
}
