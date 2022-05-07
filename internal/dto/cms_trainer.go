package dto

type CMSTrainerSummary struct {
	UserID        int64  `json:"user_id" json:"user_id" example:"1001"`                    // 用戶id
	Nickname      string `json:"nickname" json:"nickname" example:"Henry"`                 // 教練暱稱
	TrainerStatus int    `json:"trainer_status" json:"trainer_status" example:"1"`         // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
	Email         string `json:"email" json:"email" example:"henry@gmail.com"`             // 信箱
	CreateAt      string `json:"create_at" json:"create_at" example:"2021-06-01 12:00:00"` // 創建日期
	UpdateAt      string `json:"update_at" json:"update_at" example:"2021-06-01 12:00:00"` // 修改日期
}

func (CMSTrainerSummary) TableName() string {
	return "trainers"
}

type CMSTrainer struct {
	UserID             int64                `json:"user_id" gorm:"column:user_id" example:"10001"`                         // 用戶id
	Name               string               `json:"name" gorm:"column:name" example:"王小明"`                                 // 教練本名
	Nickname           string               `json:"nickname" gorm:"column:nickname" example:"戰車老師"`                        // 教練暱稱
	Skill              string               `json:"skill" gorm:"column:skill" example:"1,3,5"`                             // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
	Avatar             string               `json:"avatar" gorm:"column:avatar" example:"ld3ae0faf5we.png"`                // 教練大頭照
	TrainerStatus      int                  `json:"trainer_status" gorm:"column:trainer_status" example:"1"`               // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
	TrainerLevel       int                  `json:"trainer_level" gorm:"column:trainer_level" example:"1"`                 // 教練評鑑等級
	Email              string               `json:"email" gorm:"column:email" example:"henry@gmail.com"`                   // 信箱
	Phone              string               `json:"phone" gorm:"column:phone" example:"0955321456"`                        // 電話
	Address            string               `json:"address" gorm:"column:address" example:"台北市信義區信義路五段五號"`                 // 住址
	Intro              string               `json:"intro" gorm:"column:intro" example:"我叫戰車老師"`                            // 個人介紹
	Experience         int                  `json:"experience" gorm:"column:experience" example:"10"`                      // 年資
	Motto              string               `json:"motto" gorm:"column:motto" example:"我不會煎一塊牛排，因為我都煎兩塊"`                  // 座右銘
	FacebookURL        string               `json:"facebook_url" gorm:"column:facebook_url" example:"www.facebook.com"`    // 臉書連結
	InstagramURL       string               `json:"instagram_url" gorm:"column:instagram_url" example:"www.instagram.com"` // ig連結
	YoutubeURL         string               `json:"youtube_url" gorm:"column:youtube_url" example:"www.youtube.com"`       // youtube連結
	CreateAt           string               `json:"create_at" gorm:"column:create_at" example:"2021-06-01 12:00:00"`       // 創建日期
	UpdateAt           string               `json:"update_at" gorm:"column:update_at" example:"2021-06-01 12:00:00"`       // 修改日期
	BankAccount        *BankAccount         `json:"bank_account" gorm:"-"`                                                 // 銀行資訊
	Card               *Card                `json:"card" gorm:"-"`                                                        // 身分證
	TrainerAlbumPhotos []*TrainerAlbumPhoto `json:"trainer_album_photos" gorm:"-"`                                         // 教練相簿
	Certificates       []*Certificate       `json:"certificates" gorm:"-"`                                                 // 教練證照
}

type FinsCMSTrainersParam struct {
	UserID        *int64  // 用戶ID
	NickName      *string // 教練暱稱(1~40字元)
	Email         *string // 用戶Email
	TrainerStatus *int    // 教練狀態 教練狀態(1:正常/2:審核中/3:停權/4:未啟用)
}
