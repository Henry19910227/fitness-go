package food

import (
	foodService "github.com/Henry19910227/fitness-go/internal/v2/service/food"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	foodSvc := foodService.NewService(db)
	return New(foodSvc)
}
