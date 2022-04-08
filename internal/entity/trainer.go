package entity

type Trainer struct {
	UserID        int64  `gorm:"column:user_id"`        // 關聯的用戶id
	Name          string `gorm:"column:name"`           // 教練本名
	Nickname      string `gorm:"column:nickname"`       // 教練暱稱
	Skill         string `gorm:"column:skill"`          // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
	Avatar        string `gorm:"column:avatar"`         // 教練大頭照
	TrainerStatus int    `gorm:"column:trainer_status"` // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
	TrainerLevel  int    `gorm:"column:trainer_level"`  // 教練評鑑等級
	Email         string `gorm:"column:email"`          // 信箱
	Phone         string `gorm:"column:phone"`          // 電話
	Address       string `gorm:"column:address"`        // 住址
	Intro         string `gorm:"column:intro"`          // 個人介紹
	Experience    int    `gorm:"column:experience"`     // 年資
	Motto         string `gorm:"column:motto"`          // 座右銘
	FacebookURL   string `gorm:"column:facebook_url"`   // 臉書連結
	InstagramURL  string `gorm:"column:instagram_url"`  // ig連結
	YoutubeURL    string `gorm:"column:youtube_url"`    // youtube連結
	CreateAt      string `gorm:"column:create_at"`      // 創建日期
	UpdateAt      string `gorm:"column:update_at"`      // 修改日期
}

func (Trainer) TableName() string {
	return "trainers"
}
