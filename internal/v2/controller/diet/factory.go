package diet

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/diet"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := diet.NewResolver(db)
	return New(resolver)
}
