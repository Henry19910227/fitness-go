package food

import (
	foodCalorie "github.com/Henry19910227/fitness-go/internal/pkg/tool/food_calorie"
	foodService "github.com/Henry19910227/fitness-go/internal/v2/service/food"
	foodCategoryService "github.com/Henry19910227/fitness-go/internal/v2/service/food_category"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	foodSvc := foodService.NewService(db)
	categorySvc := foodCategoryService.NewService(db)
	calorieTool := foodCalorie.New()
	return New(foodSvc, categorySvc, calorieTool)
}
