package dto

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"mime/multipart"
)

type CourseCover struct {
	Cover string `json:"cover" example:"dkf2se51fsdds.png"` // 課表封面照片
}

type CreateCourseParam struct {
	Name          string
	Level         int
	Category      int
	CategoryOther string
	ScheduleType  int
}

type UpdateCourseParam struct {
	Category    *int    `gorm:"column:category"`     // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	SaleType    *int    `gorm:"column:sale_type"`    // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	SaleID      *int64  `gorm:"column:sale_id"`      // 銷售id
	Name        *string `gorm:"column:name"`         // 課表名稱
	Intro       *string `gorm:"column:intro"`        // 課表介紹
	Food        *string `gorm:"column:food"`         // 飲食建議
	Level       *int    `gorm:"column:level"`        // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit        *string `gorm:"column:suit"`         // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment   *string `gorm:"column:equipment"`    // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place       *string `gorm:"column:place"`        // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget *string `gorm:"column:train_target"` // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget  *string `gorm:"column:body_target"`  // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice      *string `gorm:"column:notice"`       // 注意事項
	UpdateAt    *string `gorm:"column:update_at"`    // 更新時間
}

type UploadCourseCoverParam struct {
	CoverNamed string
	File       multipart.File
}

type CourseSummary struct {
	ID           int64           `json:"id" example:"2"`                                                  // 課表 id
	Trainer      *TrainerSummary `json:"trainer"`                                                         // 教練簡介
	SaleType     int             `json:"sale_type" example:"1"`                                           // 銷售類型
	Sale         *SaleItem       `json:"sale"`                                                            // 銷售資料
	CourseStatus int             `json:"course_status" example:"1"`                                       // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category     int             `json:"category" gorm:"column:category" example:"3"`                     // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int             `json:"schedule_type" gorm:"column:schedule_type" example:"2"`           // 排課類別(1:單一訓練/2:多項計畫)
	Name         string          `json:"name" example:"Henry課表"`                                          // 課表名稱
	Cover        string          `json:"cover" example:"d2w3e15d3awe.jpg"`                                // 課表封面
	Level        int             `json:"level" example:"3"`                                               // 強度(1:初級/2:中級/3:中高級/4:高級)
	PlanCount    int             `json:"plan_count" gorm:"column:plan_count" example:"2"`                 // 計畫總數
	WorkoutCount int             `json:"workout_count" gorm:"column:workout_count" example:"10"`          // 訓練總數
	CreateAt     string          `json:"create_at" gorm:"column:create_at" example:"2021-06-01 12:00:00"` // 創建日期
	UpdateAt     string          `json:"update_at" gorm:"column:update_at" example:"2021-06-01 12:00:00"` // 修改日期
}

type Course struct {
	ID           int64           `json:"id" gorm:"column:id" example:"2"`                                 // 課表 id
	Trainer      *TrainerSummary `json:"trainer"`                                                         // 教練簡介
	SaleType     int             `json:"sale_type" example:"2"`                                           // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Sale         *SaleItem       `json:"sale"`                                                            // 銷售資料
	CourseStatus int             `json:"course_status" gorm:"column:course_status" example:"1"`           // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category     int             `json:"category" gorm:"column:category" example:"3"`                     // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int             `json:"schedule_type" gorm:"column:schedule_type" example:"2"`           // 排課類別(1:單一訓練/2:多項計畫)
	Name         string          `json:"name" gorm:"column:name" example:"Henry課表"`                       // 課表名稱
	Cover        string          `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`            // 課表封面
	Intro        string          `json:"intro" gorm:"column:intro" example:"佛系課表"`                        // 課表介紹
	Food         string          `json:"food" gorm:"column:food" example:"佛系飲食"`                          // 飲食建議
	Level        int             `json:"level" gorm:"column:level" example:"3"`                           // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit         string          `json:"suit" gorm:"column:suit" example:"2,5,7"`                         // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment    string          `json:"equipment" gorm:"column:equipment" example:"2,3,6"`               // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place        string          `json:"place" gorm:"column:place" example:"1,2"`                         // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget  string          `json:"train_target" gorm:"column:train_target" example:"2,3,4"`         // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget   string          `json:"body_target" gorm:"column:body_target" example:"4,5"`             // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice       string          `json:"notice" gorm:"column:notice" example:"小心不要受傷"`                    // 注意事項
	PlanCount    int             `json:"plan_count" gorm:"column:plan_count" example:"2"`                 // 計畫總數
	WorkoutCount int             `json:"workout_count" gorm:"column:workout_count" example:"10"`          // 訓練總數
	CreateAt     string          `json:"create_at" gorm:"column:create_at" example:"2021-05-28 11:00:00"` // 創建時間
	UpdateAt     string          `json:"update_at" gorm:"column:update_at" example:"2021-05-29 11:00:00"` // 更新時間
}

