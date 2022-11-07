package favorite_course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/favorite_course"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	favoriteCourseService := favorite_course.NewService(db)
	return New(favoriteCourseService)
}
