package course

type IDOptional struct {
	ID *int64 `json:"id,omitempty" uri:"course_id" form:"course_id" binding:"omitempty" example:"2"` // 課表 id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" binding:"omitempty" example:"10001"` // 用戶 id
}
type SaleTypeOptional struct {
	SaleType *int `json:"sale_type,omitempty" form:"sale_type" binding:"omitempty,oneof=1 2 3 4" example:"3"` // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表/4:個人課表)
}
type SaleIDOptional struct {
	SaleID *int64 `json:"sale_id,omitempty" binding:"omitempty" example:"3"` // 銷售 id
}
type CourseStatusOptional struct {
	CourseStatus *int `json:"course_status,omitempty" form:"course_status" binding:"omitempty,oneof=1 2 3 4 5" example:"3"` // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
}
type CategoryOptional struct {
	Category *int `json:"category,omitempty" binding:"omitempty,oneof=0 1 2 3 4 5 6" example:"1"` // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
}
type ScheduleTypeOptional struct {
	ScheduleType *int `json:"schedule_type,omitempty" binding:"omitempty,oneof=1 2" example:"2"` // 排課類別(1:單一訓練/2:多項計畫)
}
type NameOptional struct {
	Name *string `json:"name,omitempty" form:"name" binding:"omitempty,min=1,max=40" example:"增肌課表"` // 課表名稱
}
type BodyTargetOptional struct {
	BodyTarget *string `json:"body_target,omitempty" binding:"omitempty,body_target,max=5" example:"4,5"` // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
}
type NoticeOptional struct {
	Notice *string `json:"notice,omitempty" binding:"omitempty" example:"注意關節避免鎖死"` // 注意事項
}
type EquipmentOptional struct {
	Equipment *string `json:"equipment,omitempty" binding:"omitempty,equipment,max=5" example:"2,3,6"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
type PlaceOptional struct {
	Place *string `json:"place,omitempty" binding:"omitempty,place,max=5" example:"1,2"` // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
}
type IntroOptional struct {
	Intro *string `json:"intro,omitempty" binding:"omitempty,max=400" example:"增肌專用課表"` // 課表介紹
}
type FoodOptional struct {
	Food *string `json:"food,omitempty" binding:"omitempty,max=400" example:"多吃雞胸肉"` // 飲食建議
}
type LevelOptional struct {
	Level *int `json:"level,omitempty" binding:"omitempty,oneof=0 1 2 3 4" example:"4"` // 強度(1:初級/2:中級/3:中高級/4:高級)
}
type SuitOptional struct {
	Suit *string `json:"suit,omitempty" binding:"omitempty,suit,max=5" example:"2,5,7"` // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
}
type TrainTargetOptional struct {
	TrainTarget *string `json:"train_target,omitempty" binding:"omitempty,train_target,max=5" example:"2,3,4"` // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
}
