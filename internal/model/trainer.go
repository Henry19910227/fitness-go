package model

type Trainer struct {
	UserID           int64   `gorm:"column:user_id"`         // 關聯的用戶id
	Name             string  `gorm:"column:name"`            // 教練本名
	Nickname         string  `gorm:"column:nickname"`        // 教練暱稱
	Skill            string  `gorm:"column:skill"`           // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
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
	Skill            string  `gorm:"column:skill"`           // 專長
}

type CreateTrainerParam struct {
	Name               string     // 教練本名
	Nickname           string     // 教練暱稱
	Skill              string     // 專長
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
	Nickname         *string  // 教練暱稱
	Skill            *string  // 專長
	TrainerStatus    *int     // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
	Intro            *string  // 個人介紹
	Experience       *int     // 年資
	Motto            *string  // 座右銘
	FacebookURL      *string  // 臉書連結
	InstagramURL     *string  // ig連結
	YoutubeURL       *string  // youtube連結
	Avatar           *string  // 教練大頭照
	DeleteAlbumPhotosIDs []int64 // 待刪除的相簿照片id
	CreateAlbumPhotos []string // 待新增的相簿照片
	DeleteCerIDs     []int64  // 待刪除的證照照片id
	UpdateCerIDs     []int64  // 待更新的證照照片id
	UpdateCerImages   []string // 待更新的證照照片
	UpdateCerNames   []string // 待更新的證照名稱
	CreateCerImages []string  // 待新增的證照照片
	CreateCerNames   []string // 待新增的證照名稱
	UpdateAt         *string  // 修改日期
}