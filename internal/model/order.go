package model

type Order struct {
	ID             string          `gorm:"column:id"`                         // 訂單id
	UserID         int64           `gorm:"column:user_id"`                    // 用戶id
	Quantity       int64           `gorm:"column:quantity"`                   // 數量
	OrderType      int             `gorm:"column:order_type"`                 // 訂單類型(1:課表購買/2:會員訂閱)
	OrderStatus    int             `gorm:"column:order_status"`               // 訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
	CreateAt       string          `gorm:"column:create_at"`                  // 創建時間
	UpdateAt       string          `gorm:"column:update_at"`                  // 更新時間
	OrderCourse    *OrderCourse    `gorm:"foreignKey:order_id;references:id"` // 訂單課表資訊
	OrderSubscribe *OrderSubscribe `gorm:"foreignKey:order_id;references:id"` // 訂單訂閱資訊
}

func (Order) TableName() string {
	return "orders"
}

type OrderCourse struct {
	OrderID    string                `gorm:"column:order_id"`                       // 訂單id
	SaleItemID int64                 `gorm:"column:sale_item_id"`                   // 購買項目id
	CourseID   int64                 `gorm:"column:course_id"`                      // 購買課表id
	SaleItem   *SaleItem             `gorm:"foreignKey:id;references:sale_item_id"` // 銷售項目
	Course     *CourseProductSummary `gorm:"foreignKey:id;references:course_id"`    // 課表
}

func (OrderCourse) TableName() string {
	return "order_courses"
}

type OrderSubscribe struct {
	OrderID         string         `gorm:"column:order_id"`                            // 訂單id
	SubscribePlanID int64          `gorm:"column:subscribe_plan_id"`                   // 訂閱方案id
	SubscribePlan   *SubscribePlan `gorm:"foreignKey:id;references:subscribe_plan_id"` // 課表
}

func (OrderSubscribe) TableName() string {
	return "order_subscribes"
}

type CreateOrderParam struct {
	UserID     int64  // 用戶id
	SaleItemID *int64 // 銷售項目id
	CourseID   int64  // 訂單課表id
}

type CreateSubscribeOrderParam struct {
	UserID          int64 // 用戶id
	SubscribePlanID int64 // 銷售項目id
}

type UpdateOrderParam struct {
	OrderStatus int `gorm:"column:order_status"` // 訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
}
