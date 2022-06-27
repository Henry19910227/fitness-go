package body_image

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //主鍵id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}
type BodyImageField struct {
	BodyImage *string `json:"body_image,omitempty" gorm:"column:body_image" example:"1234.jpg"` //體態照片
}
type WeightField struct {
	Weight *float64 `json:"weight,omitempty" gorm:"weight:value" example:"50.5"` //體重(公斤)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	UserIDField
	BodyImageField
	WeightField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "body_images"
}