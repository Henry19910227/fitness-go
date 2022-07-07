package review

type IDOptional struct {
	ID *int64 `json:"id,omitempty" binding:"omitempty" example:"1"` //評論id
}
type CourseIDOptional struct {
	CourseID *int64 `json:"course_id,omitempty" binding:"omitempty" example:"10"` //課表id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" binding:"omitempty" example:"10001"` //用戶id
}
type ScoreOptional struct {
	Score *int `json:"score,omitempty" binding:"omitempty,oneof=1 2 3 4 5" example:"5"` //評分(1~5分)
}
type BodyOptional struct {
	Body *string `json:"body,omitempty" binding:"omitempty" example:"很棒的課表"` //內容
}
