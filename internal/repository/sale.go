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

func (s *sale) FindSaleItems() ([]*model.SaleItem, error) {
	var entities []*model.SaleItem
	if err := s.gorm.DB().
		Table("sale_items").
		Select("*").
		Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (s *sale) FindSaleItemByID(saleID int64) (*model.SaleItem, error) {
	var item model.SaleItem
	if err := s.gorm.DB().
		Table("sale_items").
		Select("*").
		Where("id = ?", saleID).
		Take(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}


