package required

type IDField struct {
	ID int64 `json:"id" uri:"course_id" form:"course_id" gorm:"column:id" binding:"required" example:"2"` // 課表 id
}
type UserIDField  struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶 id
}
type SaleTypeField  struct {
	SaleType int `json:"sale_type" form:"sale_type" gorm:"column:sale_type" binding:"required,oneof=1 2 3 4" example:"3"` // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表/4:個人課表)
}
type SaleIDField  struct {
	SaleID int64 `json:"sale_id" gorm:"column:sale_id" example:"3"` // 銷售 id
}
type CourseStatusField  struct {
	CourseStatus int `json:"course_status" form:"course_status" gorm:"column:course_status" binding:"required,oneof=1 2 3 4 5" example:"3"` // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
}
type CategoryField  struct {
	Category int `json:"category" gorm:"column:category" binding:"required,oneof=1 2 3 4 5 6" example:"1"` // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
}
type ScheduleTypeField  struct {
	ScheduleType int `json:"schedule_type" gorm:"column:schedule_type" binding:"required,oneof=1 2" example:"2"` // 排課類別(1:單一訓練/2:多項計畫)
}
type NameField  struct {
	Name string `json:"name" form:"name" gorm:"column:name" binding:"required,min=1,max=40" example:"增肌課表"` // 課表名稱
}
type CoverField  struct {
	Cover string `json:"cover" gorm:"column:cover" binding:"required" example:"abc.png"` // 課表封面
}
type IntroField  struct {
	Intro string `json:"intro" gorm:"column:intro" binding:"required" example:"增肌專用課表"` // 課表介紹
}
type FoodField  struct {
	Food string `json:"food" gorm:"column:food" binding:"required" example:"多吃雞胸肉"` // 飲食建議
}
type LevelField  struct {
	Level int `json:"level" gorm:"column:level" binding:"required,oneof=1 2 3 4" example:"4"` // 強度(1:初級/2:中級/3:中高級/4:高級)
}
type SuitField  struct {
	Suit string `json:"suit" gorm:"column:suit" binding:"required" example:"2"` // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
}
type EquipmentField  struct {
	Equipment string `json:"equipment" gorm:"column:equipment" binding:"required,equipment,max=5" example:"2,3,6"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
type PlaceField  struct {
	Place string `json:"place" gorm:"column:place" binding:"required,place,max=5" example:"1,2"` // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
}
type TrainTargetField  struct {
	TrainTarget string `json:"train_target" gorm:"column:train_target" binding:"required,train_target,max=5" example:"2,3,4"` // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
}
type BodyTargetField  struct {
	BodyTarget string `json:"body_target" gorm:"column:body_target" example:"4,5"` // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
}
type NoticeField  struct {
	Notice string `json:"notice" gorm:"column:notice" example:"注意關節避免鎖死"` // 注意事項
}
type PlanCountField  struct {
	PlanCount int `json:"plan_count" gorm:"column:plan_count" example:"10"` // 計畫總數
}
type WorkoutCountField  struct {
	WorkoutCount int `json:"workout_count" gorm:"column:workout_count" example:"50"` // 訓練總數
}
type CreateAtField  struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField  struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" example:"2022-06-12 00:00:00"` // 更新時間
}
