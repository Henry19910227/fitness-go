package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
)

type sale struct {
	gorm tool.Gorm
}

func NewSale(gorm tool.Gorm) Sale {
	return &sale{gorm: gorm}
}

func (s *sale) FindSaleItems(saleType *int) ([]*model.SaleItem, error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 user_id 篩選條件
	if saleType != nil {
		query += "AND type = ? "
		params = append(params, saleType)
	}
	var entities []*model.SaleItem
	if err := s.gorm.DB().
		Preload("ProductLabel").
		Where(query, params...).
		Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (s *sale) FindSaleItemByID(saleID int64) (*model.SaleItem, error) {
	var item model.SaleItem
	if err := s.gorm.DB().
		Preload("ProductLabel").
		Where("id = ?", saleID).
		Take(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
