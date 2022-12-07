package rda

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/rda"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := rda.New(db)
	return New(repository)
}
