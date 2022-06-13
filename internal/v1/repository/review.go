package repository

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"gorm.io/gorm"
	"time"
)

type review struct {
	gorm tool.Gorm
}

func NewReview(gorm tool.Gorm) Review {
	return &review{gorm: gorm}
}

func (r *review) CreateReview(tx *gorm.DB, param *model.CreateReviewParam) (int64, error) {
	db := r.gorm.DB()
	if tx != nil {
		db = tx
	}
	review := entity.Review{
		CourseID: param.CourseID,
		UserID:   param.UserID,
		Score:    param.Score,
		Body:     param.Body,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	//創建一筆評論
	if err := db.Create(&review).Error; err != nil {
		return 0, err
	}
	return review.ID, nil
}

func (r *review) DeleteReview(tx *gorm.DB, reviewID int64) error {
	db := r.gorm.DB()
	if tx != nil {
		db = tx
	}
	//刪除該評論
	if err := db.
		Where("id = ?", reviewID).
		Delete(&entity.Review{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *review) FindReviewByID(tx *gorm.DB, reviewID int64) (*model.Review, error) {
	db := r.gorm.DB()
	if tx != nil {
		db = tx
	}
	var review model.Review
	if err := db.
		Preload("User").
		Preload("Images").
		Take(&review, reviewID).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *review) FindReviews(uid int64, param *model.FindReviewsParam, paging *model.PagingParam) ([]*model.Review, error) {
	if param == nil {
		return nil, nil
	}
	//將查詢會員的ID排在第一個
	orderQuery := fmt.Sprintf("reviews.user_id <> %v, reviews.create_at ASC", uid)
	var reviews []*model.Review
	db := r.gorm.DB().Preload("User").Preload("Images").Order(orderQuery)
	if param.FilterType == global.PhotoReviewType {
		db = db.Joins("INNER JOIN review_images ON review_images.review_id = reviews.id").Group("reviews.id")
	}
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	if err := db.Find(&reviews, "course_id = ?", param.CourseID).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *review) FindReviewsCount(param *model.FindReviewsParam) (int, error) {
	if param == nil {
		return 0, nil
	}
	db := r.gorm.DB().Table("reviews")
	if param.FilterType == global.PhotoReviewType {
		db = db.Joins("INNER JOIN review_images ON review_images.review_id = reviews.id").Group("reviews.id")
	}
	var count int64
	if err := db.Where("reviews.course_id = ?", param.CourseID).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
