package optional

type UserIDField struct {
	UserID     *int64  `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"`     // 關聯的用戶id
}
type CardIDField struct {
	CardID     *string `json:"card_id,omitempty" gorm:"column:card_id" binding:"omitempty" example:"A123456789"`     // 身分證字號
}
type FrontImageField struct {
	FrontImage *string `json:"front_image,omitempty" gorm:"column:front_image" binding:"omitempty" example:"123.jpg"` // 身分證正面照
}
type BackImageField struct {
	BackImage  *string `json:"back_image,omitempty" gorm:"column:back_image" binding:"omitempty" example:"123.jpg"`  // 身分證背面照
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
