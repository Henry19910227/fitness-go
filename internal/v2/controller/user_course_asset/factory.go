package user_course_asset

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/user_course_asset"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := user_course_asset.NewResolver(db)
	return New(resolver)
}
