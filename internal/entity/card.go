package entity

type Card struct {
	UserID     int64  `gorm:"column:user_id"`     // 關聯的用戶id
	CardID     string `gorm:"column:card_id"`     // 身分證字號
	FrontImage string `gorm:"column:front_image"` // 身分證正面照
	BackImage  string `gorm:"column:back_image"`  // 身分證背面照
	CreateAt   string `gorm:"column:create_at"`   // 創建日期
	UpdateAt   string `gorm:"column:update_at"`   // 更新時間
}

func (Card) TableName() string {
	return "cards"
}
