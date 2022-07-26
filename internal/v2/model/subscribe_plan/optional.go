package subscribe_plan

type IDOptional struct {
	ID *int64 `json:"id,omitempty" example:"1"` //訂閱項目id
}
type ProductLabelIDOptional struct {
	ProductLabelID *int64 `json:"product_label_id,omitempty" example:"1"` //產品標籤id
}
type PeriodOptional struct {
	Period *int `json:"period,omitempty" example:"12"` //週期(1:一個月/2:二個月/3:三個月/6:六個月/12:一年/99:永久)
}
type NameOptional struct {
	Name *string `json:"name,omitempty" example:"金牌課表"` //銷售名稱
}
type EnableOptional struct {
	Enable *int `json:"enable,omitempty" example:"1"` //是否啟用
}
