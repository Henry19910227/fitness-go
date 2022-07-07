package review

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //評論id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" example:"10"` //課表id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}
type ScoreField struct {
	Score *int `json:"score,omitempty" gorm:"column:score" example:"5"` //評分
}
type BodyField struct {
	Body *string `json:"body,omitempty" gorm:"column:body" example:"很棒的課表"` //內容
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}

type Table struct {
	IDField
	CourseIDField
	UserIDField
	ScoreField
	BodyField
	CreateAtField
}

func (Table) TableName() string {
	return "reviews"
}
