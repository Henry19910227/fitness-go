package entity

type CourseStatistic struct {
	ID                int64  `gorm:"column:id"`                  // 課表統計報表id
	UserID            int64  `gorm:"column:user_id"`             // 用戶id
	CourseID          int64  `gorm:"column:course_id"`           // 課表id
	WorkoutCourt      int64  `gorm:"column:workout_count"`       // 完成訓練數量(去除重複)
	TotalWorkoutCourt int64  `gorm:"column:total_workout_count"` // 訓練總量(可重複並累加)
	Duration          int64  `gorm:"column:duration"`            // 總花費時間(秒)
	UpdateAt          string `gorm:"column:update_at"`           // 更新日期
}

func (CourseStatistic) TableName() string {
	return "course_statistics"
}
