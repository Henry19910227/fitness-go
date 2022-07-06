package user

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/user"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := user.NewResolver(db)
	return New(resolver)
}
