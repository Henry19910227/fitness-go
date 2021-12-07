package model

type Course struct {
	ID       int64  `gorm:"column:id"`                       // 課表 id
	UserID   int64  `gorm:"column:user_id"`                  // 用戶 id
	SaleID *int64 `gorm:"column:sale_id"`                    // 銷售 id
	CourseStatus int `gorm:"column:course_status"`           // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category int `gorm:"column:category"`                    // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int `gorm:"column:schedule_type"`           // 排課類別(1:單一訓練/2:多項計畫)
	Name string `gorm:"column:name"`                         // 課表名稱
	Cover string `gorm:"column:cover"`                       // 課表封面
	Intro string `gorm:"column:intro"`                       // 課表介紹
	Food string `gorm:"column:food"`                         // 飲食建議
	Level int `gorm:"column:level"`                          // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit string `gorm:"column:suit"`                         // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment string `gorm:"column:equipment"`               // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place string `gorm:"column:place"`                       // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget string `gorm:"column:train_target"`          // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget string `gorm:"column:body_target"`            // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice string `gorm:"column:notice"`                     // 注意事項
	PlanCount int `gorm:"column:plan_count"`                 // 計畫總數
	WorkoutCount int `gorm:"column:workout_count"`           // 訓練總數
	CreateAt string `gorm:"column:create_at"`                // 創建時間
	UpdateAt string `gorm:"column:update_at"`                // 更新時間
}

func (Course) TableName() string {
	return "courses"
}

type CourseSummaryEntity struct {
	ID       int64  `gorm:"column:id"`                       // 課表 id
	Trainer  TrainerSummaryEntity                            // 教練簡介
	Sale     SaleItemEntity                                  // 銷售項目
	CourseStatus int `gorm:"column:course_status"`           // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category int `gorm:"column:category"`                    // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int `gorm:"column:schedule_type"`           // 排課類別(1:單一訓練/2:多項計畫)
	Name string `gorm:"column:name"`                         // 課表名稱
	Cover string `gorm:"column:cover"`                       // 課表封面
	Level int `gorm:"column:level"`                          // 強度(1:初級/2:中級/3:中高級/4:高級)
	PlanCount int `gorm:"column:plan_count"`                 // 計畫總數
	WorkoutCount int `gorm:"column:workout_count"`           // 訓練總數
}

type CourseDetailEntity struct {
	ID       int64  `gorm:"column:id"`                       // 課表 id
	Trainer  TrainerSummaryEntity                            // 教練簡介
	Sale     SaleItemEntity                                  // 銷售項目
	CourseStatus int `gorm:"column:course_status"`           // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category int `gorm:"column:category"`                    // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int `gorm:"column:schedule_type"`           // 排課類別(1:單一訓練/2:多項計畫)
	Name string `gorm:"column:name"`                         // 課表名稱
	Cover string `gorm:"column:cover"`                       // 課表封面
	Intro string `gorm:"column:intro"`                       // 課表介紹
	Food string `gorm:"column:food"`                         // 飲食建議
	Level int `gorm:"column:level"`                          // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit string `gorm:"column:suit"`                         // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment string `gorm:"column:equipment"`               // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place string `gorm:"column:place"`                       // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget string `gorm:"column:train_target"`          // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget string `gorm:"column:body_target"`            // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice string `gorm:"column:notice"`                     // 注意事項
	PlanCount int `gorm:"column:plan_count"`                 // 計畫總數
	WorkoutCount int `gorm:"column:workout_count"`           // 訓練總數
	CreateAt string `gorm:"column:create_at"`                // 創建時間
	UpdateAt string `gorm:"column:update_at"`                // 更新時間
}

