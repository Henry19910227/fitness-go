package review

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/review"
	"github.com/Henry19910227/fitness-go/internal/v2/service/review_image"
	"github.com/Henry19910227/fitness-go/internal/v2/service/review_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer_statistic"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	reviewService := review.NewService(db)
	reviewImageService := review_image.NewService(db)
	reviewStatisticService := review_statistic.NewService(db)
	trainerStatisticService := trainer_statistic.NewService(db)
	courseService := course.NewService(db)
	uploadTool := uploader.NewReviewImageTool()
	return New(reviewService, reviewImageService, reviewStatisticService, trainerStatisticService, courseService, uploadTool)
}
