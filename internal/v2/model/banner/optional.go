package banner

type IDOptional struct {
	ID *int64 `json:"id,omitempty" binding:"omitempty" example:"1"` //id
}
type CourseIDOptional struct {
	CourseID *int64 `json:"course_id,omitempty" form:"course_id" binding:"omitempty" example:"10"` //課表id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" form:"user_id" binding:"omitempty" example:"10001"` //用戶id
}
type ImageOptional struct {
	Image *string `json:"image,omitempty" binding:"omitempty" example:"1234.jpg"` //圖片
}
type TypeOptional struct {
	Type *int `json:"type,omitempty" binding:"omitempty" example:"1"` //類型(1:課表/2:教練/3:訂閱)
}