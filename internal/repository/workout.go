package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type workout struct {
	gorm tool.Gorm
}

func NewWorkout(gorm tool.Gorm) Workout {
	return &workout{gorm: gorm}
}

func (w *workout) CreateWorkout(planID int64, name string) (int64, error) {
	workout := model.Workout{
		PlanID: planID,
		Name: name,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := w.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//創建訓練
		if err := tx.Create(&workout).Error; err != nil {
			return err
		}
		//查詢關聯課表的計畫數量
		var workoutCount int
		if err := tx.
			Raw("SELECT COUNT(*) FROM workouts WHERE plan_id = ? FOR UPDATE", planID).
			Scan(&workoutCount).Error; err != nil {
			return err
		}
		//更新計畫擁有的訓練數量
		if err := tx.
			Table("plans").
			Where("id = ?", planID).
			Update("workout_count", workoutCount).Error; err != nil {
			return err
		}
		//取得關聯課表id
		var courseID int64
		if err := tx.
			Table("plans").
			Select("course_id").
			Where("id = ?", planID).
			Take(&courseID).Error; err != nil {
			return err
		}
		//更新課表擁有的訓練數量
		if err := tx.
			Table("courses").
			Where("id = ?", courseID).
			Update("workout_count", workoutCount).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return workout.ID, nil
}
