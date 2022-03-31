package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type actionPR struct {
	gorm tool.Gorm
}

func NewActionPR(gorm tool.Gorm) ActionPR {
	return &actionPR{gorm: gorm}
}

func (a *actionPR) FindActionPRs(userID int64, actionIDs []int64) ([]*model.ActionPR, error) {
	actionPRs := make([]*model.ActionPR, 0)
	if err := a.gorm.DB().
		Table("action_personal_records").
		Where("user_id = ? AND action_id IN (?)", userID, actionIDs).
		Find(&actionPRs).Error; err != nil {
		return nil, err
	}
	return actionPRs, nil
}

func (a *actionPR) SaveActionPRs(tx *gorm.DB, userID int64, params []*model.CreateActionPRParam) error {
	if len(params) == 0 {
		return nil
	}
	db := a.gorm.DB()
	if tx != nil {
		db = tx
	}
	actionPRs := make([]*entity.ActionPR, 0)
	for _, param := range params {
		actionPR := entity.ActionPR{
			UserID:   userID,
			ActionID: param.ActionID,
			Weight:   param.Weight,
			Reps:     param.Reps,
			Distance: param.Distance,
			Duration: param.Duration,
			Incline:  param.Incline,
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		actionPRs = append(actionPRs, &actionPR)
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "action_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"weight", "reps", "distance", "duration", "incline", "update_at"}),
	}).Create(&actionPRs).Error; err != nil {
		return err
	}
	return nil
}
