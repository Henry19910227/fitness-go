package entity

type Order struct {
	ID          string `gorm:"column:id"`           // 訂單id
	UserID      int64  `gorm:"column:user_id"`      // 用戶id
	Quantity    int64  `gorm:"column:quantity"`     // 數量
	OrderType   int    `gorm:"column:order_type"`   // 訂單類型(1:課表購買/2:會員訂閱)
	OrderStatus int    `gorm:"column:order_status"` // 訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
	CreateAt    string `gorm:"column:create_at"`    // 創建時間
	UpdateAt    string `gorm:"column:update_at"`    // 更新時間
}

func (Order) TableName() string {
	return "orders"
}

type OrderTemplate struct {
	ID                 string                      `gorm:"column:id"`                         // 訂單id
	UserID             int64                       `gorm:"column:user_id"`                    // 用戶id
	Quantity           int64                       `gorm:"column:quantity"`                   // 數量
	OrderType          int                         `gorm:"column:order_type"`                 // 訂單類型(1:課表購買/2:會員訂閱)
	OrderStatus        int                         `gorm:"column:order_status"`               // 訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
	CreateAt           string                      `gorm:"column:create_at"`                  // 創建時間
	UpdateAt           string                      `gorm:"column:update_at"`                  // 更新時間
	OrderCourse        *OrderCourseTemplate        `gorm:"foreignKey:order_id;references:id"` // 訂單課表資訊
	OrderSubscribePlan *OrderSubscribePlanTemplate `gorm:"foreignKey:order_id;references:id"` // 訂單訂閱資訊
}

func (OrderTemplate) TableName() string {
	return "orders"
}
