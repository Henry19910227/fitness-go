package dto

type Trainer struct {
	UserID           int64   `json:"user_id" gorm:"column:user_id" example:"1001"`                      // 用戶id
	Name             string  `json:"name" gorm:"column:name" example:"王小明"`                           // 教練本名
	Nickname         string  `json:"nickname" gorm:"column:nickname" example:"Henry"`                   // 教練暱稱
	Avatar           string  `json:"avatar" gorm:"column:avatar" example:"ld3ae0faf5we.png"`            // 教練大頭照
	TrainerStatus    int     `json:"trainer_status" gorm:"column:trainer_status" example:"1"`           // 教練帳戶狀態 (1:正常/2:審核中/3:停權/4:未啟用)
	Email            string  `json:"email" gorm:"column:email" example:"henry@gmail.com"`               // 信箱
	Phone            string  `json:"phone" gorm:"column:phone" example:"0978820789"`                    // 電話
	Address          string  `json:"address" gorm:"column:address" example:"台北市信義區信義路五段五號"`     // 住址
	Intro            string  `json:"intro" gorm:"column:intro" example:"健身比賽冠軍"`                    // 個人介紹
	Experience       int     `json:"experience" gorm:"column:experience" example:"1"`                   // 年資
	Motto            string  `json:"motto" gorm:"column:motto" example:"我的座右銘"`                      // 座右銘
	CardID           string  `json:"card_id" gorm:"column:card_id" example:"A123456789"`                 // 身分證字號
	CardFrontImage   string  `json:"card_front_image" gorm:"column:card_front_image" example:"ld3ae0faf5we.png"`  // 身分證正面
	CardBackImage    string  `json:"card_back_image" gorm:"column:card_back_image" example:"ld3ae0faf5we.png"`  // 身分證反面
	FacebookURL      string  `json:"facebook_url" gorm:"column:facebook_url" example:"www.facebook.com"`  // 臉書連結
	InstagramURL     string  `json:"instagram_url" gorm:"column:instagram_url" example:"www.instagram.com"`  // ig連結
	YoutubeURL       string  `json:"youtube_url" gorm:"column:youtube_url" example:"www.youtube.com"`  // youtube連結
	CreateAt         string  `json:"create_at" gorm:"column:create_at" example:"2021-05-10 10:00:00"`   // 創建日期
	UpdateAt         string  `json:"update_at" gorm:"column:update_at" example:"2021-05-10 10:00:00"`   // 修改日期
}

type TrainerSummary struct {
	UserID           int64   `json:"user_id" example:"10001"`                 // 關聯的用戶id
	Nickname         string  `json:"nickname" example:"Henry教練"`             // 教練暱稱
	Avatar           string  `json:"avatar"  example:"d2w3e15d3awe.jpg"`      // 教練大頭照
}

type CreateTrainerParam struct {
	Name string
	Phone string
	Email string
	Address string
}

type UpdateTrainerParam struct {
	Name             *string  // 教練本名
	Nickname         *string  // 教練暱稱
	Email            *string  // 信箱
	Phone            *string  // 電話
	Address          *string  // 住址
	Intro            *string  // 個人介紹
	Experience       *int     // 年資
	Motto            *string  // 座右銘
	FacebookURL      *string  // 臉書連結
	InstagramURL     *string  // ig連結
	YoutubeURL       *string  // youtube連結
}

type TrainerAvatar struct {
	Avatar string `json:"avatar" example:"dkf2se51fsdds.png"` // 教練大頭照
}

type TrainerCardFront struct {
	Image string `json:"card_front_image" example:"dkf2se51fsdds.png"` // 身分證正面
}

type TrainerCardBack struct {
	Image string `json:"card_back_image" example:"dkf2se51fsdds.png"` // 身分證背面
}