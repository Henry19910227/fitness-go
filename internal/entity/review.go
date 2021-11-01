package entity

type Review struct {
	ID int64 `gorm:"column:id primaryKey"` //評論id
	CourseID int64 `gorm:"column:course_id"` //課表id
	UserID int64 `gorm:"column:user_id"` //用戶id
	Score int `gorm:"column:score"` //評分
	Body string `gorm:"column:body"` //內容
	CreateAt string `gorm:"column:create_at"` //創建時間
}

func (Review) TableName() string {
	return "reviews"
}

type ReviewImage struct {
	ID int64 `gorm:"column:id"` //圖片id
	ReviewID int64 `gorm:"column:review_id"` //評論id
	Image string `gorm:"column:image"` //圖片
	CreateAt string `gorm:"column:create_at"` //創建時間
}

func (ReviewImage) TableName() string {
	return "review_images"
}