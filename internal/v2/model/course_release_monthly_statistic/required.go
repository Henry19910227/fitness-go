package course_release_monthly_statistic

type YearRequired struct {
	Year int `json:"year" form:"year" binding:"required,max=2500" example:"2022"` //年份
}
type MonthRequired struct {
	Month int `json:"month" form:"month" binding:"required,min=1,max=12" example:"12"` //月份
}
