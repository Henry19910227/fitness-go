package favorite_course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/favorite_course"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := favorite_course.New(db)
	return New(repository)
}
