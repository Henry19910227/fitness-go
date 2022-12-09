package meal

import (
	foodOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food/optional"
	mealOptional "github.com/Henry19910227/fitness-go/internal/v2/field/meal/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food"
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
)

type Output struct {
	Table
	Food *food.Output `json:"food,omitempty" gorm:"foreignKey:id;references:food_id"` // 食物
}

func (Output) TableName() string {
	return "meals"
}

// APIPutMealsOutput /v2/diet/{diet_id}/meals [PUT] 修改並覆蓋餐食
type APIPutMealsOutput struct {
	base.Output
}

// APIGetMealsOutput /v2/diet/{diet_id}/meals [Get]
type APIGetMealsOutput struct {
	base.Output
	Data APIGetMealsData `json:"data"`
}
type APIGetMealsData []*struct {
	mealOptional.IDField
	mealOptional.TypeField
	mealOptional.AmountField
	mealOptional.CreateAtField
	Food *struct {
		foodOptional.IDField
		foodOptional.SourceField
		foodOptional.NameField
		foodOptional.CalorieField
		foodOptional.AmountDescField
		FoodCategory *struct {
			foodCategory.IDField
			foodCategory.TagField
			foodCategory.TitleField
		} `json:"food_category,omitempty"`
	} `json:"food,omitempty"`
}
