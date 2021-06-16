package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type plan struct {
	gorm  tool.Gorm
}

func NewPlan(gorm tool.Gorm) Plan {
	return &plan{gorm: gorm}
}

func (p *plan) CreatePlan(courseID int64, name string) (int64, error) {
	plan := model.Plan{
		CourseID: courseID,
		Name: name,
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
		return err
	}
	return nil
}

func (p *plan) CheckPlanExistByUID(uid int64, planID int64) (bool, error) {
	var result int
	if err := p.gorm.DB().
		Table("plans").
		Select("1").
		Joins("INNER JOIN courses ON plans.course_id = courses.id ").
		Joins("INNER JOIN users ON courses.user_id = users.id ").
		Where("plans.id = ? AND users.id = ?", planID, uid).
		Find(&result).Error; err != nil {
		return false, err
	}
	return result > 0, nil
}
