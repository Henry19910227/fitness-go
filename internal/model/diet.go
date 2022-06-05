package model

type Diet struct {
	ID         int64  `json:"id" gorm:"column:id"`                              //id
	UserID     int64  `json:"user_id" gorm:"column:user_id"`                    //用戶id
	RdaID      int64  `json:"rda_id" gorm:"column:rda_id"`                      //建議營養id
	ScheduleAt string `json:"schedule_at" gorm:"column:schedule_at"`            //排程時間
	CreateAt   string `json:"create_at" gorm:"column:create_at"`                //創建時間
	UpdateAt   string `json:"update_at" gorm:"column:update_at"`                //更新時間
	RDA        *RDA   `json:"rda" gorm:"foreignkey:id;references:rda_id"`       //營養建議
	Meals 	   []*Meal	`json:"meals" gorm:"foreignkey:diet_id;references:id"`  //餐食
}

type FindDietParam struct {
	ID         *int64
	UserID     *int64
	ScheduleAt *string
}
