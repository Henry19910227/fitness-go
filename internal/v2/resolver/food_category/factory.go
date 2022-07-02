package food_category

import (
	foodCategoryService "github.com/Henry19910227/fitness-go/internal/v2/service/food_category"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	foodCategorySvc := foodCategoryService.NewService(db)
	return New(foodCategorySvc)
}
