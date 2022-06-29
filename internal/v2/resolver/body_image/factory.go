package body_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	bodyImageService "github.com/Henry19910227/fitness-go/internal/v2/service/body_image"
	bodyService "github.com/Henry19910227/fitness-go/internal/v2/service/body_record"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	bodyImageSvc := bodyImageService.NewService(db)
	BodySvc := bodyService.NewService(db)
	bodyUploadTool := uploader.NewBodyImageTool()
	return New(bodyImageSvc, BodySvc, bodyUploadTool)
}
