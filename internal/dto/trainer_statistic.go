package dto

type TrainerStatistic struct {
	StudentCount int     `json:"student_count" example:"10"` // 學生總數
	CourseCount  int     `json:"course_count" example:"15"`  // 課表總數
	ReviewScore  float64 `json:"review_score" example:"4.5"` // 課表總評分
}
