package required

type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` //用戶id
}
type StudentCountField struct {
	StudentCount int     `json:"student_count" gorm:"column:student_count" binding:"required" example:"10"` // 學生總數
}
type CourseCountField struct {
	CourseCount  int     `json:"course_count" gorm:"column:course_count" binding:"required" example:"15"`  // 課表總數
}
type ReviewScoreField struct {
	ReviewScore  float64 `json:"review_score" gorm:"column:review_score" binding:"required" example:"4.5"` // 課表總評分
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
