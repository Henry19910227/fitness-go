package meal

import (
	mealService "github.com/Henry19910227/fitness-go/internal/v2/service/meal"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	foodSvc := mealService.NewService(db)
	return New(foodSvc)
}
