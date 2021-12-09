package dto

import "github.com/Henry19910227/fitness-go/internal/global"

type Review struct {
	ID       int64 `json:"id" example:"1"` //評論id
	CourseID int64 `json:"course_id" example:"2"` //課表id
	User *UserSummary `json:"user"` //用戶id
	Score int `json:"score" example:"1000"` //評分
	Body string `json:"body" example:"這個課表很棒!"` //評論內容
	Images []*ReviewImage `json:"images"` //評論照片
	CreateAt string `json:"create_at" example:"2021-06-01 12:00:00"` //創建時間
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

type ReviewStatisticSummary struct {
	ScoreTotal int `json:"score_total" gorm:"column:score_total" example:"1000"` //評分累積
	Amount int `json:"amount" gorm:"column:amount" example:"450"` //評分筆數
}

type CreateReviewParam struct {
	CourseID int64 //課表id
	UserID int64 //用戶id
	Score int //評分
	Body string //內容
	Images []*File //圖片
}

type GetReviewsParam struct {
	CourseID int64
	FilterType global.ReviewFilterType //篩選類型(1:全部/2:有照片)
}
