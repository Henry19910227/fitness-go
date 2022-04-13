package entity

type ReviewImage struct {
	ID       int64  `gorm:"column:id"`        //圖片id
	ReviewID int64  `gorm:"column:review_id"` //評論id
	Image    string `gorm:"column:image"`     //圖片
	CreateAt string `gorm:"column:create_at"` //創建時間
}

func (ReviewImage) TableName() string {
	return "review_images"
}
