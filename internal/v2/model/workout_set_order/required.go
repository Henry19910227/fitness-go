package workout_set_order

type WorkoutIDRequired struct {
	WorkoutID int64 `json:"workout_id" uri:"workout_id" example:"1"` //訓練id
}
type WorkoutSetIDRequired struct {
	WorkoutSetID int64 `json:"workout_set_id" example:"1"` //訓練組id
}
type SeqRequired struct {
	Seq int `json:"seq" example:"1"` //排序號
}
