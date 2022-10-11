package favorite_action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/favorite_action"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := favorite_action.NewResolver(db)
	return New(resolver)
}