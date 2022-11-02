package certificate

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/certificate"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := certificate.New(db)
	return New(repository)
}

