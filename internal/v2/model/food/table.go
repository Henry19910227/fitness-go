package food

import foodOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food/optional"

type Table struct {
	foodOptional.IDField
	foodOptional.UserIDField
	foodOptional.FoodCategoryIDField
	foodOptional.SourceField
	foodOptional.NameField
	foodOptional.CalorieField
	foodOptional.AmountDescField
	foodOptional.StatusField
	foodOptional.IsDeletedField
	foodOptional.CreateAtField
	foodOptional.UpdateAtField
}

func (Table) TableName() string {
	return "foods"
}
