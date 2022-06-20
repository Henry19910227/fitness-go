package food

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/food"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := food.NewResolver(db)
	return New(resolver)
}
