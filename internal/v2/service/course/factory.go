package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/course"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := course.New(db)
	return New(repository)
}
