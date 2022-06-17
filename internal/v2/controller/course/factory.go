package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := course.NewResolver(db)
	return New(resolver)
}
