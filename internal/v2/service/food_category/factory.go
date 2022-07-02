package food_category

import (
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/repository/food_category"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := foodCategory.New(db)
	return New(repository)
}
