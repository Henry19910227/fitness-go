package model

import "github.com/Henry19910227/fitness-go/internal/pkg/global"

type Review struct {
	ID       int64              `gorm:"column:id"`                          //評論id
	CourseID int64              `gorm:"column:course_id"`                   //課表id
	UserID   int64              `gorm:"column:user_id"`                     //用戶id
	User     *UserSummary       `gorm:"foreignkey:id;references:user_id"`   //用戶
	Score    int                `gorm:"column:score"`                       //評分
	Body     string             `gorm:"column:body"`                        //內容
	Images   []*ReviewImageItem `gorm:"foreignKey:review_id;references:id"` //圖片
	CreateAt string             `gorm:"column:create_at"`                   //創建時間
}

func (Review) TableName() string {
	return "reviews"
}

type ReviewImageItem struct {
	ID       int64  `gorm:"column:id"`        //圖片id
	ReviewID int64  `gorm:"column:review_id"` //評論id
	Image    string `gorm:"column:image"`     //圖片
	CreateAt string `gorm:"column:create_at"` //創建時間
}

func (ReviewImageItem) TableName() string {
	return "review_images"
}

type CreateReviewParam struct {
	CourseID   int64    //課表id
	UserID     int64    //用戶id
	Score      int      //評分
	Body       string   //內容
	ImageNames []string //圖片名稱
}

type FindReviewsParam struct {
	CourseID   int64
	FilterType global.ReviewFilterType //篩選類型(1:全部/2:有照片)
}
