package validator

type GetLatestCoursesQuery struct {
	Page int `form:"page" binding:"required,min=1" example:"1"` // 頁數
	Size int `form:"size" binding:"required,min=1" example:"5"` // 筆數
}
