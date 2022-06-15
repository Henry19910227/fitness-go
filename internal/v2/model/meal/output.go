package meal

import "github.com/Henry19910227/fitness-go/internal/v2/model/food"

type Table struct {
	IDField
	DietIDField
	FoodIDField
	TypeField
	AmountField
	CreateAtField
	Food *food.Table `json:"food,omitempty" gorm:"foreignKey:id;references:food_id"` // 食物
}

func (Table) TableName() string {
	return "meals"
}
