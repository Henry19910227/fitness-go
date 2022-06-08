package model

type Diet struct {
	ID         int64   `json:"id" gorm:"column:id"`                           //id
	UserID     int64   `json:"user_id" gorm:"column:user_id"`                 //用戶id
	RdaID      int64   `json:"rda_id" gorm:"column:rda_id"`                   //建議營養id
	ScheduleAt string  `json:"schedule_at" gorm:"column:schedule_at"`         //排程時間
	CreateAt   string  `json:"create_at" gorm:"column:create_at"`             //創建時間
	UpdateAt   string  `json:"update_at" gorm:"column:update_at"`             //更新時間
	RDA        *RDA    `json:"rda" gorm:"foreignkey:id;references:rda_id"`    //營養建議
	Meals      []*Meal `json:"meals" gorm:"foreignkey:diet_id;references:id"` //餐食
}

type DietItem struct {
	ID         int64  `gorm:"column:id"`          //id
	UserID     int64  `gorm:"column:user_id"`     //用戶id
	RdaID      int64  `gorm:"column:rda_id"`      //建議營養id
	ScheduleAt string `gorm:"column:schedule_at"` //排程時間
	CreateAt   string `gorm:"column:create_at"`   //創建時間
	UpdateAt   string `gorm:"column:update_at"`   //更新時間
}

func (DietItem) TableName() string {
	return "diets"
}

type FindDietParam struct {
	ID         *int64
	UserID     *int64
	ScheduleAt *string
}

type FindDietsParam struct {
	PreloadParam
	UserID          *int64
	AfterScheduleAt *string
}

type SaveDietsParam struct {
	Diets []*DietItem
}
