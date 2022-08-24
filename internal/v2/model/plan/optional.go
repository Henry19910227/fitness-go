package plan

type IDOptional struct {
	ID *int64 `json:"id,omitempty" uri:"plan_id" gorm:"column:id" example:"1"` //計畫id
}
