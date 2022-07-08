package review_image

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/review_image"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := review_image.NewResolver(db)
	return New(resolver)
}
