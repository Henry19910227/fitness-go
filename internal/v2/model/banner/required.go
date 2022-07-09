package banner

type IDRequired struct {
	ID int64 `json:"id" uri:"banner_id" binding:"required" example:"1"` //id
}
type CourseIDRequired struct {
	CourseID int64 `json:"course_id" binding:"required" example:"10"` //課表id
}
type UserIDRequired struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` //用戶id
}
type ImageRequired struct {
	Image string `json:"image" binding:"required" example:"1234.jpg"` //圖片
}
type TypeRequired struct {
	Type int `json:"type" form:"type" binding:"required" example:"1"` //類型(1:課表/2:教練/3:訂閱)
}