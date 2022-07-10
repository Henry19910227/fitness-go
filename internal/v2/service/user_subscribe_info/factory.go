package user_subscribe_info

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/user_subscribe_info"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := user_subscribe_info.New(db)
	return New(repository)
}