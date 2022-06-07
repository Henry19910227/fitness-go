package entity

type Diet struct {
	ID         int64  `gorm:"column:id"`          //id
	UserID     int64  `gorm:"column:user_id"`     //用戶id
	RdaID      int64  `gorm:"column:rda_id"`      //建議營養id
	ScheduleAt string `gorm:"column:schedule_at"` //排程時間
	CreateAt   string `gorm:"column:create_at"`   //創建時間
	UpdateAt   string `gorm:"column:update_at"`   //更新時間
}

func (Diet) TableName() string {
	return "diets"
}
