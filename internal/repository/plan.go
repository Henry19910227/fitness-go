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

func (p *plan) CheckPlanExistByUID(uid int64, planID int64) (bool, error) {
	panic("implement me")
}
