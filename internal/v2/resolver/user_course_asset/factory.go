package user_course_asset

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_course_asset"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	userService := user.NewService(db)
	userCourseAssetService := user_course_asset.NewService(db)
	return New(userService, userCourseAssetService)
}
