package model

import "github.com/Henry19910227/fitness-go/internal/pkg/global"

type SaleItem struct {
	ID   int64 `gorm:"column:id"`  // 銷售id
	ProductLabelID int `gorm:"column:product_label_id"`  // 產品標籤id
	Type int `gorm:"column:type"`  // 銷售類型(1:免費課表/3:付費課表)
	Enable int `gorm:"column:enable"`  // 是否啟用
	Name string `gorm:"column:name"`  // 銷售名稱
	CreateAt string `gorm:"column:create_at"` //創建時間
	UpdateAt string `gorm:"column:update_at"` //更新時間
	ProductLabel *ProductLabel `gorm:"foreignKey:id;references:product_label_id"` // 產品標籤
}

func (SaleItem) TableName() string {
	return "sale_items"
}

type FindSaleItemsParam struct {
	Type *global.SaleType // 銷售類型(1:免費課表/3:付費課表)
}