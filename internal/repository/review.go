package repository

import (
	"fmt"
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

func (r *review) FindReviewByCourseIDAndUserID(courseID int64, userID int64) (*model.ReviewItem, error) {
	var review model.ReviewItem
	var user model.UserSummary
	if err := r.gorm.DB().
		Table("reviews").
		Select("reviews.course_id AS course_id", "reviews.score AS score", "reviews.body AS body", "reviews.create_at",
			"users.id AS user_id", "users.avatar AS avatar").
		Joins("INNER JOIN users ON reviews.user_id = users.id").
		Where("reviews.course_id = ? AND reviews.user_id = ?", courseID, userID).
		Row().
		Scan(&review.CourseID, &review.Score, &review.Body, &review.CreateAt,
			&user.ID, &user.Avatar); err != nil {
		return nil, err
	}
	reviewImages := make([]*model.ReviewImageItem, 0)
	if err := r.gorm.DB().Model(&model.ReviewImage{}).
		Where("course_id = ? AND user_id = ?", courseID, userID).
		Find(&reviewImages).Error; err != nil {
		return nil, err
	}
	review.User = &user
	review.Images = reviewImages
	return &review, nil
}

func (r *review) FindReviewsByCourseIDAndUserID(courseID int64, userID int64, paging *model.PagingParam) ([]*model.ReviewItem, error) {
	//將查詢會員的ID排在第一個
	orderQuery := fmt.Sprintf("reviews.user_id <> %v, reviews.create_at ASC", userID)
	db := r.gorm.DB().
		Table("reviews").
		Select("reviews.course_id AS course_id", "reviews.score AS score", "reviews.body AS body", "reviews.create_at",
			"users.id AS user_id", "users.avatar AS avatar").
		Joins("INNER JOIN users ON reviews.user_id = users.id").
		Where("reviews.course_id = ?", courseID).
		Order(orderQuery)
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	rows, err := db.Rows()
	if err != nil {
		return nil, err
	}
	reviewMap := make(map[int64]*model.ReviewItem)
	userIDs := make([]int64, 0)
	courseIDs := make([]int64, 0)
	reviews := make([]*model.ReviewItem, 0)
	for rows.Next() {
		var review model.ReviewItem
		var user model.UserSummary
		if err := rows.Scan(&review.CourseID, &review.Score, &review.Body, &review.CreateAt,
			&user.ID, &user.Avatar); err != nil {
			return nil, err
		}
		review.User = &user
		userIDs = append(userIDs, user.ID)
		courseIDs = append(courseIDs, review.CourseID)
		reviewMap[user.ID] = &review
		reviews = append(reviews, &review)
	}
	var reviewImages []*model.ReviewImageItem
	if err := r.gorm.DB().
		Model(&model.ReviewImage{}).
		Where("course_id IN (?) AND user_id IN (?)",courseIDs, userIDs).
		Find(&reviewImages).Error; err != nil {
			return nil, err
	}
	for _, image := range reviewImages {
		if review, ok := reviewMap[image.UserID]; ok {
			review.Images = append(review.Images, image)
		}
	}
	return reviews, nil
}
