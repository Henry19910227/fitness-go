package ios_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/ios_version"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := ios_version.New(db)
	return New(repository)
}
