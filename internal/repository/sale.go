package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type sale struct {
	gorm  tool.Gorm
}

func NewSale(gorm tool.Gorm) Sale {
	return &sale{gorm: gorm}
}

func (s *sale) FinsSaleItems() ([]*model.SaleItemEntity, error) {
	var entities []*model.SaleItemEntity
	if err := s.gorm.DB().
		Table("sale_items").
		Select("*").
		Find(&entities).Error; err != nil {
			return nil, err
	}
	return entities, nil
}

