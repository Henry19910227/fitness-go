package entity

type UserPlanStatistic struct {
	ID                 int64  `gorm:"column:id"`                   // 課表統計報表id
	UserID             int64  `gorm:"column:user_id"`              // 用戶id
	PlanID             int64  `gorm:"column:plan_id"`              // 計畫id
	FinishWorkoutCourt int    `gorm:"column:finish_workout_count"` // 完成訓練數量(去除重複)
	Duration           int    `gorm:"column:duration"`             // 總花費時間(秒)
	UpdateAt           string `gorm:"column:update_at"`            // 更新日期
}

func (UserPlanStatistic) TableName() string {
	return "user_plan_statistics"
}
