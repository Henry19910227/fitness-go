package food

import "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"

type Table struct {
	IDField
	UserIDField
	FoodCategoryIDField
	SourceField
	NameField
	CalorieField
	AmountDescField
	IsDeletedField
	CreateAtField
	UpdateAtField
	FoodCategory *food_category.Table `json:"food_category,omitempty" gorm:"foreignKey:id;references:food_category_id"` // 餐食
}

func (Table) TableName() string {
	return "foods"
}