type CourseProduct struct {
	ID           int64            `json:"id" gorm:"column:id" example:"2"`                         // 課表 id
	Trainer      *TrainerSummary  `json:"trainer"`                                                 // 教練簡介
	SaleType     int              `json:"sale_type" example:"2"`                                   // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Sale         *SaleItem        `json:"sale"`                                                    // 銷售項目
	AllowAccess  int              `json:"allow_access" example:"0"`                                // 是否允許訪問此課表(0:否/1:是)
	CourseStatus int              `json:"course_status" gorm:"column:course_status" example:"1"`   // 課表狀態(1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category     int              `json:"category" gorm:"column:category" example:"3"`             // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int              `json:"schedule_type" gorm:"column:schedule_type" example:"2"`   // 排課類別(1:單一訓練/2:多項計畫)
	Name         string           `json:"name" gorm:"column:name" example:"Henry課表"`               // 課表名稱
	Cover        string           `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`    // 課表封面
	Intro        string           `json:"intro" gorm:"column:intro" example:"佛系課表"`                // 課表介紹
	Food         string           `json:"food" gorm:"column:food" example:"佛系飲食"`                  // 飲食建議
	Level        int              `json:"level" gorm:"column:level" example:"3"`                   // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit         string           `json:"suit" gorm:"column:suit" example:"2,5,7"`                 // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment    string           `json:"equipment" gorm:"column:equipment" example:"2,3,6"`       // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place        string           `json:"place" gorm:"column:place" example:"1,2"`                 // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget  string           `json:"train_target" gorm:"column:train_target" example:"2,3,4"` // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget   string           `json:"body_target" gorm:"column:body_target" example:"4,5"`     // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice       string           `json:"notice" gorm:"column:notice" example:"小心不要受傷"`            // 注意事項
	Favorite     int              `json:"favorite" gorm:"column:favorite" example:"1"`             //  是否收藏(0:否/1:是)
	PlanCount    int              `json:"plan_count" gorm:"column:plan_count" example:"2"`         // 計畫總數
	WorkoutCount int              `json:"workout_count" gorm:"column:workout_count" example:"10"`  // 訓練總數
	Review       *ReviewStatistic `json:"review"`                                                  // 評分統計
}

