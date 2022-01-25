package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type sale struct {
	saleRepo          repository.Sale
	subscribePlanRepo repository.SubscribePlan
	jwtTool           tool.JWT
	errHandler        errcode.Handler
}

func NewSale(saleRepo repository.Sale, subscribePlanRepo repository.SubscribePlan, jwtTool tool.JWT, errHandler errcode.Handler) Sale {
	return &sale{saleRepo: saleRepo, subscribePlanRepo: subscribePlanRepo, jwtTool: jwtTool, errHandler: errHandler}
}

func (s *sale) GetSaleItems(c *gin.Context) ([]*dto.SaleItem, errcode.Error) {
	entities, err := s.saleRepo.FindSaleItems(nil)
	if err != nil {
		return nil, s.errHandler.Set(c, "sale item repo", err)
	}
	var saleItems []*dto.SaleItem
	for _, entity := range entities {
		saleItem := dto.SaleItem{
			ID:   entity.ID,
			Type: entity.Type,
			Name: entity.Name,
		}
		if entity.ProductLabel != nil {
			saleItem.ProductID = entity.ProductLabel.ProductID
			saleItem.Twd = entity.ProductLabel.Twd
		}
		saleItems = append(saleItems, &saleItem)
	}
	return saleItems, nil
}

func (s *sale) GetSubscribePlans(c *gin.Context) ([]*dto.SubscribePlan, errcode.Error) {
	models, err := s.subscribePlanRepo.FindSubscribePlans()
	if err != nil {
		return nil, s.errHandler.Set(c, "subscribe plan repo", err)
	}
	var subscribePlans []*dto.SubscribePlan
	for _, model := range models {
		subscribePlan := dto.SubscribePlan{
			ID:     model.ID,
			Period: model.Period,
			Name:   model.Name,
		}
		if model.ProductLabel != nil {
			subscribePlan.Twd = model.ProductLabel.Twd
			subscribePlan.ProductID = model.ProductLabel.ProductID
		}
		subscribePlans = append(subscribePlans, &subscribePlan)
	}
	return subscribePlans, nil
}
