package banner

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/banner"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := banner.New(db)
	return New(repository)
}
