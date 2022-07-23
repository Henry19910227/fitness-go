package user_subscribe_info

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/user_subscribe_info"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := user_subscribe_info.NewResolver(db)
	return New(resolver)
}
