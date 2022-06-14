package dto

import "github.com/Henry19910227/fitness-go/internal/v1/model"

type Trainer struct {
	UserID             int64                `json:"user_id" gorm:"column:user_id" example:"1001"`                          // 用戶id
	Name               string               `json:"name" gorm:"column:name" example:"王小明"`                                 // 教練本名
	Nickname           string               `json:"nickname" gorm:"column:nickname" example:"Henry"`                       // 教練暱稱
	Skill              string               `json:"skill" gorm:"column:skill" example:"1,3,5"`                             // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
	Avatar             string               `json:"avatar" gorm:"column:avatar" example:"ld3ae0faf5we.png"`                // 教練大頭照
	TrainerStatus      int                  `json:"trainer_status" gorm:"column:trainer_status" example:"1"`               // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
	TrainerLevel       int                  `json:"trainer_level" gorm:"column:trainer_level" example:"1"`                 // 教練評鑑等級
	Email              string               `json:"email" gorm:"column:email" example:"henry@gmail.com"`                   // 信箱
	Phone              string               `json:"phone" gorm:"column:phone" example:"0978820789"`                        // 電話
	Address            string               `json:"address" gorm:"column:address" example:"台北市信義區信義路五段五號"`                 // 住址
	Intro              string               `json:"intro" gorm:"column:intro" example:"健身比賽冠軍"`                            // 個人介紹
	Experience         int                  `json:"experience" gorm:"column:experience" example:"1"`                       // 年資
	Motto              string               `json:"motto" gorm:"column:motto" example:"我的座右銘"`                             // 座右銘
	FacebookURL        string               `json:"facebook_url" gorm:"column:facebook_url" example:"www.facebook.com"`    // 臉書連結
	InstagramURL       string               `json:"instagram_url" gorm:"column:instagram_url" example:"www.instagram.com"` // ig連結
	YoutubeURL         string               `json:"youtube_url" gorm:"column:youtube_url" example:"www.youtube.com"`       // youtube連結
	Favorite           int                  `json:"favorite" gorm:"-" example:"1"`                                         // 是否收藏(0:否/1:是)
	TrainerStatistic   *TrainerStatistic    `json:"trainer_statistic" gorm:"-"`                                            // 教練統計
	TrainerAlbumPhotos []*TrainerAlbumPhoto `json:"trainer_album_photos" gorm:"-"`                                         // 教練相簿
	Certificates       []*Certificate       `json:"certificates" gorm:"-"`                                                 // 教練證照
}

type TrainerSummary struct {
	UserID   int64  `json:"user_id" example:"10001"`            // 關聯的用戶id
	Nickname string `json:"nickname" example:"Henry教練"`         // 教練暱稱
	Avatar   string `json:"avatar"  example:"d2w3e15d3awe.jpg"` // 教練大頭照
	Skill    string `json:"skill"  example:"d2w3e15d3awe.jpg"`  // 專長
}

type GetTrainerSummariesParam struct {
	OrderType *string // 排序類型(latest:最新/popular:熱門)-單選
}

type CreateTrainerParam struct {
	Name               string  // 教練本名
	Nickname           string  // 教練暱稱
	Skill              []int   // 專長
	Email              string  // 信箱
	Phone              string  // 電話
	Address            string  // 住址
	Intro              string  // 個人介紹
	Experience         int     // 年資
	Motto              *string // 座右銘
	FacebookURL        *string // 臉書連結
	InstagramURL       *string // ig連結
	YoutubeURL         *string // youtube連結
	Avatar             *File
	CardFrontImage     *File
	CardBackImage      *File
	TrainerAlbumPhotos []*File
	CertificateImages  []*File
	CertificateNames   []string
	AccountName        string // 帳戶名稱
	AccountImage       *File  // 帳戶照片
	BankCode           string // 銀行代號
	Account            string // 帳戶
	Branch             string // 銀行分行
}

type UpdateTrainerParam struct {
	Nickname             *string // 暱稱 (1~20字元)
	Skill                []int   // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
	TrainerStatus        *int    // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
	TrainerLevel         *int    // 教練評鑑等級
	Intro                *string // 教練介紹 (1~400字元)
	Experience           *int    // 年資
	Motto                *string // 座右銘 (1~100字元)
	FacebookURL          *string // 臉書連結
	InstagramURL         *string // instagram連結
	YoutubeURL           *string // youtube連結
	Avatar               *File
	DeleteAlbumPhotosIDs []int64  // 待刪除的相簿照片id
	CreateAlbumPhotos    []*File  // 待新增的相簿照片
	DeleteCerIDs         []int64  // 待刪除的證照照片id
	UpdateCerIDs         []int64  // 待更新的證照照片id
	UpdateCerImages      []*File  // 待更新的證照照片
	UpdateCerNames       []string // 待更新的證照名稱
	CreateCerNames       []string // 待新增的證照名稱
	CreateCerImages      []*File  // 待更新的證照照片
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

func NewTrainer(data *model.Trainer) Trainer {
	trainer := Trainer{
		UserID:        data.UserID,
		Name:          data.Name,
		Nickname:      data.Nickname,
		Skill:         data.Skill,
		Avatar:        data.Avatar,
		TrainerStatus: data.TrainerStatus,
		TrainerLevel:  data.TrainerLevel,
		Email:         data.Email,
		Phone:         data.Phone,
		Address:       data.Address,
		Intro:         data.Intro,
		Experience:    data.Experience,
		Motto:         data.Motto,
		FacebookURL:   data.FacebookURL,
		InstagramURL:  data.InstagramURL,
		YoutubeURL:    data.YoutubeURL,
	}
	trainer.TrainerStatistic = &TrainerStatistic{}
	if data.TrainerStatistic != nil {
		trainer.TrainerStatistic.ReviewScore = data.TrainerStatistic.ReviewScore
		trainer.TrainerStatistic.StudentCount = data.TrainerStatistic.StudentCount
		trainer.TrainerStatistic.CourseCount = data.TrainerStatistic.CourseCount
	}
	return trainer
}
