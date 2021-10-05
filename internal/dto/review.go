package dto

type Review struct {
	CourseID int64 `json:"course_id" gorm:"column:course_id" example:"1000"` //課表id
	UserID int64 `json:"user_id" gorm:"column:user_id" example:"1000"` //用戶id
	Score int `json:"score" gorm:"column:score" example:"1000"` //評分
	Body string `json:"body" gorm:"column:body" example:"這個課表很棒!"` //評論內容
	Images []*ReviewImage `json:"images" gorm:"-"` //評論照片
	CreateAt string `json:"create_at" gorm:"column:create_at" example:"1000"` //創建時間
}

type ReviewImage struct {
	ID int64 `json:"id" gorm:"column:id" example:"1"` //圖片id
	Image string `json:"image" gorm:"column:image" example:"ds241w564d5e2.png"` //圖片
}

type ReviewStatistic struct {
	ScoreTotal int `json:"score_total" gorm:"column:score_total" example:"1000"` //評分累積
	Amount int `json:"amount" gorm:"column:amount" example:"450"` //評分筆數
	FiveTotal int `json:"five_total" gorm:"column:five_total" example:"100"` //五分總筆數
	FourTotal int `json:"four_total" gorm:"column:four_total" example:"100"` //四分總筆數
	ThreeTotal int `json:"three_total" gorm:"column:three_total" example:"150"` //三分總筆數
	TwoTotal int `json:"two_total" gorm:"column:two_total" example:"100"` //二分總筆數
	OneTotal int `json:"one_total" gorm:"column:one_total" example:"50"` //一分總筆數
	UpdateAt string `json:"update_at" gorm:"column:update_at" example:"100"` //更新時間
}

type CreateReviewParam struct {
	CourseID int64 //課表id
	UserID int64 //用戶id
	Score int //評分
	Body string //內容
	Images []*File //圖片
}
