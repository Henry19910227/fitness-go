package course

type IDOptional struct {
	ID *int64 `json:"id,omitempty" uri:"course_id" form:"course_id" gorm:"column:id" binding:"omitempty" example:"2"` // 課表 id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` // 用戶 id
}
type SaleTypeOptional struct {
	SaleType *int `json:"sale_type,omitempty" form:"sale_type" gorm:"column:sale_type" binding:"omitempty,oneof=1 2 3" example:"3"` // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表/4:個人課表)
}
type CourseStatusOptional struct {
	CourseStatus *int `json:"course_status,omitempty" form:"course_status" gorm:"column:course_status" binding:"omitempty,oneof=1 2 3 4 5" example:"3"` // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
}
type CategoryOptional struct {
	Category *int `json:"category,omitempty" gorm:"column:category" example:"1"` // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
}
type ScheduleTypeOptional struct {
	ScheduleType *int `json:"schedule_type,omitempty" gorm:"column:schedule_type" example:"2"` // 排課類別(1:單一訓練/2:多項計畫)
}
type NameOptional struct {
	Name *string `json:"name,omitempty" form:"name" gorm:"column:name" binding:"omitempty" example:"增肌課表"` // 課表名稱
}
