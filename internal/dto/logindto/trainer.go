package logindto

type Trainer struct {
	UserID           int64   `json:"-" gorm:"column:user_id"`                                           // 用戶id
	Name             string  `json:"name" gorm:"column:name" example:"王小明"`                           // 教練本名
	Nickname         string  `json:"nickname" gorm:"column:nickname" example:"Henry"`                   // 教練暱稱
	Avatar           string  `json:"avatar" gorm:"column:avatar" example:"ld3ae0faf5we.png"`            // 教練大頭照
	TrainerStatus    int     `json:"trainer_status" gorm:"column:trainer_status" example:"1"`           // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
	Email            string  `json:"email" gorm:"column:email" example:"henry@gmail.com"`               // 信箱
	Phone            string  `json:"phone" gorm:"column:phone" example:"0978820789"`                    // 電話
	Address          string  `json:"address" gorm:"column:address" example:"台北市信義區信義路五段五號"`     // 住址
	Intro            string  `json:"intro" gorm:"column:intro" example:"健身比賽冠軍"`                    // 個人介紹
}
