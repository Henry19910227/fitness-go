package optional

type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` // 用戶 id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" uri:"course_id" gorm:"column:course_id" binding:"omitempty" example:"10"` //課表id
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 創建時間
}
