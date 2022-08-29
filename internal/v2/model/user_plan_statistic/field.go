package user_plan_statistic

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` // 用戶id
}
type PlanIDField struct {
	PlanID *int64 `json:"plan_id,omitempty" gorm:"column:plan_id" example:"1"` // 計畫 id
}
type FinishWorkoutCountField struct {
	FinishWorkoutCount *int `json:"finish_workout_count,omitempty" gorm:"column:finish_workout_count" example:"10"` // 完成訓練數量(去除重複)
}
type DurationField struct {
	Duration *int `json:"duration,omitempty" gorm:"column:duration" example:"60"` // 總花費時間(秒)
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	UserIDField
	PlanIDField
	FinishWorkoutCountField
	DurationField
	UpdateAtField
}

func (Table) TableName() string {
	return "user_plan_statistics"
}
