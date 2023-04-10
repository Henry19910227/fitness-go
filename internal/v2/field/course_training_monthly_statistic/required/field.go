package required

type YearField struct {
	Year int `json:"year" form:"year" gorm:"column:year" binding:"required,max=2500" example:"2022"` //年份
}
type MonthField struct {
	Month int `json:"month" form:"month" gorm:"column:month" binding:"required,min=1,max=12" example:"12"` //月份
}
type TotalField struct {
	Total int `json:"total" gorm:"column:total" binding:"required" example:"1000"` //當月總訓練數
}
type FreeField struct {
	Free int `json:"free" gorm:"column:free" binding:"required" example:"200"` //免費課表訓練數
}
type SubscribeField struct {
	Subscribe int `json:"subscribe" gorm:"column:subscribe" binding:"required" example:"400"` //訂閱課表訓練數
}
type ChargeField struct {
	Charge int `json:"charge" gorm:"column:charge" binding:"required" example:"400"` //付費課表訓練數
}
type AerobicField struct {
	Aerobic int `json:"aerobic" gorm:"column:aerobic" binding:"required" example:"100"` //有氧課表訓練數
}
type IntervalTrainingField struct {
	IntervalTraining int `json:"interval_training" gorm:"column:interval_training" binding:"required" example:"100"` //間歇肌力訓練課表訓練數
}
type WeightTrainingField struct {
	WeightTraining int `json:"weight_training" gorm:"column:weight_training" binding:"required" example:"200"` //重量訓練課表訓練數
}
type ResistanceTrainingField struct {
	ResistanceTraining int `json:"resistance_training" gorm:"column:resistance_training" binding:"required" example:"200"` //阻力訓練課表訓練數
}
type BodyweightTrainingField struct {
	BodyweightTraining int `json:"bodyweight_training" gorm:"column:bodyweight_training" binding:"required" example:"200"` //徒手訓練課表訓練數
}
type OtherTrainingField struct {
	OtherTraining int `json:"other_training" gorm:"column:other_training" binding:"required" example:"200"` //付費課表訓練數
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
