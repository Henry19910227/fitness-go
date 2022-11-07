package favorite_course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/favorite_course"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := favorite_course.NewResolver(db)
	return New(resolver)
}
