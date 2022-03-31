package model

type WorkoutLog struct {
	ID        int64    `gorm:"column:id"`                           // 課表統計報表id
	UserID    int64    `gorm:"column:user_id"`                      // 用戶id
	WorkoutID int64    `gorm:"column:workout_id"`                   // 訓練id
	Duration  int      `gorm:"column:duration"`                     // 訓練時長
	Intensity int      `gorm:"column:intensity"`                    // 訓練強度(0:未指定/1:輕鬆/2:適中/3:稍難/4:很累)
	Place     int      `gorm:"column:place"`                        // 地點(0:未指定/1:住家/2:健身房/3:戶外)
	CreateAt  string   `gorm:"column:create_at"`                    // 更新日期
	Workout   *Workout `gorm:"foreignKey:id;references:workout_id"` // 訓練id
}

type WorkoutLogCourseStatistic struct {
	CourseID                int64 `gorm:"column:course_id"`
	FinishWorkoutCount      int   `gorm:"column:finish_workout_count"`
	TotalFinishWorkoutCount int   `gorm:"column:total_finish_workout_count"`
	Duration                int   `gorm:"column:duration"`
}

type WorkoutLogPlanStatistic struct {
	PlanID             int64 `gorm:"column:plan_id"`              // 計畫id
	FinishWorkoutCount int   `gorm:"column:finish_workout_count"` // 完成訓練數量(去除重複)
	Duration           int   `gorm:"column:duration"`             // 總花費時間(秒)
}

type CreateWorkoutLogParam struct {
	UserID    int64 // 用戶id
	WorkoutID int64 // 訓練id
	Duration  int   // 訓練時長
	Intensity int   // 訓練強度(0:未指定/1:輕鬆/2:適中/3:稍難/4:很累)
	Place     int   // 地點(0:未指定/1:住家/2:健身房/3:戶外)
}
