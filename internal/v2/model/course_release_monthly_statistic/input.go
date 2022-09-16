package course_release_monthly_statistic

// APIGetCMSCourseReleaseStatisticInput /v2/cms/statistic_monthly/course/release [GET]
type APIGetCMSCourseReleaseStatisticInput struct {
	Query APIGetCMSCourseReleaseStatisticQuery
}
type APIGetCMSCourseReleaseStatisticQuery struct {
	YearRequired
	MonthRequired
}
