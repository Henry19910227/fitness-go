package coursedto

type CreateResult struct {
	ID int64 `json:"course_id" example:"1"` //課表 id
}

type CreateCourseParam struct {
	Name string
	Level int
	Category int
	CategoryOther string
	ScheduleType int
}


type UpdateCourseParam struct {
	Category *int `gorm:"column:category"`                    // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType *int `gorm:"column:schedule_type"`           // 排課類別(1:單一訓練/2:多項計畫)
	SaleType *int `gorm:"column:sale_type"`                   // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Price *int64 `gorm:"column:price"`                        // 售價
	Name *string `gorm:"column:name"`                         // 課表名稱
	Intro *string `gorm:"column:intro"`                       // 課表介紹
	Food *string `gorm:"column:food"`                         // 飲食建議
	Level *int `gorm:"column:level"`                          // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit *string `gorm:"column:suit"`                         // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment *string `gorm:"column:equipment"`               // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place *string `gorm:"column:place"`                       // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget *string `gorm:"column:train_target"`          // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget *string `gorm:"column:body_target"`            // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice *string `gorm:"column:notice"`                     // 注意事項
	UpdateAt *string `gorm:"column:update_at"`                // 更新時間
}