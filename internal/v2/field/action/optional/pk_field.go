package optional

type ActionIDField struct {
	ActionID *int64 `json:"action_id,omitempty" uri:"action_id" gorm:"column:action_id" binding:"omitempty" example:"1"` //動作id
}
