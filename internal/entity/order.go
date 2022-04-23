package entity

type Order struct {
	ID          string `gorm:"column:id"`           // 訂單id
	UserID      int64 `gorm:"column:user_id"`      // 用戶id
	Quantity    int64  `gorm:"column:quantity"`     // 數量
	OrderType   int    `gorm:"column:order_type"`   // 訂單類型(1:課表購買/2:會員訂閱)
	OrderStatus int    `gorm:"column:order_status"` // 訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
	CreateAt    string `gorm:"column:create_at"`    // 創建時間
	UpdateAt    string `gorm:"column:update_at"`    // 更新時間
}

func (Order) TableName() string {
	return "orders"
}

type OrderCourse struct {
	OrderID    string `gorm:"column:order_id"`     // 訂單id
	SaleItemID *int64 `gorm:"column:sale_item_id"` // 銷售項目 id
	CourseID   int64  `gorm:"column:course_id"`    // 課表id
}

func (OrderCourse) TableName() string {
	return "order_courses"
}
