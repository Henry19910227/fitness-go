package model

type CourseUsageStatisticResult struct {
	CourseID int64 `gorm:"column:course_id"`
	Value int `gorm:"column:value"`
}
