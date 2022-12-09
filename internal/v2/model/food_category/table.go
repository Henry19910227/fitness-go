package food_category

import foodCategoryOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food_category/optional"

type Table struct {
	foodCategoryOptional.IDField
	foodCategoryOptional.TagField
	foodCategoryOptional.TitleField
	foodCategoryOptional.IsDeletedField
	foodCategoryOptional.CreateAtField
	foodCategoryOptional.UpdateAtField
}

func (Table) TableName() string {
	return "food_categories"
}
