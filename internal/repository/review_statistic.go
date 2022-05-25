package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type reviewStatistic struct {
	gorm tool.Gorm
}

func NewReviewStatistic(gorm tool.Gorm) ReviewStatistic {
	return &reviewStatistic{gorm: gorm}
}

func (r *reviewStatistic) FindReviewStatisticOutput(courseID int64, output interface{}) error {
	if err := r.gorm.DB().
		Model(&model.ReviewStatistic{}).
		Where("course_id = ?", courseID).
		Take(output).Error; err != nil {
		return err
	}
	return nil
}

func (r *reviewStatistic) SaveReviewStatistic(tx *gorm.DB, courseID int64, param *model.SaveReviewStatisticParam) error {
	db := r.gorm.DB()
	if tx != nil {
		db = tx
	}
	stat := entity.ReviewStatistic{
		CourseID:   courseID,
		ScoreTotal: param.ScoreTotal,
		Amount:     param.Amount,
		FiveTotal:  param.FiveTotal,
		FourTotal:  param.FourTotal,
		ThreeTotal: param.ThreeTotal,
		TwoTotal:   param.TwoTotal,
		OneTotal:   param.OneTotal,
		UpdateAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "course_id"}},
		DoUpdates: clause.AssignmentColumns(
			[]string{"score_total", "amount", "five_total", "four_total", "three_total", "two_total", "one_total", "update_at"}),
	}).Create(&stat).Error; err != nil {
		return err
	}
	return nil
}

func (r *reviewStatistic) CalculateReviewStatistic(tx *gorm.DB, courseID int64) (*model.ReviewStatistic, error) {
	db := r.gorm.DB()
	if tx != nil {
		db = tx
	}
	reviews := make([]*entity.Review, 0)
	if err := db.
		Table("reviews").
		Select("*").
		Where("course_id = ?", courseID).
		Find(&reviews).Error; err != nil {
		return nil, err
	}
	reviewStat := model.ReviewStatistic{
		CourseID: courseID,
	}
	for _, item := range reviews {
		reviewStat.ScoreTotal += item.Score
		reviewStat.Amount += 1
		switch item.Score {
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
	}
	return &reviewStat, nil
}
