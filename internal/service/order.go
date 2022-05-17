package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	orderDTO "github.com/Henry19910227/fitness-go/internal/dto/order"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/gin-gonic/gin"
)

type order struct {
	Base
	orderRepo  repository.Order
	errHandler errcode.Handler
}

func NewOrder(orderRepo repository.Order, errHandler errcode.Handler) Order {
	return &order{orderRepo: orderRepo, errHandler: errHandler}
}

func (o *order) GetCMSUserOrders(c *gin.Context, userID int64, orderByParam *dto.OrderByParam, pagingParam *dto.PagingParam) ([]*orderDTO.CMSUserOrdersAPI, *dto.Paging, errcode.Error) {
	//設置排序
	var orderBy *model.OrderBy
	if orderByParam != nil {
		orderBy = &model.OrderBy{
			OrderType: global.DESC,
			Field:     "update_at",
		}
		if orderByParam.OrderType != nil {
			orderBy.OrderType = global.OrderType(*orderByParam.OrderType)
		}
		if orderByParam.OrderField != nil {
			orderBy.Field = *orderByParam.OrderField
		}
	}
	//設置分頁
	var paging *model.PagingParam
	if pagingParam != nil {
		offset, limit := o.GetPagingIndex(pagingParam.Page, pagingParam.Size)
		paging = &model.PagingParam{
			Offset: offset,
			Limit:  limit,
		}
	}
	orders := make([]*orderDTO.CMSUserOrdersAPI, 0)
	amount, err := o.orderRepo.FindCMSUserOrdersAPIItems(userID, &orders, orderBy, paging)
	if err != nil {
		return nil, nil, o.errHandler.Set(c, "order repo", err)
	}
	pagingResult := dto.Paging{
		TotalCount: amount,
		TotalPage:  o.GetTotalPage(amount, pagingParam.Size),
		Page:       pagingParam.Page,
		Size:       pagingParam.Size,
	}
	return orders, &pagingResult, nil
}
