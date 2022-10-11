package required

type ActionIDField struct {
	ActionID int64 `json:"action_id" uri:"action_id" gorm:"column:action_id" binding:"required" example:"1"` //動作id
}
