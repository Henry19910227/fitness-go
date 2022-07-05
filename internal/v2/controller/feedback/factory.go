package feedback

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/feedback"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := feedback.NewResolver(db)
	return New(resolver)
}