package subscribe_plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/subscribe_plan"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := subscribe_plan.NewResolver(db)
	return New(resolver)
}
