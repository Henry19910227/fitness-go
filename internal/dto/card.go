package dto

type Card struct {
	FrontImage string `json:"front_image" gorm:"column:front_image" example:"sd2lkd1e23w54dw3e.png"` // 身分證正面照
	BackImage  string `json:"back_image" gorm:"column:back_image" example:"sd2lkd1e23w54dw3e.png"`   // 身分證背面照
}
