package course_create_monthly_statistic

// APIGetCMSCourseCreateStatisticInput /v2/cms/statistic_monthly/course/create [GET]
type APIGetCMSCourseCreateStatisticInput struct {
	Query APIGetCMSCourseCreateStatisticQuery
}
type APIGetCMSCourseCreateStatisticQuery struct {
	YearRequired
	MonthRequired
}
