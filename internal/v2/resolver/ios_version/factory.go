package ios_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/ios_version"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	versionService := ios_version.NewService(db)
	return New(versionService)
}
