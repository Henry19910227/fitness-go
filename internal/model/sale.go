package model

type SaleItem struct {
	ID   int64 `gorm:"column:id"`  // 銷售id
	Type int64 `gorm:"column:type"`  // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Name string `gorm:"column:name"` // 銷售名稱
	Twd  float64 `gorm:"column:twd"` // 台幣價格
	Identifier string `gorm:"column:identifier"` // 銷售識別碼
	CreateAt string `gorm:"column:create_at"` //創建時間
	UpdateAt string `gorm:"column:update_at"` //更新時間
}

func (SaleItem) TableName() string {
	return "sale_items"
}

type SaleItemEntity struct {
	ID   int64 `gorm:"column:id"`  // 銷售id
	Type int64 `gorm:"column:type"`  // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Name string `gorm:"column:name"` // 銷售名稱
	Twd  float64 `gorm:"column:twd"` // 台幣價格
	Identifier string `gorm:"column:identifier"` // 銷售識別碼
	CreateAt string `gorm:"column:create_at"` //創建時間
	UpdateAt string `gorm:"column:update_at"` //更新時間
}

type SaleSummaryEntity struct {
	ID   int64   // 銷售id
	Type int64   // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Name string  // 銷售名稱
	Twd  float64 // 台幣價格
}
