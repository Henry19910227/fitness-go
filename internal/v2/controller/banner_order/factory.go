package banner_order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/banner_order"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := banner_order.NewResolver(db)
	return New(resolver)
}
