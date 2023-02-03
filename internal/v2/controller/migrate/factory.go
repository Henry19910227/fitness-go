package migrate

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/migrate"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := migrate.NewResolver(db)
	return New(resolver)
}
