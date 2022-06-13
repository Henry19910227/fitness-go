package entity

type PurchaseLog struct {
	ID int64 `gorm:"column:id primaryKey"` // id
	UserID int64 `gorm:"column:user_id"` // 用戶id
	OrderID string `gorm:"column:order_id"` // 訂單id
	Type int `gorm:"column:type"` // 訂單類型(1:購買/2:退費)
	CreateAt string `gorm:"create_at"` // 創建時間
}

func (PurchaseLog) TableName() string {
	return "purchase_logs"
}