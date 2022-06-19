package meal

import (
	dietService "github.com/Henry19910227/fitness-go/internal/v2/service/diet"
	mealService "github.com/Henry19910227/fitness-go/internal/v2/service/meal"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	foodSvc := mealService.NewService(db)
	dietSvc := dietService.NewService(db)
	return New(foodSvc, dietSvc)
}
