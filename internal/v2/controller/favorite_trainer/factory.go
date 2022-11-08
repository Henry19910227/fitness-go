package favorite_trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/favorite_trainer"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := favorite_trainer.NewResolver(db)
	return New(resolver)
}
