package plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/plan"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := plan.NewResolver(db)
	return New(resolver)
}
