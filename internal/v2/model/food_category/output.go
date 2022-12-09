package food_category

import (
	foodCategoryOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food_category/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

type Output struct {
	Table
}

func (Output) TableName() string {
	return "food_categories"
}

// APIGetCMSFoodCategoriesOutput v2/cms/food_categories 獲取食物分類
type APIGetCMSFoodCategoriesOutput struct {
	base.Output
	Data APIGetCMSFoodCategoriesData `json:"data"`
}
type APIGetCMSFoodCategoriesData []*struct {
	foodCategoryOptional.IDField
	foodCategoryOptional.TagField
	foodCategoryOptional.TitleField
	foodCategoryOptional.CreateAtField
	foodCategoryOptional.UpdateAtField
}
