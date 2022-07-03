package favorite_course


type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` // 用戶 id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" example:"10"` //課表id
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}



type Table struct {
	UserIDField
	CourseIDField
	CreateAtField
}
func (Table) TableName() string {
	return "favorite_courses"
}