package food_category

import (
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/resolver/food_category"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := foodCategory.NewResolver(db)
	return New(resolver)
}
