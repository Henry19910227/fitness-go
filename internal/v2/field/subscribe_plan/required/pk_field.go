package required

type SubscribePlanIDField struct {
	SubscribePlanID int64 `json:"subscribe_plan_id" gorm:"column:subscribe_plan_id" binding:"required" example:"1"` //訂閱項目id
}
