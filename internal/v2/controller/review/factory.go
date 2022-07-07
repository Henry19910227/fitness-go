package review

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/review"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := review.NewResolver(db)
	return New(resolver)
}
