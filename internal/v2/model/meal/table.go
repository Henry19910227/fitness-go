package meal

import mealOptional "github.com/Henry19910227/fitness-go/internal/v2/field/meal/optional"

type Table struct {
	mealOptional.IDField
	mealOptional.DietIDField
	mealOptional.FoodIDField
	mealOptional.TypeField
	mealOptional.AmountField
	mealOptional.CreateAtField
}

func (Table) TableName() string {
	return "meals"
}