type CourseProductStructure struct {
	ID           int64            `json:"id" gorm:"column:id" example:"2"`                         // 課表 id
	Trainer      *TrainerSummary  `json:"trainer"`                                                 // 教練簡介
	SaleType     int              `json:"sale_type" example:"2"`                                   // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Sale         *SaleItem        `json:"sale"`                                                    // 銷售項目
	AllowAccess  int              `json:"allow_access" example:"0"`                                // 是否允許訪問此課表(0:否/1:是)
	CourseStatus int              `json:"course_status" gorm:"column:course_status" example:"1"`   // 課表狀態(1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category     int              `json:"category" gorm:"column:category" example:"3"`             // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int              `json:"schedule_type" gorm:"column:schedule_type" example:"2"`   // 排課類別(1:單一訓練/2:多項計畫)
	Name         string           `json:"name" gorm:"column:name" example:"Henry課表"`               // 課表名稱
	Cover        string           `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`    // 課表封面
	Intro        string           `json:"intro" gorm:"column:intro" example:"佛系課表"`                // 課表介紹
	Food         string           `json:"food" gorm:"column:food" example:"佛系飲食"`                  // 飲食建議
	Level        int              `json:"level" gorm:"column:level" example:"3"`                   // 強度(1:初級/2:中級/3:中高級/4:高級)
	Suit         string           `json:"suit" gorm:"column:suit" example:"2,5,7"`                 // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment    string           `json:"equipment" gorm:"column:equipment" example:"2,3,6"`       // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place        string           `json:"place" gorm:"column:place" example:"1,2"`                 // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget  string           `json:"train_target" gorm:"column:train_target" example:"2,3,4"` // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget   string           `json:"body_target" gorm:"column:body_target" example:"4,5"`     // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
	Notice       string           `json:"notice" gorm:"column:notice" example:"小心不要受傷"`            // 注意事項
	Favorite     int              `json:"favorite" gorm:"column:favorite" example:"1"`             //  是否收藏(0:否/1:是)
	PlanCount    int              `json:"plan_count" gorm:"column:plan_count" example:"2"`         // 計畫總數
	WorkoutCount int              `json:"workout_count" gorm:"column:workout_count" example:"10"`  // 訓練總數
	Review       *ReviewStatistic `json:"review"`                                                  // 評分統計
	Plans        []*PlanStructure `json:"plans"`                                                   // 計畫列表
}

type CourseProductSummary struct {
	ID           int64                  `json:"id" gorm:"column:id" example:"2"`                        // 課表 id
	Trainer      *TrainerSummary        `json:"trainer"`                                                // 教練簡介
	SaleType     int                    `json:"sale_type" example:"2"`                                  // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Sale         *SaleItem              `json:"sale"`                                                   // 銷售項目
	CourseStatus int                    `json:"course_status" gorm:"column:course_status" example:"1"`  // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category     int                    `json:"category" gorm:"column:category" example:"3"`            // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int                    `json:"schedule_type" gorm:"column:schedule_type" example:"2"`  // 排課類別(1:單一訓練/2:多項計畫)
	Name         string                 `json:"name" gorm:"column:name" example:"Henry課表"`              // 課表名稱
	Cover        string                 `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`   // 課表封面
	Level        int                    `json:"level" gorm:"column:level" example:"3"`                  // 強度(1:初級/2:中級/3:中高級/4:高級)
	PlanCount    int                    `json:"plan_count" gorm:"column:plan_count" example:"2"`        // 計畫總數
	WorkoutCount int                    `json:"workout_count" gorm:"column:workout_count" example:"10"` // 訓練總數
	Review       ReviewStatisticSummary `json:"review"`                                                 // 評分統計
}

type CourseAsset struct {
	ID              int64                `json:"id" gorm:"column:id" example:"2"`                        // 課表 id
	Trainer         *TrainerSummary      `json:"trainer"`                                                // 教練簡介
	SaleType        int                  `json:"sale_type" example:"2"`                                  // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Sale            *SaleItem            `json:"sale"`                                                   // 銷售項目
	AllowAccess     int                  `json:"allow_access" example:"0"`                               // 是否允許訪問此課表(0:否/1:是)
	CourseStatus    int                  `json:"course_status" gorm:"column:course_status" example:"1"`  // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category        int                  `json:"category" gorm:"column:category" example:"3"`            // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType    int                  `json:"schedule_type" gorm:"column:schedule_type" example:"2"`  // 排課類別(1:單一訓練/2:多項計畫)
	Name            string               `json:"name" gorm:"column:name" example:"Henry課表"`              // 課表名稱
	Cover           string               `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`   // 課表封面
	Level           int                  `json:"level" gorm:"column:level" example:"3"`                  // 強度(1:初級/2:中級/3:中高級/4:高級)
	Favorite        int                  `json:"favorite" gorm:"column:favorite" example:"1"`            //  是否收藏(0:否/1:是)
	PlanCount       int                  `json:"plan_count" gorm:"column:plan_count" example:"2"`        // 計畫總數
	WorkoutCount    int                  `json:"workout_count" gorm:"column:workout_count" example:"10"` // 訓練總數
	CourseStatistic *UserCourseStatistic `json:"user_course_statistic"`                                  // 課表統計
}

