package order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/order"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := order.NewResolver(db)
	return New(resolver)
}
