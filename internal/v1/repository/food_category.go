package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
)

type foodCategory struct {
	gorm tool.Gorm
}

func NewFoodCategory(gorm tool.Gorm) FoodCategory {
	return &foodCategory{gorm: gorm}
}

func (f *foodCategory) FindFoodCategory(categoryID int64) (*model.FoodCategory, error) {
	var foodCategory model.FoodCategory
	if err := f.gorm.DB().
		Model(&entity.FoodCategory{}).
		Where("id = ?", categoryID).
		Take(&foodCategory).Error; err != nil {
		return nil, err
	}
	return &foodCategory, nil
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
