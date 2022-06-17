package meal

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/meal"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := meal.NewResolver(db)
	return New(resolver)
}
