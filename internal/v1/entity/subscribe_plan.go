package entity

type SubscribePlan struct {
	ID   int64 `gorm:"column:id"`  // 銷售id
	ProductLabelID int `gorm:"column:product_label_id"`  // 產品標籤id
	Period int `gorm:"column:period"`  // 週期(1:一個月/2:二個月/3:三個月/6:六個月/12:一年/99:永久)
	Enable int `gorm:"column:enable"`  // 是否啟用
	Name string `gorm:"column:name"` // 銷售名稱
	CreateAt string `gorm:"column:create_at"` //創建時間
	UpdateAt string `gorm:"column:update_at"` //更新時間
}

func (SubscribePlan) TableName() string {
	return "subscribe_plans"
}