type CourseAssetStructure struct {
	ID              int64                 `json:"id" gorm:"column:id" example:"2"`                        // 課表 id
	Trainer         *TrainerSummary       `json:"trainer"`                                                // 教練簡介
	SaleType        int                   `json:"sale_type" example:"2"`                                  // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Sale            *SaleItem             `json:"sale"`                                                   // 銷售項目
	AllowAccess     int                   `json:"allow_access" example:"0"`                               // 是否允許訪問此課表(0:否/1:是)
	CourseStatus    int                   `json:"course_status" gorm:"column:course_status" example:"1"`  // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category        int                   `json:"category" gorm:"column:category" example:"3"`            // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType    int                   `json:"schedule_type" gorm:"column:schedule_type" example:"2"`  // 排課類別(1:單一訓練/2:多項計畫)
	Name            string                `json:"name" gorm:"column:name" example:"Henry課表"`              // 課表名稱
	Cover           string                `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`   // 課表封面
	Level           int                   `json:"level" gorm:"column:level" example:"3"`                  // 強度(1:初級/2:中級/3:中高級/4:高級)
	Favorite        int                   `json:"favorite" gorm:"column:favorite" example:"1"`            //  是否收藏(0:否/1:是)
	PlanCount       int                   `json:"plan_count" gorm:"column:plan_count" example:"2"`        // 計畫總數
	WorkoutCount    int                   `json:"workout_count" gorm:"column:workout_count" example:"10"` // 訓練總數
	CourseStatistic *UserCourseStatistic  `json:"user_course_statistic"`                                  // 課表統計
	Plans           []*PlanAssetStructure `json:"plans"`                                                  // 計畫列表
}

type CourseAssetSummary struct {
	ID           int64                   `json:"id" gorm:"column:id" example:"2"`                        // 課表 id
	Trainer      *TrainerSummary         `json:"trainer"`                                                // 教練簡介
	SaleType     int                     `json:"sale_type" example:"2"`                                  // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Sale         *SaleItem               `json:"sale"`                                                   // 銷售項目
	CourseStatus int                     `json:"course_status" gorm:"column:course_status" example:"1"`  // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category     int                     `json:"category" gorm:"column:category" example:"3"`            // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType int                     `json:"schedule_type" gorm:"column:schedule_type" example:"2"`  // 排課類別(1:單一訓練/2:多項計畫)
	Name         string                  `json:"name" gorm:"column:name" example:"Henry課表"`              // 課表名稱
	Cover        string                  `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`   // 課表封面
	Level        int                     `json:"level" gorm:"column:level" example:"3"`                  // 強度(1:初級/2:中級/3:中高級/4:高級)
	PlanCount    int                     `json:"plan_count" gorm:"column:plan_count" example:"2"`        // 計畫總數
	WorkoutCount int                     `json:"workout_count" gorm:"column:workout_count" example:"10"` // 訓練總數
	Review       *ReviewStatisticSummary `json:"review"`                                                 // 評分統計
}

type CourseStatistic struct {
	ID                   int64                   `json:"id" gorm:"column:id" example:"2"`                        // 課表 id
	UserID               int64                   `json:"user_id" gorm:"column:user_id" example:"10001"`          // 教練 id
	SaleType             int                     `json:"sale_type" example:"2"`                                  // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	CourseStatus         int                     `json:"course_status" gorm:"column:course_status" example:"1"`  // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category             int                     `json:"category" gorm:"column:category" example:"3"`            // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType         int                     `json:"schedule_type" gorm:"column:schedule_type" example:"2"`  // 排課類別(1:單一訓練/2:多項計畫)
	Name                 string                  `json:"name" gorm:"column:name" example:"Henry課表"`              // 課表名稱
	Cover                string                  `json:"cover" gorm:"column:cover" example:"d2w3e15d3awe.jpg"`   // 課表封面
	Level                int                     `json:"level" gorm:"column:level" example:"3"`                  // 強度(1:初級/2:中級/3:中高級/4:高級)
	PlanCount            int                     `json:"plan_count" gorm:"column:plan_count" example:"2"`        // 計畫總數
	WorkoutCount         int                     `json:"workout_count" gorm:"column:workout_count" example:"10"` // 訓練總數
	Review               *ReviewStatisticSummary `json:"review,omitempty" gorm:"-"`                              // 評分統計
	CourseUsageStatistic *CourseUsageStatistic   `json:"course_usage_statistic,omitempty" gorm:"-"`              // 課表使用統計
}

