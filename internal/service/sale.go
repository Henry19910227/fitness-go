package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/saledto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type sale struct {
	saleRepo repository.Sale
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewSale(saleRepo repository.Sale, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Sale {
	return &sale{saleRepo: saleRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (s *sale) GetSaleItems(c *gin.Context) ([]*saledto.SaleItem, errcode.Error) {
	entities, err := s.saleRepo.FinsSaleItems()
	if err != nil {
		s.logger.Set(c, handler.Error, "PlanRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	var saleItems []*saledto.SaleItem
	for _, entity := range entities {
		saleItem := saledto.SaleItem{
			ID: entity.ID,
			Type: entity.Type,
			Name: entity.Name,
			Twd: entity.Twd,
			Identifier: entity.Identifier,
			CreateAt: entity.CreateAt,
			UpdateAt: entity.UpdateAt,
		}
		saleItems = append(saleItems, &saleItem)
	}
	return saleItems, nil
}
