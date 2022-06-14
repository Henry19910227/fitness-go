package dto

import "github.com/Henry19910227/fitness-go/internal/v1/model"

type CourseOrder struct {
	ID          string                `json:"id" example:"202105201300687423"`         // 訂單id
	UserID      int64                 `json:"user_id" example:"10001"`                 // 用戶id
	Quantity    int64                 `json:"quantity" example:"1"`                    // 數量
	OrderType   int                   `json:"order_type" example:"1"`                  // 訂單類型(1:課表購買/2:會員訂閱)
	OrderStatus int                   `json:"order_status" example:"1"`                // 訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
	SaleItem    *SaleItem             `json:"sale_item"`                               // 銷售項目
	Course      *CourseProductSummary `json:"course"`                                  // 訂單課表
	CreateAt    string                `json:"create_at" example:"2021-05-28 11:00:00"` // 創建時間
	UpdateAt    string                `json:"update_at" example:"2021-05-28 11:00:00"` // 更新時間
}

type SubscribeOrder struct {
	ID            string         `json:"id" example:"202105201300687423"`         // 訂單id
	UserID        int64          `json:"user_id" example:"10001"`                 // 用戶id
	Quantity      int64          `json:"quantity" example:"1"`                    // 數量
	OrderType     int            `json:"order_type" example:"1"`                  // 訂單類型(1:課表購買/2:會員訂閱)
	OrderStatus   int            `json:"order_status" example:"1"`                // 訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
	SubscribePlan *SubscribePlan `json:"subscribe_plan"`                          // 銷售項目
	CreateAt      string         `json:"create_at" example:"2021-05-28 11:00:00"` // 創建時間
	UpdateAt      string         `json:"update_at" example:"2021-05-28 11:00:00"` // 更新時間
}

func NewSubscribeOrder(data *model.Order) SubscribeOrder {
	order := SubscribeOrder{
		ID:          data.ID,
		UserID:      data.UserID,
		Quantity:    data.Quantity,
		OrderType:   data.OrderType,
		OrderStatus: data.OrderStatus,
		CreateAt:    data.CreateAt,
		UpdateAt:    data.UpdateAt,
	}
	if data.OrderSubscribe != nil {
		if data.OrderSubscribe.SubscribePlan != nil {
			order.SubscribePlan = &SubscribePlan{
				ID:     data.OrderSubscribe.SubscribePlan.ID,
				Period: data.OrderSubscribe.SubscribePlan.Period,
			}
			if data.OrderSubscribe.SubscribePlan.ProductLabel != nil {
				order.SubscribePlan.ProductID = data.OrderSubscribe.SubscribePlan.ProductLabel.ProductID
				order.SubscribePlan.Name = data.OrderSubscribe.SubscribePlan.ProductLabel.Name
				order.SubscribePlan.Twd = data.OrderSubscribe.SubscribePlan.ProductLabel.Twd
			}
		}
	}
	return order
}

type CreateOrderParam struct {
	UserID      int64  // 用戶id
	SaleItemID  int64  // 銷售項目id
	OrderType   int    // 訂單類型(1:課表購買/2:會員訂閱)
	PaymentType int    // 支付方式(1:apple內購/2:google內購)
	CourseID    *int64 // 訂單課表id
}