type CourseStatisticSummary struct {
	ID                   int64                        `json:"id" example:"1"`                                        // 課表 id
	UserID               int64                        `json:"user_id" example:"10001"`                               // 教練 id
	SaleType             int                          `json:"sale_type" example:"1"`                                 // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	CourseStatus         int                          `json:"course_status" example:"1"`                             // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	Category             int                          `json:"category" example:"3"`                                  // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	ScheduleType         int                          `json:"schedule_type" gorm:"column:schedule_type" example:"2"` // 排課類別(1:單一訓練/2:多項計畫)
	Name                 string                       `json:"name" example:"Henry課表"`                                // 課表名稱
	Cover                string                       `json:"cover" example:"d2w3e15d3awe.jpg"`                      // 課表封面
	Level                int                          `json:"level" example:"3"`                                     // 強度(1:初級/2:中級/3:中高級/4:高級)
	PlanCount            int                          `json:"plan_count" example:"100"`                              // 計畫總數
	WorkoutCount         int                          `json:"workout_count" example:"10"`                            // 訓練總數
	CourseUsageStatistic *CourseUsageStatisticSummary `json:"course_usage_statistic,omitempty"`                      // 課表使用統計
}

type UserCourseStatistic struct {
	FinishWorkoutCourt int `json:"finish_workout_count" example:"10"` // 完成訓練數量(去除重複)
	Duration           int `json:"duration" example:"3600"`           // 總花費時間(秒)
}

type CourseProductItem struct {
	ID       int64  `json:"id" example:"10001"`             // 課表 id
	SaleType int    `json:"sale_type" example:"1"`          // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Name     string `json:"name" example:"henry"`           // 課表名稱
	Cover    string `json:"cover" example:"f43e5715fe.jpg"` // 課表封面
}

type CourseCustom1 struct {
	ID           int64  `json:"id" example:"10001"`                                    // 課表 id
	SaleType     int    `json:"sale_type" example:"1"`                                 // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	ScheduleType int    `json:"schedule_type" gorm:"column:schedule_type" example:"2"` // 排課類別(1:單一訓練/2:多項計畫)
	Name         string `json:"name" example:"henry"`                                  // 課表名稱
	Cover        string `json:"cover" example:"f43e5715fe.jpg"`                        // 課表封面
}

type GetCourseProductSummariesParam struct {
	UserID       *int64   `form:"user_id" binding:"omitempty" example:"10001"`                                          //教練ID
	Name         *string  `form:"name" binding:"omitempty,min=1,max=20" example:"增肌課表"`                                 //課表名稱(1~20字元)
	OrderType    *string  `form:"order_type" binding:"omitempty,oneof=latest popular" example:"latest"`                 // 排序類型(latest:最新/popular:熱門)-單選
	Score        *int     `form:"score" binding:"omitempty,min=1,max=5" example:"5"`                                    // 評價(1~5分)-單選
	Level        []int    `form:"level" binding:"omitempty" example:"3"`                                                // 強度(1:初級/2:中級/3:中高級/4:高級)-複選
	Category     []int    `form:"category" binding:"omitempty,oneof=1 2 3 4 5 6" example:"3"`                           // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)-複選
	Suit         []int    `form:"suit" binding:"omitempty,oneof=1 2 3 4 5 6 7 8 9 10" example:"7"`                      // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)-複選
	Equipment    []int    `form:"equipment" binding:"omitempty,oneof=1 2 3 4 5 6 7 8 9" example:"5"`                    // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)-複選
	Place        []int    `form:"place" binding:"omitempty,oneof=1 2 3 4 5" example:"3"`                                // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)-複選
	TrainTarget  []int    `form:"train_target" binding:"omitempty,oneof=1 2 3 4 5" example:"4"`                         // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)-複選
	BodyTarget   []int    `form:"body_target" binding:"omitempty,oneof=1 2 3 4 5 6 7" example:",6"`                     // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)-複選
	SaleType     []int    `form:"sale_type" binding:"omitempty,oneof=1 2 3" example:"2"`                                // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)-複選
	TrainerSex   []string `form:"trainer_sex" binding:"omitempty,oneof=m f" example:"m"`                                // 教練性別(m:男性/f:女性)-複選
	TrainerSkill []int    `form:"trainer_skill" binding:"omitempty,oneof=1 2 3 4 5 6 7 8 9 10 11 12 13 14" example:"1"` // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
}

