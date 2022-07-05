package feedback_image

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //主鍵id
}
type FeedbackIDField struct {
	FeedbackID *int64 `json:"feedback_id,omitempty" gorm:"column:feedback_id" example:"1"` //反饋id
}
type ImageField struct {
	Image *string `json:"image,omitempty" gorm:"column:image" example:"123.png"` //反饋圖片
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}

type Table struct {
	IDField
	FeedbackIDField
	ImageField
	CreateAtField
}

func (Table) TableName() string {
	return "feedback_images"
}
