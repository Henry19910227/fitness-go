package android_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/android_version"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := android_version.NewResolver(db)
	return New(resolver)
}
