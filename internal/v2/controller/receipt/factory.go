package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/receipt"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := receipt.NewResolver(db)
	return New(resolver)
}
