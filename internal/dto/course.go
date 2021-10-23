package dto

import (
	"mime/multipart"
)

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
	SaleID *int `gorm:"column:sale_id"`                       // 銷售id
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

type CourseSummary struct {
	ID       int64           `json:"id" example:"2"`                                         // 課表 id
	Trainer  *TrainerSummary `json:"trainer"`                                                // 教練簡介
	Sale     *SaleItem       `json:"sale"`                                                   // 銷售資料
	CourseStatus int         `json:"course_status" example:"1"`                              // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category int             `json:"category" gorm:"column:category" example:"3"`            // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int         `json:"schedule_type" gorm:"column:schedule_type" example:"2"`  // 排課類別(1:單一訓練/2:多項計畫)
	Name string              `json:"name" example:"Henry課表"`                                 // 課表名稱
	Cover string             `json:"cover" example:"d2w3e15d3awe.jpg"`                       // 課表封面
	Level int                `json:"level" example:"3"`                                      // 強度(1:初級/2:中級/3:中高級/4:高級)
	PlanCount int            `json:"plan_count" gorm:"column:plan_count" example:"2"`        // 計畫總數
	WorkoutCount int         `json:"workout_count" gorm:"column:workout_count" example:"10"` // 訓練總數
}

type Course struct {
	ID       int64           `json:"id" gorm:"column:id" example:"2"`                        // 課表 id
	Trainer  *TrainerSummary `json:"trainer"`                                                // 教練簡介
	Sale     *SaleItem       `json:"sale"`                                                   // 銷售資料
	Restricted int           `json:"restricted" example:"0"`                                 // 是否是限制訪問狀態(0:否/1:是)
	CourseStatus int         `json:"course_status" gorm:"column:course_status" example:"1"`  // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category int             `json:"category" gorm:"column:category" example:"3"`            // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int         `json:"schedule_type" gorm:"column:schedule_type" example:"2"`  // 排課類別(1:單一訓練/2:多項計畫)
	Name string              `json:"name" gorm:"column:name" example:"Henry課表"`              // 課表名稱
	Cover string             `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`   // 課表封面
	Intro string             `json:"intro" gorm:"column:intro" example:"佛系課表"`               // 課表介紹
	Food string              `json:"food" gorm:"column:food" example:"佛系飲食"`                 // 飲食建議
	Level int                `json:"level" gorm:"column:level" example:"3"`                  // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit string              `json:"suit" gorm:"column:suit" example:"2,5,7"`                // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment string `json:"equipment" gorm:"column:equipment" example:"2,3,6"`              // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place string `json:"place" gorm:"column:place" example:"1,2"`                            // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget string `json:"train_target" gorm:"column:train_target" example:"2,3,4"`      // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget string `json:"body_target" gorm:"column:body_target" example:"4,5"`           // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice string `json:"notice" gorm:"column:notice" example:"小心不要受傷"`                      // 注意事項
	PlanCount int `json:"plan_count" gorm:"column:plan_count" example:"2"`                   // 計畫總數
	WorkoutCount int `json:"workout_count" gorm:"column:workout_count" example:"10"`         // 訓練總數
	CreateAt string `json:"create_at" gorm:"column:create_at" example:"2021-05-28 11:00:00"` // 創建時間
	UpdateAt string `json:"update_at" gorm:"column:update_at" example:"2021-05-29 11:00:00"` // 更新時間
}

