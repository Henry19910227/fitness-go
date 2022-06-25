package action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/action"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := action.NewResolver(db)
	return New(resolver)
}
