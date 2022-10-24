package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" binding:"omitempty" example:"1"` //評論id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty"  uri:"course_id" binding:"omitempty" example:"10"` //課表id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" binding:"omitempty" example:"10001"` //用戶id
}
type ScoreField struct {
	Score *int `json:"score,omitempty" form:"score" binding:"omitempty,oneof=1 2 3 4 5" example:"5"` //評分(1~5分)
}
type BodyField struct {
	Body *string `json:"body,omitempty" form:"body" binding:"omitempty,max=400" example:"很棒的課表"` //內容
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
