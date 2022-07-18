package trainer_statistic

type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}

type StudentCount struct {
	StudentCount int     `json:"student_count,omitempty" gorm:"column:student_count" example:"10"` // 學生總數
}

type CourseCount struct {
	CourseCount  int     `json:"course_count,omitempty" gorm:"column:course_count" example:"15"`  // 課表總數
}

type ReviewScore struct {
	ReviewScore  float64 `json:"review_score,omitempty" gorm:"column:review_score" example:"4.5"` // 課表總評分
}

type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	UserIDField
	StudentCount
	CourseCount
	ReviewScore
	UpdateAtField
}

func (Table) TableName() string {
	return "trainer_statistics"
}
