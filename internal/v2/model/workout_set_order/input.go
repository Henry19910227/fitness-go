package workout_set_order

type DeleteInput struct {
	WorkoutIDRequired
}

// APIUpdateUserWorkoutSetOrdersInput /v2/user/workout/{workout_id}/workout_set_orders [PUT] 修改訓練組的順序
type APIUpdateUserWorkoutSetOrdersInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri APIUpdateUserWorkoutSetOrderUri
	Body APIUpdateUserWorkoutSetOrderBody
}
type APIUpdateUserWorkoutSetOrderUri struct {
	WorkoutIDRequired
}
type APIUpdateUserWorkoutSetOrderBody struct {
	WorkoutSetOrders []struct{
		WorkoutSetIDRequired
		SeqRequired
	} `json:"workout_set_orders" binding:"required"` //訓練組排序
}