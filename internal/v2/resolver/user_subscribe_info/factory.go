package user_subscribe_info

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	subscribeInfoService := user_subscribe_info.NewService(db)
	return New(subscribeInfoService)
}
