package model

type plan struct {
	ID int64 `gorm:"column:id"` //計畫id
	CourseID int64 `gorm:"column:course_id"` //課表id
	Name int64 `gorm:"column:name"` //計畫名稱
	WorkoutCount int64 `gorm:"column:workout_count"` //訓練數量
	CreateAt string `gorm:"column:create_at"` //創建時間
	UpdateAt string `gorm:"column:update_at"` //更新時間
}
