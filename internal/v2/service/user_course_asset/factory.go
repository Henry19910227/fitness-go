package user_course_asset

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/user_course_asset"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := user_course_asset.New(db)
	return New(repository)
}
