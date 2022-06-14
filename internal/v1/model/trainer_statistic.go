package model

type TrainerStatistic struct {
	UserID       int64   `gorm:"column:user_id"`       // 教練id
	StudentCount int     `gorm:"column:student_count"` // 學生總數
	CourseCount  int     `gorm:"column:course_count"`  // 課表總數
	ReviewScore  float64 `gorm:"column:review_score"`  // 課表總評分
	UpdateAt     string  `gorm:"column:update_at"`     // 更新日期
}

func (TrainerStatistic) TableName() string {
	return "trainer_statistics"
}

type SaveTrainerStatisticParam struct {
	StudentCount *int     // 學生總數
	CourseCount  *int     // 課表總數
	ReviewScore  *float64 // 課表總評分
}
