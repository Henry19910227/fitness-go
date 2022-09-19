package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` // 銷售id
}
type ProductLabelIDField struct {
	ProductLabelID int `json:"product_label_id" gorm:"column:product_label_id" binding:"required" example:"2"` // 產品標籤id
}
type TypeField struct {
	Type int `json:"type" gorm:"column:type" binding:"required" example:"3"` // 銷售類型(1:免費課表/3:付費課表)
}
type EnableField struct {
	Enable int `json:"enable" gorm:"column:enable" binding:"required" example:"1"` // 是否啟用
}
type NameField struct {
	Name string `json:"name" gorm:"column:name" binding:"required" example:"銅級課表 "` // 銷售名稱
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-12 00:00:00"` // 更新時間
}
