package food_category

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

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
	IDField
	TagField
	TitleField
	CreateAtField
	UpdateAtField
}
