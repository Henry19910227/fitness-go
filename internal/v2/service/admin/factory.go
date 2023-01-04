package admin

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/admin"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := admin.New(db)
	return New(repository)
}
