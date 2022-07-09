package banner

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/banner"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := banner.NewResolver(db)
	return New(resolver)
}
