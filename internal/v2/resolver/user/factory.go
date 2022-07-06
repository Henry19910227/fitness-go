package user

import (
	userService "github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	userSvc := userService.NewService(db)
	return New(userSvc)
}