type CourseProduct struct {
	ID       int64  `gorm:"column:id"`                       // 課表 id
	UserID   int64  `gorm:"column:user_id"`                  // 用戶 id
	SaleID *int64 `gorm:"column:sale_id"`                    // 銷售 id
	CourseStatus int `gorm:"column:course_status"`           // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category int `gorm:"column:category"`                    // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int `gorm:"column:schedule_type"`           // 排課類別(1:單一訓練/2:多項計畫)
	Name string `gorm:"column:name"`                         // 課表名稱
	Cover string `gorm:"column:cover"`                       // 課表封面
	Intro string `gorm:"column:intro"`                       // 課表介紹
	Food string `gorm:"column:food"`                         // 飲食建議
	Level int `gorm:"column:level"`                          // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit string `gorm:"column:suit"`                         // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment string `gorm:"column:equipment"`               // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place string `gorm:"column:place"`                       // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget string `gorm:"column:train_target"`          // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget string `gorm:"column:body_target"`            // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice string `gorm:"column:notice"`                     // 注意事項
	PlanCount int `gorm:"column:plan_count"`                 // 計畫總數
	WorkoutCount int `gorm:"column:workout_count"`           // 訓練總數
	CreateAt string `gorm:"column:create_at"`                // 創建時間
	UpdateAt string `gorm:"column:update_at"`                // 更新時間
	Trainer  *TrainerSummaryEntity `gorm:"foreignkey:user_id;references:user_id"` // 教練簡介
	Sale     *SaleItem             `gorm:"foreignkey:id;references:sale_id"`             // 銷售項目
	Review   ReviewStatistic       `gorm:"foreignkey:course_id;references:id"`           // 評分統計
}

func (CourseProduct) TableName() string {
	return "courses"
}

type CourseProductSummary struct {
	ID       int64  `gorm:"column:id"`                       // 課表 id
	Trainer  TrainerSummaryEntity                            // 教練簡介
	Sale     SaleItemEntity                                  // 銷售項目
	CourseStatus int `gorm:"column:course_status"`           // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category int `gorm:"column:category"`                    // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int `gorm:"column:schedule_type"`           // 排課類別(1:單一訓練/2:多項計畫)
	Name string `gorm:"column:name"`                         // 課表名稱
	Cover string `gorm:"column:cover"`                       // 課表封面
	Level int `gorm:"column:level"`                          // 強度(1:初級/2:中級/3:中高級/4:高級)
	PlanCount int `gorm:"column:plan_count"`                 // 計畫總數
	WorkoutCount int `gorm:"column:workout_count"`           // 訓練總數
	ReviewStatistic ReviewStatisticSummary                   // 統計資料
}

type FindCourseProductSummariesParam struct {
	Name *string //課表名稱
	Score *int // 評價(1~5分)-單選
	Level []int // 強度(1:初級/2:中級/3:中高級/4:高級)-複選
	Category []int // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)-複選
	Suit []int  // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)-複選
	Equipment []int  // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)-複選
	Place []int // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)-複選
	TrainTarget []int // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)-複選
	BodyTarget []int // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)-複選
	SaleType []int // 銷售類型(1:免費課表/2:付費課表/3:訂閱課表)-複選
	TrainerSex []string // 教練性別(m:男性/f:女性)-複選
	TrainerSkill []int // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
}

type FindCourseProductCountParam struct {
	Name *string //課表名稱
	Score *int // 評價(1~5分)-單選
	Level []int // 強度(1:初級/2:中級/3:中高級/4:高級)-複選
	Category []int // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)-複選
	Suit []int  // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)-複選
	Equipment []int  // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)-複選
	Place []int // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)-複選
	TrainTarget []int // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)-複選
	BodyTarget []int // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)-複選
	SaleType []int // 銷售類型(1:免費課表/2:付費課表/3:訂閱課表)-複選
	TrainerSex []string // 教練性別(m:男性/f:女性)-複選
	TrainerSkill []int // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
}

type FindCourseSummariesParam struct {
	UID *int64
	Status *int
}

type CreateCourseParam struct {
	Name string `gorm:"column:name"`
	Level int `gorm:"column:level"`
	Category int `gorm:"column:category"`
}

type UpdateCourseParam struct {
	CourseStatus *int    `gorm:"column:course_status"` // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category     *int    `gorm:"column:category"`      // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType *int    `gorm:"column:schedule_type"` // 排課類別(1:單一訓練/2:多項計畫)
	SaleID       *int    `gorm:"column:sale_id"`       // 銷售id
	Name         *string `gorm:"column:name"`          // 課表名稱
	Cover        *string `gorm:"column:cover"`         // 課表封面
	Intro        *string `gorm:"column:intro"`         // 課表介紹
	Food         *string `gorm:"column:food"`          // 飲食建議
	Level        *int    `gorm:"column:level"`         // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit         *string `gorm:"column:suit"`          // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment    *string `gorm:"column:equipment"`     // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place        *string `gorm:"column:place"`         // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget  *string `gorm:"column:train_target"`  // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget   *string `gorm:"column:body_target"`   // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice       *string `gorm:"column:notice"`        // 注意事項
	UpdateAt     *string `gorm:"column:update_at"`     // 更新時間
}
