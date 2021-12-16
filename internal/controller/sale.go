package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
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
	baseGroup.GET("/course_sale_items",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		sale.GetCourseSaleItems)

	baseGroup.GET("/sale_items",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		sale.GetCourseSaleItems)
}

// GetCourseSaleItems 取得付費課表銷售項目清單
// @Summary  取得付費課表銷售項目清單
// @Description  取得付費課表銷售項目清單
// @Tags Sale
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=[]dto.SaleItem} "獲取成功!"
// @Failure 400 {object} model.ErrorResult "獲取失敗"
// @Router /course_sale_items [GET]
func (s *Sale) GetCourseSaleItems(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		s.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	saleItems, err := s.saleService.GetCourseSaleItems(c)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, saleItems, "get success!")
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
