package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` //id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` //用戶id
}
type NameField struct {
	Name *string `json:"name,omitempty" gorm:"column:name" binding:"omitempty" example:"A級教練證照"` //證照名稱
}
type ImageField struct {
	Image *string `json:"image,omitempty" gorm:"column:image" binding:"omitempty" example:"1234.jpg"` //圖片
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
