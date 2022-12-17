package optional

type IDField struct {
	ID *string `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` // 訂單id
}
type VersionField struct {
	Version *string `json:"version,omitempty" gorm:"column:version" binding:"omitempty" example:"1.0.0"` // android 版本號
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
