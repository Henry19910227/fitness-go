package validator

type PlanIDUri struct {
	PlanID int64 `uri:"plan_id" binding:"required" example:"1"`
}

type UpdatePlanBody struct {
	Name string `json:"name" binding:"required,min=1,max=20" example:"第一週增肌計畫"`
}