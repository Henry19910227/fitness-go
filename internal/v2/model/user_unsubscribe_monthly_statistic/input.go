package user_unsubscribe_monthly_statistic

// APIGetCMSUserUnsubscribeStatisticInput /v2/cms/statistic_monthly/user/unsubscribe [GET]
type APIGetCMSUserUnsubscribeStatisticInput struct {
	Query APIGetCMSUserUnsubscribeStatisticQuery
}
type APIGetCMSUserUnsubscribeStatisticQuery struct {
	YearRequired
	MonthRequired
}