type CourseProduct struct {
	ID       int64                `json:"id" gorm:"column:id" example:"2"`                       // 課表 id
	Trainer  *TrainerSummary       `json:"trainer"`                                               // 教練簡介
	Sale     *SaleItem            `json:"sale"`                                                  // 銷售項目
	Restricted int                `json:"restricted" example:"0"`                                 // 是否是限制訪問狀態(0:否/1:是)
	CourseStatus int              `json:"course_status" gorm:"column:course_status" example:"1"` // 課表狀態(1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category int                  `json:"category" gorm:"column:category" example:"3"`           // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int              `json:"schedule_type" gorm:"column:schedule_type" example:"2"` // 排課類別(1:單一訓練/2:多項計畫)
	Name string                   `json:"name" gorm:"column:name" example:"Henry課表"`                    // 課表名稱
	Cover string                  `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`         // 課表封面
	Intro string                  `json:"intro" gorm:"column:intro" example:"佛系課表"`                     // 課表介紹
	Level int                     `json:"level" gorm:"column:level" example:"3"`                        // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit string                   `json:"suit" gorm:"column:suit" example:"2,5,7"`                      // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment string              `json:"equipment" gorm:"column:equipment" example:"2,3,6"`            // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place string                  `json:"place" gorm:"column:place" example:"1,2"`                      // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget string            `json:"train_target" gorm:"column:train_target" example:"2,3,4"`      // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	PlanCount int                 `json:"plan_count" gorm:"column:plan_count" example:"2"`              // 計畫總數
	WorkoutCount int              `json:"workout_count" gorm:"column:workout_count" example:"10"`       // 訓練總數
	Plans []*Plan                 `json:"plans"`                                                        // 計畫內容
	Review ReviewStatistic        `json:"review"`                                                       // 評分統計
}

type CourseProductSummary struct {
	ID       int64                `json:"id" gorm:"column:id" example:"2"`                                 // 課表 id
	Trainer  TrainerSummary       `json:"trainer"`                                            // 教練簡介
	Sale     *SaleItem            `json:"sale"`                                                // 銷售項目
	CourseStatus int              `json:"course_status" gorm:"column:course_status" example:"1"`           // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category int                  `json:"category" gorm:"column:category" example:"3"`                    // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int              `json:"schedule_type" gorm:"column:schedule_type" example:"2"`           // 排課類別(1:單一訓練/2:多項計畫)
	Name string                   `json:"name" gorm:"column:name" example:"Henry課表"`                            // 課表名稱
	Cover string                  `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`                          // 課表封面
	Level int                     `json:"level" gorm:"column:level" example:"3"`                          // 強度(1:初級/2:中級/3:中高級/4:高級)
	PlanCount int                 `json:"plan_count" gorm:"column:plan_count" example:"2"`                 // 計畫總數
	WorkoutCount int              `json:"workout_count" gorm:"column:workout_count" example:"10"`           // 訓練總數
	Review ReviewStatisticSummary `json:"review"`                                              // 評分統計
}

type GetCourseProductSummariesParam struct {
	Name *string `form:"name" binding:"omitempty,min=1,max=20" example:"增肌課表"` //課表名稱(1~20字元)
	OrderType *string `form:"order_type" binding:"omitempty,oneof=latest popular" example:"latest"` // 排序類型(latest:最新/popular:熱門)-單選
	Score *int `form:"score" binding:"omitempty,min=1,max=5" example:"5"` // 評價(1~5分)-單選
	Level []int `form:"level" binding:"omitempty" example:"3"` // 強度(1:初級/2:中級/3:中高級/4:高級)-複選
	Category []int `form:"category" binding:"omitempty,oneof=1 2 3 4 5 6" example:"3"` // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)-複選
	Suit []int `form:"suit" binding:"omitempty,oneof=1 2 3 4 5 6 7 8 9 10" example:"7"` // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)-複選
	Equipment []int `form:"equipment" binding:"omitempty,oneof=1 2 3 4 5 6 7 8 9" example:"5"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)-複選
	Place []int `form:"place" binding:"omitempty,oneof=1 2 3 4 5" example:"3"` // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)-複選
	TrainTarget []int `form:"train_target" binding:"omitempty,oneof=1 2 3 4 5" example:"4"` // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)-複選
	BodyTarget []int `form:"body_target" binding:"omitempty,oneof=1 2 3 4 5 6 7" example:",6"` // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)-複選
	SaleType []int `form:"sale_type" binding:"omitempty,oneof=1 2 3" example:"2"` // 銷售類型(1:免費課表/2:付費課表/3:訂閱課表)-複選
	TrainerSex []string `form:"trainer_sex" binding:"omitempty,oneof=m f" example:"m"` // 教練性別(m:男性/f:女性)-複選
	TrainerSkill  []int  `form:"trainer_skill" binding:"omitempty,oneof=1 2 3 4 5 6 7 8 9 10 11 12 13 14" example:"1"`  // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
}

type CourseID struct {
	ID       int64  `json:"course_id" gorm:"column:id" example:"2"`                           // 課表 id
}