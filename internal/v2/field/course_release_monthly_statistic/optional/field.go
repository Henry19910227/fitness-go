package optional

type YearField struct {
	Year *int `json:"year,omitempty" form:"year" gorm:"column:year" binding:"omitempty,max=2500" example:"2022"` //年份
}
type MonthField struct {
	Month *int `json:"month,omitempty" form:"month" gorm:"column:month" binding:"omitempty,min=1,max=12" example:"12"` //月份
}
type TotalField struct {
	Total *int `json:"total,omitempty" gorm:"column:total" binding:"omitempty" example:"1000"` //當月總上架數
}
type FreeField struct {
	Free *int `json:"free,omitempty" gorm:"column:free" binding:"omitempty" example:"200"` //免費課表上架數
}
type SubscribeField struct {
	Subscribe *int `json:"subscribe,omitempty" gorm:"column:subscribe" binding:"omitempty" example:"400"` //訂閱課表上架數
}
type ChargeField struct {
	Charge *int `json:"charge,omitempty" gorm:"column:charge" binding:"omitempty" example:"400"` //付費課表上架數
}
type AerobicField struct {
	Aerobic *int `json:"aerobic,omitempty" gorm:"column:charge" binding:"omitempty" example:"100"` //有氧課表上架數
}
type IntervalTrainingField struct {
	IntervalTraining *int `json:"interval_training,omitempty" gorm:"column:interval_training" binding:"omitempty" example:"100"` //間歇肌力訓練課表上架數
}
type WeightTrainingField struct {
	WeightTraining *int `json:"weight_training,omitempty" gorm:"column:weight_training" binding:"omitempty" example:"200"` //重量訓練課表上架數
}
type ResistanceTrainingField struct {
	ResistanceTraining *int `json:"resistance_training,omitempty" gorm:"column:resistance_training" binding:"omitempty" example:"200"` //阻力訓練課表上架數
}
type BodyweightTrainingField struct {
	BodyweightTraining *int `json:"bodyweight_training,omitempty" gorm:"column:bodyweight_training" binding:"omitempty" example:"200"` //徒手訓練課表上架數
}
type OtherTrainingField struct {
	OtherTraining *int `json:"other_training,omitempty" gorm:"column:other_training" binding:"omitempty" example:"200"` //付費課表上架數
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
