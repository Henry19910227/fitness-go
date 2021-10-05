package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type review struct {
	gorm tool.Gorm
}

func NewReview(gorm tool.Gorm) Review {
	return &review{gorm: gorm}
}

func (r *review) CreateReview(param *model.CreateReviewParam) error {
	review := model.Review {
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
		var reviewImages []*model.ReviewImage
		for _, imageName := range param.ImageNames {
			reviewImage := model.ReviewImage{
				CourseID: param.CourseID,
				UserID: param.UserID,
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
		return err
	}
	return nil
}

func (r *review) FindReviewByCourseIDAndUserID(courseID int64, userID int64, entity interface{}) error {
	if err := r.gorm.DB().
		Model(&model.Review{}).
		Where("course_id = ? AND user_id = ?", courseID, userID).
		Take(entity).Error; err != nil{
		return err
	}
	return nil
}

func (r *review) FindReviewImagesByReviewID(courseID, userID int64, entity interface{}) error {
	if err := r.gorm.DB().
		Model(&model.ReviewImage{}).
		Where("course_id = ? AND user_id = ?", courseID, userID).
		Find(entity).Error; err != nil{
		return err
	}
	return nil
}
