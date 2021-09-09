package model

type Trainer struct {
	UserID           int64   `gorm:"column:user_id"`         // 關聯的用戶id
	Name             string  `gorm:"column:name"`            // 教練本名
	Nickname         *string  `gorm:"column:nickname"`        // 教練暱稱
	Avatar           string  `gorm:"column:avatar"`          // 教練大頭照
	TrainerStatus    int     `gorm:"column:trainer_status"`  // 教練帳戶狀態 (1:正常/2:審核中/3:停權/4:未啟用)
	Email            string  `gorm:"column:email"`           // 信箱
	Phone            string  `gorm:"column:phone"`           // 電話
	Address          string  `gorm:"column:address"`         // 住址
	Intro            string  `gorm:"column:intro"`           // 個人介紹
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
	Name string `gorm:"column:name"`
	Address string `gorm:"column:address"`
	Phone string `gorm:"column:phone"`
	Email string `gorm:"column:email"`
}

type UpdateTrainerParam struct {
	Name             *string  `gorm:"column:name"`            // 教練本名
	Nickname         *string  `gorm:"column:nickname"`        // 教練暱稱
	Avatar           *string  `gorm:"column:avatar"`          // 教練大頭照
	TrainerStatus    *int     `gorm:"column:trainer_status"`  // 教練帳戶狀態 (1:正常/2:審核中/3:停權/4:未啟用)
	Email            *string  `gorm:"column:email"`           // 信箱
	Phone            *string  `gorm:"column:phone"`           // 電話
	Address          *string  `gorm:"column:address"`         // 住址
	Intro            *string  `gorm:"column:intro"`           // 個人介紹
	Experience       *int     // 年資
	Motto            *string  // 座右銘
	FacebookURL      *string  // 臉書連結
	InstagramURL     *string  // ig連結
	YoutubeURL       *string  // youtube連結
	UpdateAt         *string  `gorm:"column:update_at"`       // 修改日期
}