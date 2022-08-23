package plan

type CourseIDRequired struct {
	CourseID int64 `json:"course_id" uri:"course_id" binding:"required" example:"1"` //課表id
}

type NameRequired struct {
	Name string `json:"name" form:"name" binding:"required,min=1,max=20" example:"減脂計畫"` //計畫名稱
}
