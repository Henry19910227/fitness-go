package course_training_monthly_statistic

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //報表id
}
type YearField struct {
	Year *int `json:"year,omitempty" gorm:"column:year" example:"2022"` //年份
}
type MonthField struct {
	Month *int `json:"month,omitempty" gorm:"column:month" example:"12"` //月份
}
type TotalField struct {
	Total *int `json:"total,omitempty" gorm:"column:total" example:"1000"` //當月總訓練數
}
type FreeField struct {
	Free *int `json:"free,omitempty" gorm:"column:free" example:"200"` //免費課表訓練數
}
type SubscribeField struct {
	Subscribe *int `json:"subscribe,omitempty" gorm:"column:subscribe" example:"400"` //訂閱課表訓練數
}
type ChargeField struct {
	Charge *int `json:"charge,omitempty" gorm:"column:charge" example:"400"` //付費課表訓練數
}
type AerobicField struct {
	Aerobic *int `json:"aerobic,omitempty" gorm:"column:charge" example:"100"` //有氧課表訓練數
}
type IntervalTrainingField struct {
	IntervalTraining *int `json:"interval_training,omitempty" gorm:"column:interval_training" example:"100"` //間歇肌力訓練課表訓練數
}
type WeightTrainingField struct {
	WeightTraining *int `json:"weight_training,omitempty" gorm:"column:weight_training" example:"200"` //重量訓練課表訓練數
}
type ResistanceTrainingField struct {
	ResistanceTraining *int `json:"resistance_training,omitempty" gorm:"column:resistance_training" example:"200"` //阻力訓練課表訓練數
}
type BodyweightTrainingField struct {
	BodyweightTraining *int `json:"bodyweight_training,omitempty" gorm:"column:bodyweight_training" example:"200"` //徒手訓練課表訓練數
}
type OtherTrainingField struct {
	OtherTraining *int `json:"other_training,omitempty" gorm:"column:other_training" example:"200"` //付費課表訓練數
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	YearField
	MonthField
	TotalField
	FreeField
	SubscribeField
	ChargeField
	AerobicField
	IntervalTrainingField
	WeightTrainingField
	ResistanceTrainingField
	BodyweightTrainingField
	OtherTrainingField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "course_training_monthly_statistics"
}
