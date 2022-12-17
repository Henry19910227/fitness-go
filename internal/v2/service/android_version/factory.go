package android_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/android_version"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := android_version.New(db)
	return New(repository)
}
