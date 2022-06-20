package meal

import (
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
	IDField
	TypeField
	AmountField
	CreateAtField
	Food *struct {
		food.IDField
		food.SourceField
		food.NameField
		food.CalorieField
		food.AmountDescField
		FoodCategory *struct {
			foodCategory.IDField
			foodCategory.TagField
			foodCategory.TitleField
		} `json:"food_category,omitempty"`
	} `json:"food,omitempty"`
}
