package course_training_monthly_statistic

// APIGetCMSCourseTrainingStatisticInput /v2/cms/statistic_monthly/course/training [GET]
type APIGetCMSCourseTrainingStatisticInput struct {
	Query APIGetCMSCourseTrainingStatisticQuery
}
type APIGetCMSCourseTrainingStatisticQuery struct {
	YearRequired
	MonthRequired
}
