package admin

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/admin"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := admin.NewResolver(db)
	return New(resolver)
}
