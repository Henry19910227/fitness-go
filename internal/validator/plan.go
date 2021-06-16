package validator

type PlanIDUri struct {
	PlanID int64 `uri:"plan_id" binding:"required" example:"1"`
}
