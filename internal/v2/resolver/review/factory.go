package review

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	reviewService "github.com/Henry19910227/fitness-go/internal/v2/service/review"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	receiptSvc := reviewService.NewService(db)
	uploadTool := uploader.NewReviewImageTool()
	return New(receiptSvc, uploadTool)
}
