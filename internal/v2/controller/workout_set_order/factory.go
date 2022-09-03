package workout_set_order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/workout_set_order"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := workout_set_order.NewResolver(db)
	return New(resolver)
}
