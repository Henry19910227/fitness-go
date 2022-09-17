package course_category_training_monthly_statistic

type CategoryRequired struct {
	Category int `json:"category" form:"category" binding:"required,oneof=1 2 3 4 5 6" example:"1"` // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
}
type YearRequired struct {
	Year int `json:"year" form:"year" binding:"required,max=2500" example:"2022"` //年份
}
type MonthRequired struct {
	Month int `json:"month" form:"month" binding:"required,min=1,max=12" example:"12"` //月份
}
