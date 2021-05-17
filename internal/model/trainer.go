package model

type Trainer struct {
	ID               int64   `gorm:"column:id"`              // 教練id
	Name             string  `gorm:"column:name"`            // 教練本名
	Nickname         string  `gorm:"column:nickname"`        // 教練暱稱
	TrainerStatus    int     `gorm:"column:trainer_status"`  // 教練帳戶狀態
	Birthday         string  `gorm:"column:birthday"`        // 生日
	Email            string  `gorm:"column:email"`           // 信箱
	Phone            string  `gorm:"column:phone"`           // 電話
	Address          string  `gorm:"column:address"`         // 住址
	Intro            string  `gorm:"column:intro"`           // 個人介紹
	CreateAt         string  `gorm:"column:create_at"`       // 創建日期
	UpdateAt         string  `gorm:"column:update_at"`       // 修改日期
	UserID           int64   `gorm:"column:user_id"`              // 關聯的用戶id
}

func (Trainer) TableName() string {
	return "trainers"
}

type CreateTrainerParam struct {
	Name string `gorm:"column:name"`
	Nickname string `gorm:"column:nickname"`
	Phone string `gorm:"column:phone"`
	Email string `gorm:"column:email"`
}