package validator

type TokenHeader struct {
	Token string `header:"Token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQ0MDc3NjMsInN1YiI6IjEwMDEzIn0.Z5UeEC8ArCVYej9kI1paXD2f5FMFiTfeLpU6e_CZZw0"`
}

type UserIDUri struct {
	UserID *int64 `uri:"user_id" binding:"required" example:"10001"`
}

type CourseIDUri struct {
	CourseID int64 `uri:"course_id" binding:"required" example:"1"`
}

type PagingQuery struct {
	Page *int `form:"page" binding:"omitempty,min=1" example:"1"`
	Size *int `form:"size" binding:"omitempty,min=1" example:"5"`
}
