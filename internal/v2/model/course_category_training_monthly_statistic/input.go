package course_category_training_monthly_statistic

// APIGetCMSCategoryTrainingStatisticInput /v2/cms/statistic_monthly/course_category/training [GET]
type APIGetCMSCategoryTrainingStatisticInput struct {
	Query APIGetCMSCategoryTrainingStatisticQuery
}
type APIGetCMSCategoryTrainingStatisticQuery struct {
	CategoryRequired
	YearRequired
	MonthRequired
}
