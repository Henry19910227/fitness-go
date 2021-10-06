package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Sale struct {
	Base
	saleService service.Sale
}

func NewSale(baseGroup *gin.RouterGroup, saleService service.Sale, userMiddleware gin.HandlerFunc) {
	sale := Sale{
		saleService: saleService,
	}
	saleItemsGroup := baseGroup.Group("/sale_items")
	saleItemsGroup.Use(userMiddleware)
	saleItemsGroup.GET("", sale.GetSaleItems)
}

// GetSaleItems 取得銷售項目清單
// @Summary  取得銷售項目清單
// @Description  取得銷售項目清單
// @Tags Sale
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=[]dto.SaleItem} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /sale_items [GET]
func (s *Sale) GetSaleItems(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		s.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	saleItems, err := s.saleService.GetSaleItems(c)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, saleItems, "get success!")
}
