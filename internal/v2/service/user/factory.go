package user

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/user"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := user.New(db)
	return New(repository)
}
