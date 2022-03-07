package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type userPlanStatistic struct {
	gorm tool.Gorm
}

func NewUserPlanStatistic(gorm tool.Gorm) UserPlanStatistic {
	return &userPlanStatistic{gorm: gorm}
}

func (p *userPlanStatistic) FindUserPlanStatistics(userID int64, planID int64) ([]*model.UserPlanStatistic, error) {
	items := make([]*model.UserPlanStatistic, 0)
	if err := p.gorm.DB().
		Where("user_id = ? AND plan_id = ?", userID, planID).
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (p *userPlanStatistic) SaveUserPlanStatistic(tx *gorm.DB, param *model.SaveUserPlanStatisticParam) (int64, error) {
	db := p.gorm.DB()
	if tx != nil {
		db = tx
	}
	planStatistic := entity.UserPlanStatistic{
		UserID:             param.UserID,
		PlanID:             param.PlanID,
		Duration:           param.Duration,
		FinishWorkoutCourt: param.FinishWorkoutCount,
		UpdateAt:           time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "plan_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"finish_workout_count", "duration", "update_at"}),
	}).Create(&planStatistic).Error; err != nil {
		return 0, err
	}
	return planStatistic.ID, nil
}
