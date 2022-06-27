package body_record

import (
	body "github.com/Henry19910227/fitness-go/internal/v2/repository/body_record"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := body.New(db)
	return New(repository)
}
