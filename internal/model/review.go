package model

type Review struct {
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
	CourseID int64 `gorm:"column:course_id"` //課表id
	UserID int64 `gorm:"column:user_id"` //用戶id
	Image string `gorm:"column:image"` //圖片
	CreateAt string `gorm:"column:create_at"` //創建時間
}

func (ReviewImage) TableName() string {
	return "review_images"
}

type ReviewStatistic struct {
	CourseID int64 `gorm:"column:course_id"` //課表id
	ScoreTotal int `gorm:"column:score_total"` //評分累積
	Amount int `gorm:"column:amount"` //評分筆數
	FiveTotal int `gorm:"column:five_total"` //五分總筆數
	FourTotal int `gorm:"column:four_total"` //四分總筆數
	ThreeTotal int `gorm:"column:three_total"` //三分總筆數
	TwoTotal int `gorm:"column:two_total"` //二分總筆數
	OneTotal int `gorm:"column:one_total"` //一分總筆數
	UpdateAt string `gorm:"column:update_at"` //更新時間
}

func (ReviewStatistic) TableName() string {
	return "review_statistics"
}

type CreateReviewParam struct {
	CourseID int64 //課表id
	UserID int64 //用戶id
	Score int //評分
	Body string //內容
	ImageNames []string //圖片名稱
}