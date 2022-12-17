package android_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/android_version"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	versionService := android_version.NewService(db)
	return New(versionService)
}
