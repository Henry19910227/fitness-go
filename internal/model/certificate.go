package model

type Certificate struct {
	ID               int64   `gorm:"column:id"`              // 教練相片id
	UserID           int64   `gorm:"column:user_id"`         // 關聯的用戶id
	Name             string  `gorm:"column:name"`            // 證照名稱
	Image            string  `gorm:"column:image"`           // 照片
	CreateAt         string  `gorm:"column:create_at"`       // 創建日期
	UpdateAt         string  `gorm:"column:update_at"`       // 更新時間
}

func (Certificate) TableName() string {
	return "certificates"
}
