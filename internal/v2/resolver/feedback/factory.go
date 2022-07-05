package feedback

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	feedbackService "github.com/Henry19910227/fitness-go/internal/v2/service/feedback"
	feedbackImageService "github.com/Henry19910227/fitness-go/internal/v2/service/feedback_image"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	feedbackSvc := feedbackService.NewService(db)
	feedbackImageSvc := feedbackImageService.NewService(db)
	uploadTool := uploader.NewFeedbackImageTool()
	return New(feedbackSvc, feedbackImageSvc, uploadTool)
}
