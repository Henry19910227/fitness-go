package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type plan struct {
	gorm tool.Gorm
}

func NewPlan(gorm tool.Gorm) Plan {
	return &plan{gorm: gorm}
}

func (p *plan) CreatePlan(courseID int64, name string) (int64, error) {
	plan := model.Plan{
		CourseID: courseID,
		Name:     name,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := p.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//創建計畫
		if err := tx.Create(&plan).Error; err != nil {
			return err
		}
		//查詢關聯課表的計畫數量
		var planCount int
		if err := tx.
			Raw("SELECT COUNT(*) FROM plans WHERE course_id = ? FOR UPDATE", courseID).
			Scan(&planCount).Error; err != nil {
			return err
		}
		//更新課表擁有的計畫數量
		if err := tx.
			Table("courses").
			Where("id = ?", courseID).
			Update("plan_count", planCount).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return plan.ID, nil
}

func (p *plan) FindPlanByID(planID int64, entity interface{}) error {
	if err := p.gorm.DB().
		Model(&model.Plan{}).
		Where("id = ?", planID).
		Find(entity).Error; err != nil {
		return err
	}
	return nil
}

func (p *plan) FindPlansByCourseID(courseID int64) ([]*model.Plan, error) {
	plans := make([]*model.Plan, 0)
	if err := p.gorm.DB().
		Table("plans").
		Select("*").
		Where("course_id = ?", courseID).
		Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

func (p *plan) FindPlanAssets(userID int64, courseID int64) ([]*model.PlanAsset, error) {
	plans := make([]*model.PlanAsset, 0)
	if err := p.gorm.DB().
		Table("plans").
		Select("plans.id AS id", "plans.course_id AS course_id",
			"plans.name AS name", "plans.workout_count AS workout_count",
			"IFNULL(stat.finish_workout_count, 0) AS finish_workout_count",
			"plans.create_at", "plans.update_at").
		Joins("INNER JOIN user_plan_statistics AS stat ON plans.id = stat.plan_id AND stat.user_id = ?", userID).
		Where("plans.course_id = ?", courseID).
		Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

func (p *plan) UpdatePlanByID(planID int64, name string) error {
	if err := p.gorm.DB().
		Table("plans").
		Where("id = ?", planID).
		Updates(map[string]interface{}{
			"name":      name,
			"update_at": time.Now().Format("2006-01-02 15:04:05"),
		}).Error; err != nil {
		return err
	}
	return nil
}

func (p *plan) DeletePlanByID(planID int64) error {
	if err := p.gorm.DB().Transaction(func(tx *gorm.DB) error {
		//查詢關聯課表id
		var courseID int64
		if err := tx.
			Table("plans").
			Select("course_id").
			Where("id = ?", planID).
			Take(&courseID).Error; err != nil {
			return err
		}
		//刪除計畫
		if err := tx.
			Where("id = ?", planID).
			Delete(&model.Plan{}).Error; err != nil {
			return err
		}
		//查詢關聯課表的訓練數量
		var workoutCount int
		if err := tx.
			Raw("SELECT COUNT(*) FROM courses "+
				"INNER JOIN plans ON courses.id = plans.course_id "+
				"INNER JOIN workouts ON plans.id = workouts.plan_id "+
				"WHERE course_id = ? FOR UPDATE", courseID).
			Scan(&workoutCount).Error; err != nil {
			return err
		}
		//查詢關聯課表的計畫數量
		var planCount int
		if err := tx.
			Raw("SELECT COUNT(*) FROM plans WHERE course_id = ? FOR UPDATE", courseID).
			Scan(&planCount).Error; err != nil {
			return err
		}
		//更新課表擁有的計畫與訓練數量
		if err := tx.
			Table("courses").
			Where("id = ?", courseID).
			Updates(map[string]interface{}{
				"plan_count":    planCount,
				"workout_count": workoutCount,
			}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (p *plan) FindPlanOwnerByID(planID int64) (int64, error) {
	var userID int64
	if err := p.gorm.DB().
		Table("plans").
		Select("courses.user_id").
		Joins("INNER JOIN courses ON plans.course_id = courses.id").
		Where("plans.id = ?", planID).
		Take(&userID).Error; err != nil {
		return 0, err
	}
	return userID, nil
}
