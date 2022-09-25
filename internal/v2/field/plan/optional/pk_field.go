package optional

type PlanIDField struct {
	PlanID *int64 `json:"plan_id,omitempty" uri:"plan_id" gorm:"column:plan_id" binding:"omitempty" example:"1"` //計畫id
}
