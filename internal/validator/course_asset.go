package validator

type GetCourseAssetQuery struct {
	Type int `form:"type" binding:"required,oneof=1 2" example:"1"` // 搜尋類別(1:進行中課表/2:付費課表)
	Page int `form:"page" binding:"required,min=1" example:"1"`     // 頁數
	Size int `form:"size" binding:"required,min=1" example:"5"`     // 筆數
}
