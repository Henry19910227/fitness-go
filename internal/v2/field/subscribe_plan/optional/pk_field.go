package optional

type SubscribePlanIDField struct {
	SubscribePlanID *int64 `json:"subscribe_plan_id,omitempty" gorm:"column:subscribe_plan_id" binding:"omitempty" example:"1"` //訂閱項目id
}
