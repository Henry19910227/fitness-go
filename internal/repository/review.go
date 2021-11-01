package repository

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type review struct {
	gorm tool.Gorm
}

func NewReview(gorm tool.Gorm) Review {
	return &review{gorm: gorm}
}

func (r *review) CreateReview(param *model.CreateReviewParam) (int64, error) {
	review := entity.Review{
		CourseID: param.CourseID,
		UserID: param.UserID,
		Score: param.Score,
		Body: param.Body,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := r.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//創建一筆評論
		if err := tx.Create(&review).Error; err != nil {
			return err
		}
		//創建評論照片
		var reviewImages []*entity.ReviewImage
		for _, imageName := range param.ImageNames {
			reviewImage := entity.ReviewImage{
				ReviewID: review.ID,
				Image: imageName,
				CreateAt: time.Now().Format("2006-01-02 15:04:05"),
			}
			reviewImages = append(reviewImages, &reviewImage)
		}
		if reviewImages != nil {
			if err := tx.Create(reviewImages).Error; err != nil {
				return err
			}
		}
		//查詢並修改當前評論統計狀態
		var reviewStat *model.ReviewStatistic
		if err := tx.Table("review_statistics").
			Where("course_id = ? FOR UPDATE", param.CourseID).
			Find(&reviewStat).Error; err != nil {
				return err
		}
		reviewStat.CourseID = param.CourseID
		reviewStat.ScoreTotal += param.Score
		reviewStat.Amount += 1
		reviewStat.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
		switch param.Score {
		case 1:
			reviewStat.OneTotal += 1
		case 2:
			reviewStat.TwoTotal += 1
		case 3:
			reviewStat.ThreeTotal += 1
		case 4:
			reviewStat.FourTotal += 1
		case 5:
			reviewStat.FiveTotal += 1
		}
		//沒有有評論統計紀錄
		if reviewStat.CourseID == 0 {
			reviewStat.CourseID = param.CourseID
			if err := tx.Create(&reviewStat).Error; err != nil {
				return err
			}
			return nil
		}
		//已經有評論統計紀錄
		if err := tx.Where("course_id = ?", reviewStat.CourseID).Save(&reviewStat).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return review.ID, nil
}

func (r *review) DeleteReview(reviewID int64) error {
	if err := r.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//查詢當前評論狀態
		var review model.Review
		if err := tx.Table("reviews").
			Where("id = ?", reviewID).
			Find(&review).Error; err != nil {
				return err
		}
		//刪除該評論
		if err := tx.Delete(&review).Error; err != nil {
			return err
		}
		//修改當前評論統計狀態
		var reviewStat *model.ReviewStatistic
		if err := tx.Table("review_statistics").
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("course_id = ?", review.CourseID).
			Find(&reviewStat).Error; err != nil {
			return err
		}
		reviewStat.ScoreTotal -= review.Score
		reviewStat.Amount -= 1
		reviewStat.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
		switch review.Score {
		case 1:
			reviewStat.OneTotal -= 1
		case 2:
			reviewStat.TwoTotal -= 1
		case 3:
			reviewStat.ThreeTotal -= 1
		case 4:
			reviewStat.FourTotal -= 1
		case 5:
			reviewStat.FiveTotal -= 1
		}
		if err := tx.Where("course_id = ?", reviewStat.CourseID).Save(&reviewStat).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *review) FindReviewByID(reviewID int64) (*model.Review, error) {
	var review model.Review
	if err := r.gorm.DB().
		Preload("User").
		Preload("Images").
		Take(&review, reviewID).Error; err != nil {
			return nil, err
	}
	return &review, nil
}

func (r *review) FindReviewsByCourseID(courseID int64, uid int64, paging *model.PagingParam) ([]*model.Review, error) {
	//將查詢會員的ID排在第一個
	orderQuery := fmt.Sprintf("reviews.user_id <> %v, reviews.create_at ASC", uid)
	var reviews []*model.Review
	db := r.gorm.DB().
		Preload("User").
		Preload("Images").
		Order(orderQuery)
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	if err := db.Find(&reviews, "course_id = ?", courseID).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *review) FindReviewImages(courseID int64, userID int64) ([]*model.ReviewImageItem, error) {
	var reviewImages []*model.ReviewImageItem
	if err := r.gorm.DB().Table("review_images").
		Where("course_id = ? AND user_id = ?",courseID, userID).
		Find(&reviewImages).Error; err != nil {
		return nil, err
	}
	return reviewImages, nil
}
