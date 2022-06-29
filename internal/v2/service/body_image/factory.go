package body_image

import (
	bodyImage "github.com/Henry19910227/fitness-go/internal/v2/repository/body_image"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := bodyImage.New(db)
	return New(repository)
}
