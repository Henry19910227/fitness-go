package required

type UserIDField struct {
	UserID     int64  `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"`     // 關聯的用戶id
}
type CardIDField struct {
	CardID     string `json:"card_id" gorm:"column:card_id" binding:"required" example:"A123456789"`     // 身分證字號
}
type FrontImageField struct {
	FrontImage string `json:"front_image" gorm:"column:front_image" binding:"required" example:"123.jpg"` // 身分證正面照
}
type BackImageField struct {
	BackImage  string `json:"back_image" gorm:"column:back_image" binding:"required" example:"123.jpg"`  // 身分證背面照
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}