type CourseID struct {
	ID int64 `json:"course_id" gorm:"column:id" example:"2"` // 課表 id
}

func NewCourseProduct(data *model.CourseProduct) CourseProduct {
	course := CourseProduct{
		ID:           data.ID,
		CourseStatus: data.CourseStatus,
		Category:     data.Category,
		SaleType:     data.SaleType,
		ScheduleType: data.ScheduleType,
		Name:         data.Name,
		Cover:        data.Cover,
		Intro:        data.Intro,
		Food:         data.Food,
		Level:        data.Level,
		Suit:         data.Suit,
		Equipment:    data.Equipment,
		Place:        data.Place,
		TrainTarget:  data.TrainTarget,
		BodyTarget:   data.BodyTarget,
		Notice:       data.Notice,
		PlanCount:    data.PlanCount,
		WorkoutCount: data.WorkoutCount,
	}
	//配置教練資訊
	trainer := &TrainerSummary{
		UserID:   data.Trainer.UserID,
		Nickname: data.Trainer.Nickname,
		Avatar:   data.Trainer.Avatar,
		Skill:    data.Trainer.Skill,
	}
	course.Trainer = trainer
	//配置銷售資訊
	if data.Sale != nil {
		sale := &SaleItem{
			ID:   data.Sale.ID,
			Type: data.Sale.Type,
			Name: data.Sale.Name,
		}
		course.Sale = sale
		if data.Sale.ProductLabel != nil {
			course.Sale.Twd = data.Sale.ProductLabel.Twd
			course.Sale.ProductID = data.Sale.ProductLabel.ProductID
		}
	}
	//配置評論統計
	course.Review = &ReviewStatistic{}
	if data.Review != nil {
		course.Review.ScoreTotal = data.Review.ScoreTotal
		course.Review.Amount = data.Review.Amount
		course.Review.FiveTotal = data.Review.FiveTotal
		course.Review.FourTotal = data.Review.FourTotal
		course.Review.ThreeTotal = data.Review.ThreeTotal
		course.Review.TwoTotal = data.Review.TwoTotal
		course.Review.OneTotal = data.Review.OneTotal
		course.Review.UpdateAt = data.Review.UpdateAt
	}
	return course
}

