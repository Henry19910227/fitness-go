package repository

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type subscribePlan struct {
	gorm  tool.Gorm
}

func NewSubscribePlan(gorm tool.Gorm) SubscribePlan {
	return &subscribePlan{gorm: gorm}
}

func (s *subscribePlan) FindSubscribePlansByPeriod(period global.PeriodType) ([]*model.SubscribePlan, error) {
	var items []*model.SubscribePlan
	if err := s.gorm.DB().
		Preload("ProductLabel").
		Where("period = ?", period).
		Order("create_at ASC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
