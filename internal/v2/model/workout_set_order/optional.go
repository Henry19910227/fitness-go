package workout_set_order

type WorkoutIDOptional struct {
	WorkoutID *int64 `json:"workout_id,omitempty" uri:"workout_id" example:"1"` //訓練id
}
type WorkoutSetIDOptional struct {
	WorkoutSetID *int64 `json:"workout_set_id,omitempty" example:"1"` //訓練組id
}
type SeqOptional struct {
	Seq *int `json:"seq,omitempty" example:"1"` //排序號
}
