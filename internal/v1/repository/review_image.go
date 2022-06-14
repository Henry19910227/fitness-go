package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"gorm.io/gorm"
	"time"
)

type reviewImage struct {
	gorm tool.Gorm
}

func NewReviewImage(gorm tool.Gorm) ReviewImage {
	return &reviewImage{gorm: gorm}
}

func (r *reviewImage) CreateReviewImages(tx *gorm.DB, reviewID int64, imageNames []string) error {
	if len(imageNames) == 0 || imageNames == nil {
		return nil
	}
	db := r.gorm.DB()
	if tx != nil {
		db = tx
	}
	var reviewImages []*entity.ReviewImage
	for _, imageName := range imageNames {
		reviewImage := entity.ReviewImage{
			ReviewID: reviewID,
			Image:    imageName,
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		reviewImages = append(reviewImages, &reviewImage)
	}
	if err := db.Create(reviewImages).Error; err != nil {
		return err
	}
	return nil
}