func NewCourseProductStructure(data *model.CourseProduct) CourseProductStructure {
	course := CourseProductStructure{
		ID:           data.ID,
		CourseStatus: data.CourseStatus,
		Category:     data.Category,
		SaleType:     data.SaleType,
		ScheduleType: data.ScheduleType,
		Name:         data.Name,
		Cover:        data.Cover,
		Intro:        data.Intro,
		Food:         data.Food,
		Level:        data.Level,
		Suit:         data.Suit,
		Equipment:    data.Equipment,
		Place:        data.Place,
		TrainTarget:  data.TrainTarget,
		BodyTarget:   data.BodyTarget,
		Notice:       data.Notice,
		PlanCount:    data.PlanCount,
		WorkoutCount: data.WorkoutCount,
	}
	//配置教練資訊
	trainer := &TrainerSummary{
		UserID:   data.Trainer.UserID,
		Nickname: data.Trainer.Nickname,
		Avatar:   data.Trainer.Avatar,
		Skill:    data.Trainer.Skill,
	}
	course.Trainer = trainer
	//配置銷售資訊
	if data.Sale != nil {
		sale := &SaleItem{
			ID:   data.Sale.ID,
			Type: data.Sale.Type,
			Name: data.Sale.Name,
		}
		course.Sale = sale
		if data.Sale.ProductLabel != nil {
			course.Sale.Twd = data.Sale.ProductLabel.Twd
			course.Sale.ProductID = data.Sale.ProductLabel.ProductID
		}
	}
	//配置評論統計
	course.Review = &ReviewStatistic{}
	if data.Review != nil {
		course.Review.ScoreTotal = data.Review.ScoreTotal
		course.Review.Amount = data.Review.Amount
		course.Review.FiveTotal = data.Review.FiveTotal
		course.Review.FourTotal = data.Review.FourTotal
		course.Review.ThreeTotal = data.Review.ThreeTotal
		course.Review.TwoTotal = data.Review.TwoTotal
		course.Review.OneTotal = data.Review.OneTotal
		course.Review.UpdateAt = data.Review.UpdateAt
	}
	course.Plans = make([]*PlanStructure, 0)
	return course
}

func NewCourseAsset(data *model.CourseAsset) CourseAsset {
	course := CourseAsset{
		ID:           data.ID,
		CourseStatus: data.CourseStatus,
		Category:     data.Category,
		SaleType:     data.SaleType,
		ScheduleType: data.ScheduleType,
		Name:         data.Name,
		Cover:        data.Cover,
		Level:        data.Level,
		PlanCount:    data.PlanCount,
		WorkoutCount: data.WorkoutCount,
	}
	//配置教練資訊
	trainer := &TrainerSummary{
		UserID:   data.Trainer.UserID,
		Nickname: data.Trainer.Nickname,
		Avatar:   data.Trainer.Avatar,
		Skill:    data.Trainer.Skill,
	}
	course.Trainer = trainer
	//配置銷售資訊
	if data.Sale != nil {
		sale := &SaleItem{
			ID:   data.Sale.ID,
			Type: data.Sale.Type,
			Name: data.Sale.Name,
		}
		course.Sale = sale
		if data.Sale.ProductLabel != nil {
			course.Sale.Twd = data.Sale.ProductLabel.Twd
			course.Sale.ProductID = data.Sale.ProductLabel.ProductID
		}
	}
	//配置個人課表統計
	course.CourseStatistic = &UserCourseStatistic{
		FinishWorkoutCourt: data.FinishWorkoutCourt,
		Duration:           data.Duration,
	}
	return course
}

func NewCourseAssetStructure(data *model.CourseAsset) CourseAssetStructure {
	course := CourseAssetStructure{
		ID:           data.ID,
		CourseStatus: data.CourseStatus,
		Category:     data.Category,
		SaleType:     data.SaleType,
		ScheduleType: data.ScheduleType,
		Name:         data.Name,
		Cover:        data.Cover,
		Level:        data.Level,
		PlanCount:    data.PlanCount,
		WorkoutCount: data.WorkoutCount,
	}
	course.Plans = make([]*PlanAssetStructure, 0)
	//配置教練資訊
	trainer := &TrainerSummary{
		UserID:   data.Trainer.UserID,
		Nickname: data.Trainer.Nickname,
		Avatar:   data.Trainer.Avatar,
		Skill:    data.Trainer.Skill,
	}
	course.Trainer = trainer
	//配置銷售資訊
	if data.Sale != nil {
		sale := &SaleItem{
			ID:   data.Sale.ID,
			Type: data.Sale.Type,
			Name: data.Sale.Name,
		}
		course.Sale = sale
		if data.Sale.ProductLabel != nil {
			course.Sale.Twd = data.Sale.ProductLabel.Twd
			course.Sale.ProductID = data.Sale.ProductLabel.ProductID
		}
	}
	//配置個人課表統計
	course.CourseStatistic = &UserCourseStatistic{
		FinishWorkoutCourt: data.FinishWorkoutCourt,
		Duration:           data.Duration,
	}
	return course
}

