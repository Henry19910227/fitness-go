package model

type Trainer struct {
	UserID           int64   `gorm:"column:user_id"`         // 關聯的用戶id
	Name             string  `gorm:"column:name"`            // 教練本名
	Nickname         string  `gorm:"column:nickname"`        // 教練暱稱
	Avatar           string  `gorm:"column:avatar"`          // 教練大頭照
	TrainerStatus    int     `gorm:"column:trainer_status"`  // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
	Email            string  `gorm:"column:email"`           // 信箱
	Phone            string  `gorm:"column:phone"`           // 電話
	Address          string  `gorm:"column:address"`         // 住址
	Intro            string  `gorm:"column:intro"`           // 個人介紹
	Experience       int     `gorm:"column:experience"`      // 年資
	Motto            string  `gorm:"column:motto"`           // 座右銘
	FacebookURL      string  `gorm:"column:facebook_url"`        // 臉書連結
	InstagramURL     string  `gorm:"column:instagram_url"`       // ig連結
	YoutubeURL       string  `gorm:"column:youtube_url"`         // youtube連結
	CreateAt         string  `gorm:"column:create_at"`       // 創建日期
	UpdateAt         string  `gorm:"column:update_at"`       // 修改日期
}

func (Trainer) TableName() string {
	return "trainers"
}

type TrainerSummaryEntity struct {
	UserID           int64   `gorm:"column:user_id"`         // 關聯的用戶id
	Nickname         string  `gorm:"column:nickname"`        // 教練暱稱
	Avatar           string  `gorm:"column:avatar"`          // 教練大頭照
}

type CreateTrainerParam struct {
	Name               string     // 教練本名
	Nickname           string     // 教練暱稱
	Email              string     // 信箱
	Phone              string     // 電話
	Address            string     // 住址
	Intro              string     // 個人介紹
	Experience         int        // 年資
	Motto              *string    // 座右銘
	CardFrontImage     string     // 身分證正面
	CardBackImage      string     // 身分證反面
	FacebookURL        *string    // 臉書連結
	InstagramURL       *string    // ig連結
	YoutubeURL         *string    // youtube連結
	Avatar             string     // 大頭照
	TrainerAlbumPhotos []string   // 教練相簿照片
	CertificateImages  []string   // 證照照片
	CertificateNames   []string   // 證照名稱
	AccountName        string     // 帳戶名稱
	AccountImage       string     // 帳戶照片
	BankCode           string     // 銀行代號
	Branch             string     // 分行
	Account            string     // 帳戶
}

type UpdateTrainerParam struct {
	Name             *string  `gorm:"column:name"`            // 教練本名
	Nickname         *string  `gorm:"column:nickname"`        // 教練暱稱
	Avatar           *string  `gorm:"column:avatar"`          // 教練大
	TrainerStatus    *int     `gorm:"column:trainer_status"`  // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
	Email            *string  `gorm:"column:email"`           // 信箱
	Phone            *string  `gorm:"column:phone"`           // 電話
	Address          *string  `gorm:"column:address"`         // 住址
	Intro            *string  `gorm:"column:intro"`           // 個人介紹
	Experience       *int     // 年資
	Motto            *string  // 座右銘
	CardID           *string  // 身分證字號
	CardFrontImage   *string  // 身分證正面
	CardBackImage   *string   // 身分證反面
	FacebookURL      *string  // 臉書連結
	InstagramURL     *string  // ig連結
	YoutubeURL       *string  // youtube連結
	UpdateAt         *string  `gorm:"column:update_at"`       // 修改日期
}