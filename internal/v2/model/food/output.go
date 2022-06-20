package food

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
)

type Output struct {
	Table
	FoodCategory *foodCategory.Output `json:"food_category,omitempty" gorm:"foreignkey:id;references:food_category_id"` // 餐食
}

func (Output) TableName() string {
	return "foods"
}

type APIGetFoodsOutput struct {
	base.Output
	Data APIGetFoodsData `json:"data"`
}
type APIGetFoodsData []*struct {
	IDField
	UserIDField
	SourceField
	NameField
	CalorieField
	AmountDescField
	CreateAtField
	UpdateAtField
	FoodCategory *struct {
		foodCategory.IDField
		foodCategory.TagField
		foodCategory.TitleField
	} `json:"food_category,omitempty"`
}
