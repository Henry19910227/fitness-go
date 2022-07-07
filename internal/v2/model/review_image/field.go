package review_image

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //圖片id
}
type ReviewIDField struct {
	ID *int64 `json:"review_id,omitempty" gorm:"column:review_id" example:"1"` //評論id
}
type ImageField struct {
	Image *string `json:"image,omitempty" gorm:"column:image" example:"1234.jpg"` //圖片
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}

type Table struct {
	IDField
	ReviewIDField
	ImageField
	CreateAtField
}

func (Table) TableName() string {
	return "review_images"
}
