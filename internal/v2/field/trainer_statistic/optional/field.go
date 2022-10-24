package optional

type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` //用戶id
}
type StudentCountField struct {
	StudentCount *int     `json:"student_count,omitempty" gorm:"column:student_count" binding:"omitempty" example:"10"` // 學生總數
}
type CourseCountField struct {
	CourseCount  *int     `json:"course_count,omitempty" gorm:"column:course_count" binding:"omitempty" example:"15"`  // 課表總數
}
type ReviewScoreField struct {
	ReviewScore  *float64 `json:"review_score,omitempty" gorm:"column:review_score" binding:"omitempty" example:"4.5"` // 課表總評分
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
