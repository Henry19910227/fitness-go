package dto

type Certificate struct {
	ID    int64  `json:"id" gorm:"column:id" example:"1"`                       // 證照id
	Name  string `json:"name" gorm:"column:name" example:"A級教練證照"`              // 證照名稱
	Image string `json:"image" gorm:"column:image" example:"dkf2se51fsdds.png"` // 證照照片
}
