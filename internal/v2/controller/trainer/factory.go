package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/trainer"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := trainer.NewResolver(db)
	return New(resolver)
}
