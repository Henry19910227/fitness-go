package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type foodCategory struct {
	gorm tool.Gorm
}

func NewFoodCategory(gorm tool.Gorm) FoodCategory {
	return &foodCategory{gorm: gorm}
}

func (f *foodCategory) FindFoodCategories() ([]*model.FoodCategory, error) {
	results := make([]*model.FoodCategory, 0)
	if err := f.gorm.DB().
		Model(&entity.FoodCategory{}).
		Where("is_deleted = ?", 0).
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
