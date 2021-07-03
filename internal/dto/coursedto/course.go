package coursedto

import "mime/multipart"

type CreateResult struct {
	ID int64 `json:"course_id" example:"1"` //課表 id
}

type CourseCover struct {
	Cover string `json:"cover" example:"dkf2se51fsdds.png"` // 課表封面照片
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

type UploadCourseCoverParam struct {
	CoverNamed string
	File       multipart.File
}

type Course struct {
	ID       int64  `json:"id" gorm:"column:id" example:"2"`                                  // 課表 id
	UserID   int64  `json:"user_id" gorm:"column:user_id" example:"10001"`                    // 用戶 id
	CourseStatus int `json:"course_status" gorm:"column:course_status" example:"1"`           // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category int `json:"category" gorm:"column:category" example:"3"`                         // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int `json:"schedule_type" gorm:"column:schedule_type" example:"2"`           // 排課類別(1:單一訓練/2:多項計畫)
	SaleType int `json:"sale_type" gorm:"column:sale_type" example:"2"`                       // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Price int64 `json:"price" gorm:"column:price" example:"330"`                              // 售價
	Name string `json:"name" gorm:"column:name" example:"Henry課表"`                           // 課表名稱
	Cover string `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`                // 課表封面
	Intro string `json:"intro" gorm:"column:intro" example:"佛系課表"`                          // 課表介紹
	Food string `json:"food" gorm:"column:food" example:"佛系飲食"`                             // 飲食建議
	Level int `json:"level" gorm:"column:level" example:"3"`                                  // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit string `json:"suit" gorm:"column:suit" example:"2,5,7"`                              // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment string `json:"equipment" gorm:"column:equipment" example:"2,3,6"`               // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place string `json:"place" gorm:"column:place" example:"1,2"`                             // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget string `json:"train_target" gorm:"column:train_target" example:"2,3,4"`       // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget string `json:"body_target" gorm:"column:body_target" example:"4,5"`            // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice string `json:"notice" gorm:"column:notice" example:"小心不要受傷"`                   // 注意事項
	PlanCount int `json:"plan_count" gorm:"column:plan_count" example:"2"`                    // 計畫總數
	WorkoutCount int `json:"workout_count" gorm:"column:workout_count" example:"10"`          // 訓練總數
	CreateAt string `json:"create_at" gorm:"column:create_at" example:"2021-05-28 11:00:00"`  // 創建時間
	UpdateAt string `json:"update_at" gorm:"column:update_at" example:"2021-05-29 11:00:00"`  // 更新時間
}

type CourseID struct {
	ID       int64  `json:"id" gorm:"column:id" example:"2"`                                  // 課表 id
}