package dto

type Certificate struct {
	ID  int64   `json:"id" gorm:"column:id" example:"1"` // 證照id
	Name  string  `json:"name" gorm:"column:name" example:"A級教練證照"` // 證照名稱
	Image string `json:"image" example:"dkf2se51fsdds.png"` // 證照照片
	CreateAt  string   `json:"create_at" gorm:"column:create_at" example:"2021-06-01 12:00:00"`  // 創建日期
	UpdateAt  string   `json:"update_at" gorm:"column:update_at" example:"2021-06-01 12:00:00"`  // 修改日期
}
