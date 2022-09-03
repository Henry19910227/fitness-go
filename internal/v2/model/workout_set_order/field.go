package workout_set_order

type WorkoutIDField struct {
	WorkoutID *int64 `json:"workout_id,omitempty" uri:"workout_id" gorm:"column:workout_id" example:"1"` //訓練id
}
type WorkoutSetIDField struct {
	WorkoutSetID *int64 `json:"workout_set_id,omitempty" gorm:"column:workout_set_id" example:"1"` //訓練組id
}
type SeqField struct {
	Seq *int `json:"seq,omitempty" gorm:"column:seq" example:"1"` //排序號
}


type Table struct {
	WorkoutIDField
	WorkoutSetIDField
	SeqField
}

func (Table) TableName() string {
	return "workout_set_orders"
}
