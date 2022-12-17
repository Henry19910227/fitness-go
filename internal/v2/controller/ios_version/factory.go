package ios_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/ios_version"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := ios_version.NewResolver(db)
	return New(resolver)
}
