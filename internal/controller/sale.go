package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/gin-gonic/gin"
)

type Sale struct {
	Base
	saleService service.Sale
}

func NewSale(baseGroup *gin.RouterGroup, saleService service.Sale, userMidd midd.User) {
	sale := Sale{
		saleService: saleService,
	}
	baseGroup.GET("/sale_items",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		sale.GetSaleItems)

	baseGroup.GET("/subscribe_plans",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		sale.GetSubscribePlans)
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
	saleItems, err := s.saleService.GetSaleItems(c)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, saleItems, "get success!")
}

// GetSubscribePlans 取得訂閱方案清單
// @Summary  取得訂閱方案清單
// @Description  取得訂閱方案清單
// @Tags Sale
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=[]dto.SubscribePlan} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /subscribe_plans [GET]
func (s *Sale) GetSubscribePlans(c *gin.Context) {
	subscribePlans, err := s.saleService.GetSubscribePlans(c)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, subscribePlans, "get success!")
}
