package review_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	reviewImageService "github.com/Henry19910227/fitness-go/internal/v2/service/review_image"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	receiptImageSvc := reviewImageService.NewService(db)
	uploadTool := uploader.NewReviewImageTool()
	return New(receiptImageSvc, uploadTool)
}
