package repository

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type subscribePlan struct {
	gorm tool.Gorm
}

func NewSubscribePlan(gorm tool.Gorm) SubscribePlan {
	return &subscribePlan{gorm: gorm}
}

func (s *subscribePlan) FindSubscribePlans() ([]*model.SubscribePlan, error) {
	var items []*model.SubscribePlan
	if err := s.gorm.DB().
		Preload("ProductLabel").
		Where("enable = ?", 1).
		Order("create_at ASC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (s *subscribePlan) FinsSubscribePlanByID(subscribePlanID int64) (*model.SubscribePlan, error) {
	var item *model.SubscribePlan
	if err := s.gorm.DB().
		Preload("ProductLabel").
		Where("id = ?", subscribePlanID).
		Take(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (s *subscribePlan) FindSubscribePlansByPeriod(period global.PeriodType) ([]*model.SubscribePlan, error) {
	var items []*model.SubscribePlan
	if err := s.gorm.DB().
		Preload("ProductLabel").
		Where("period = ? AND enable = ?", period, 1).
		Order("create_at ASC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