func NewCourseAssetSummary(data *model.CourseAssetSummary) CourseAssetSummary {
	course := CourseAssetSummary{
		ID:           data.ID,
		SaleType:     data.SaleType,
		CourseStatus: data.CourseStatus,
		Category:     data.Category,
		ScheduleType: data.ScheduleType,
		Name:         data.Name,
		Cover:        data.Cover,
		Level:        data.Level,
		PlanCount:    data.PlanCount,
		WorkoutCount: data.WorkoutCount,
	}
	if data.Trainer != nil {
		course.Trainer = &TrainerSummary{
			UserID:   data.Trainer.UserID,
			Nickname: data.Trainer.Nickname,
			Avatar:   data.Trainer.Avatar,
			Skill:    data.Trainer.Skill,
		}
	}
	course.Review = &ReviewStatisticSummary{}
	if data.Review != nil {
		course.Review.Amount = data.Review.Amount
		course.Review.ScoreTotal = data.Review.ScoreTotal
	}
	if data.Sale != nil {
		sale := &SaleItem{
			ID:   data.Sale.ID,
			Type: data.Sale.Type,
			Name: data.Sale.Name,
		}
		course.Sale = sale
		if data.Sale.ProductLabel != nil {
			course.Sale.Twd = data.Sale.ProductLabel.Twd
			course.Sale.ProductID = data.Sale.ProductLabel.ProductID
		}
	}
	return course
}

func NewCourseSummary(data *model.CourseSummary) *CourseSummary {
	course := CourseSummary{
		ID:           data.ID,
		SaleType:     data.SaleType,
		CourseStatus: data.CourseStatus,
		Category:     data.Category,
		ScheduleType: data.ScheduleType,
		Name:         data.Name,
		Cover:        data.Cover,
		Level:        data.Level,
		PlanCount:    data.PlanCount,
		WorkoutCount: data.WorkoutCount,
		CreateAt:     data.CreateAt,
		UpdateAt:     data.UpdateAt,
	}
	trainer := &TrainerSummary{
		UserID:   data.Trainer.UserID,
		Nickname: data.Trainer.Nickname,
		Avatar:   data.Trainer.Avatar,
		Skill:    data.Trainer.Skill,
	}
	course.Trainer = trainer
	if data.Sale != nil {
		sale := &SaleItem{
			ID:   data.Sale.ID,
			Type: data.Sale.Type,
		}
		course.Sale = sale
		if data.Sale.ProductLabel != nil {
			course.Sale.Name = data.Sale.ProductLabel.Name
			course.Sale.Twd = data.Sale.ProductLabel.Twd
			course.Sale.ProductID = data.Sale.ProductLabel.ProductID
		}
	}
	return &course
}

func NewCourseStatisticSummary(data *model.CourseStatisticSummary) *CourseStatisticSummary {
	course := CourseStatisticSummary{
		ID:           data.ID,
		UserID:       data.UserID,
		SaleType:     data.SaleType,
		CourseStatus: data.CourseStatus,
		Category:     data.Category,
		ScheduleType: data.ScheduleType,
		Name:         data.Name,
		Cover:        data.Cover,
		Level:        data.Level,
		PlanCount:    data.PlanCount,
		WorkoutCount: data.WorkoutCount,
	}
	course.CourseUsageStatistic = &CourseUsageStatisticSummary{}
	if data.CourseUsageStatistic != nil {
		course.CourseUsageStatistic.TotalFinishWorkoutCount = data.CourseUsageStatistic.TotalFinishWorkoutCount
		course.CourseUsageStatistic.UserFinishCount = data.CourseUsageStatistic.UserFinishCount
		course.CourseUsageStatistic.FinishCountAvg = data.CourseUsageStatistic.FinishCountAvg
	}
	return &course
}
