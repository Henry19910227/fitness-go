package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
)

type Order struct {
	Base
	orderService service.Order
}

func NewOrder(baseGroup *gin.RouterGroup, orderService service.Order, userMiddleware middleware.User) {
	order := Order{
		orderService: orderService,
	}
	baseGroup.GET("/cms/user/:user_id/orders",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		order.GetCMSUserOrders)
}

// GetCMSUserOrders 獲取用戶購買歷史訂單
// @Summary 獲取用戶購買歷史訂單
// @Description 獲取用戶購買歷史訂單
// @Tags CMS/User_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "用戶ID"
// @Param order_field query string true "排序欄位 (create_at:創建時間)"
// @Param order_type query string true "排序類型 (ASC:由低到高/DESC:由高到低)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} model.SuccessResult{data=[]order.CMSUserOrdersAPI} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/cms/user/{user_id}/orders [GET]
func (o *Order) GetCMSUserOrders(c *gin.Context) {
	var uri validator.UserIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		o.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var orderByQuery validator.OrderByQuery
	if err := c.ShouldBind(&orderByQuery); err != nil {
		o.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var pagingQuery validator.PagingQuery
	if err := c.ShouldBind(&pagingQuery); err != nil {
		o.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	orders, paging, err := o.orderService.GetCMSUserOrders(c, *uri.UserID, &dto.OrderByParam{
		OrderField: orderByQuery.OrderField,
		OrderType:  orderByQuery.OrderType,
	}, &dto.PagingParam{
		Page: pagingQuery.Page,
		Size: pagingQuery.Size,
	})
	if err != nil {
		o.JSONErrorResponse(c, err)
		return
	}
	o.JSONSuccessPagingResponse(c, orders, paging, "success!")
}
