package body_image

import (
	bodyImageService "github.com/Henry19910227/fitness-go/internal/v2/service/body_image"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	bodyImageSvc := bodyImageService.NewService(db)
	return New(bodyImageSvc)
}
