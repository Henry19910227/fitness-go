package required

type PlanIDField struct {
	PlanID int64 `json:"plan_id" uri:"plan_id" gorm:"column:plan_id" binding:"required" example:"1"` //計畫id
}
